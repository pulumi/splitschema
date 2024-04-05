package bschema

import (
	"encoding/json"
	"fmt"
	"hash/crc32"
	"io/fs"
	"os"
	"path/filepath"
	"slices"
	"strings"
	"sync"

	ccmap "github.com/orcaman/concurrent-map/v2"
	"github.com/pulumi/pulumi/pkg/v3/codegen/schema"
	"github.com/pulumi/pulumi/sdk/v3/go/common/tokens"
	yaml "gopkg.in/yaml.v3"
)

func WritePackageSpec(path string, pkg *schema.PackageSpec) error {
	pkgCopy := *pkg
	functions := pkg.Functions
	pkgCopy.Functions = nil
	types := pkg.Types
	pkgCopy.Types = nil
	resources := pkg.Resources
	pkgCopy.Resources = nil

	writer := NewWriter(path, "json", "    ")
	if err := writer.WriteData("core", pkgCopy); err != nil {
		return err
	}

	typeTokens := make(map[string]string, len(types))
	for token, typ := range types {
		path, err := writer.WriteType(token, typ)
		if err != nil {
			return err
		}
		typeTokens[token] = path
	}
	if err := writer.WriteData("types", typeTokens); err != nil {
		return err
	}

	resourceTokens := make(map[string]string, len(resources))
	for token, resource := range resources {
		path, err := writer.WriteResource(token, resource)
		if err != nil {
			return err
		}
		resourceTokens[token] = path
	}
	if err := writer.WriteData("resources", resourceTokens); err != nil {
		return err
	}

	functionTokens := make(map[string]string, len(functions))
	for token, function := range functions {
		path, err := writer.WriteFunction(token, function)
		if err != nil {
			return err
		}
		functionTokens[token] = path
	}
	if err := writer.WriteData("functions", functionTokens); err != nil {
		return err
	}

	return nil
}

func ReadPackageSpec(path string) (*schema.PackageSpec, error) {
	partialPkg := NewPartialFsPackage(path)
	return partialPkg.ReadPackageSpec()
}

type partialPackage struct {
	core *schema.PackageSpec

	reader                reader
	resourceTokenMappings map[string]string
	resourceTokensList    []string
	resourceTokensRead    bool
	resourceTokensLock    sync.Mutex
	resources             ccmap.ConcurrentMap[string, *schema.ResourceSpec]

	functionTokenMappings map[string]string
	functionTokensList    []string
	functionTokensRead    bool
	functionTokensLock    sync.Mutex
	functions             ccmap.ConcurrentMap[string, *schema.FunctionSpec]

	typeTokenMappings map[string]string
	typeTokensList    []string
	typeTokensRead    bool
	typeTokenLock     sync.Mutex
	types             ccmap.ConcurrentMap[string, *schema.ComplexTypeSpec]
}

func NewPartialFsPackage(basePath string) partialPackage {
	return NewPartialPackage(os.DirFS(basePath), ".")
}

func NewPartialPackage(fs fs.FS, basePath string) partialPackage {
	return partialPackage{
		reader:    NewReader(fs, basePath, "json"),
		resources: ccmap.New[*schema.ResourceSpec](),
		types:     ccmap.New[*schema.ComplexTypeSpec](),
		functions: ccmap.New[*schema.FunctionSpec](),
	}
}

func (p *partialPackage) ReadPackageSpec() (*schema.PackageSpec, error) {
	core, err := p.getCore()
	if err != nil {
		return nil, err
	}
	resourceTokens, err := p.GetResourceTokens()
	if err != nil {
		return nil, err
	}
	core.Resources = make(map[string]schema.ResourceSpec, len(resourceTokens))
	for _, v := range resourceTokens {
		res, err := p.GetResource(v)
		if err != nil {
			return nil, err
		}
		core.Resources[v] = *res
	}

	functionTokens, err := p.GetFunctionTokens()
	if err != nil {
		return nil, err
	}
	core.Functions = make(map[string]schema.FunctionSpec, len(functionTokens))
	for _, v := range functionTokens {
		fn, err := p.GetFunction(v)
		if err != nil {
			return nil, err
		}
		core.Functions[v] = *fn
	}

	typeTokens, err := p.GetTypeTokens()
	if err != nil {
		return nil, err
	}
	core.Types = make(map[string]schema.ComplexTypeSpec, len(typeTokens))
	for _, v := range typeTokens {
		typ, err := p.GetType(v)
		if err != nil {
			return nil, err
		}
		core.Types[v] = *typ
	}
	return core, nil
}

func (p *partialPackage) getCore() (*schema.PackageSpec, error) {
	if p.core != nil {
		return p.core, nil
	}
	var core schema.PackageSpec
	if err := p.reader.ReadData("core", &core); err != nil {
		return nil, err
	}
	p.core = &core
	return p.core, nil
}

func (p *partialPackage) GetResource(token string) (*schema.ResourceSpec, error) {
	if spec, ok := p.resources.Get(token); ok {
		return spec, nil
	}
	tokens, _, err := p.getResourceTokenMappings()
	if err != nil {
		return nil, err
	}
	path, ok := tokens[token]
	if !ok {
		return nil, nil
	}
	var spec schema.ResourceSpec
	description, err := p.reader.ReadSpec(path, &spec)
	if err != nil {
		return nil, err
	}
	if description != nil {
		spec.Description = *description
	}
	if !p.resources.SetIfAbsent(token, &spec) {
		// Use the first written spec if another goroutine wrote the spec first.
		if spec, ok := p.resources.Get(token); ok {
			return spec, nil
		}
	}
	return &spec, nil
}

// GetResourceTokens returns the resource tokens for the package sorted alphabetically.
func (p *partialPackage) GetResourceTokens() ([]string, error) {
	_, tokens, err := p.getResourceTokenMappings()
	return tokens, err
}

func (p *partialPackage) GetFunction(token string) (*schema.FunctionSpec, error) {
	if spec, ok := p.functions.Get(token); ok {
		return spec, nil
	}
	tokens, _, err := p.getFunctionTokenMappings()
	if err != nil {
		return nil, err
	}
	path, ok := tokens[token]
	if !ok {
		return nil, nil
	}
	var spec schema.FunctionSpec
	description, err := p.reader.ReadSpec(path, &spec)
	if err != nil {
		return nil, err
	}
	if description != nil {
		spec.Description = *description
	}
	if !p.functions.SetIfAbsent(token, &spec) {
		// Use the first written spec if another goroutine wrote the spec first.
		if spec, ok := p.functions.Get(token); ok {
			return spec, nil
		}
	}
	return &spec, nil
}

func (p *partialPackage) GetFunctionTokens() ([]string, error) {
	_, tokens, err := p.getFunctionTokenMappings()
	return tokens, err
}

func (p *partialPackage) GetType(token string) (*schema.ComplexTypeSpec, error) {
	if spec, ok := p.types.Get(token); ok {
		return spec, nil
	}
	tokens, _, err := p.getTypeTokenMappings()
	if err != nil {
		return nil, err
	}
	path, ok := tokens[token]
	if !ok {
		return nil, nil
	}
	var spec schema.ComplexTypeSpec
	description, err := p.reader.ReadSpec(path, &spec)
	if err != nil {
		return nil, err
	}
	if description != nil {
		spec.Description = *description
	}
	if !p.types.SetIfAbsent(token, &spec) {
		// Use the first written spec if another goroutine wrote the spec first.
		if spec, ok := p.types.Get(token); ok {
			return spec, nil
		}
	}
	return &spec, nil
}

func (p *partialPackage) GetTypeTokens() ([]string, error) {
	_, tokens, err := p.getTypeTokenMappings()
	return tokens, err
}

// getResourceTokenMappings returns the resource token mappings and a sorted list of resource tokens.
func (p *partialPackage) getResourceTokenMappings() (map[string]string, []string, error) {
	if p.resourceTokensRead {
		return p.resourceTokenMappings, p.resourceTokensList, nil
	}

	p.resourceTokensLock.Lock()
	defer p.resourceTokensLock.Unlock()
	if p.resourceTokensRead {
		return p.resourceTokenMappings, p.resourceTokensList, nil
	}

	var resourceTokenMappings map[string]string
	if err := p.reader.ReadData("resources", &resourceTokenMappings); err != nil {
		return nil, nil, err
	}
	resourceTokensList := make([]string, 0, len(resourceTokenMappings))
	for token := range resourceTokenMappings {
		resourceTokensList = append(resourceTokensList, token)
	}
	slices.Sort(resourceTokensList)

	p.resourceTokenMappings = resourceTokenMappings
	p.resourceTokensList = resourceTokensList
	p.resourceTokensRead = true
	return p.resourceTokenMappings, p.resourceTokensList, nil
}

func (p *partialPackage) getTypeTokenMappings() (map[string]string, []string, error) {
	if p.typeTokensRead {
		return p.typeTokenMappings, p.typeTokensList, nil
	}

	p.typeTokenLock.Lock()
	defer p.typeTokenLock.Unlock()
	if p.typeTokensRead {
		return p.typeTokenMappings, p.typeTokensList, nil
	}

	var typeTokenMappings map[string]string
	if err := p.reader.ReadData("types", &typeTokenMappings); err != nil {
		return nil, nil, err
	}
	typeTokensList := make([]string, 0, len(typeTokenMappings))
	for token := range typeTokenMappings {
		typeTokensList = append(typeTokensList, token)
	}
	slices.Sort(typeTokensList)

	p.typeTokenMappings = typeTokenMappings
	p.typeTokensList = typeTokensList
	p.typeTokensRead = true
	return p.typeTokenMappings, p.typeTokensList, nil
}

func (p *partialPackage) getFunctionTokenMappings() (map[string]string, []string, error) {
	if p.functionTokensRead {
		return p.functionTokenMappings, p.functionTokensList, nil
	}

	p.functionTokensLock.Lock()
	defer p.functionTokensLock.Unlock()
	if p.functionTokensRead {
		return p.functionTokenMappings, p.functionTokensList, nil
	}

	var functionTokenMappings map[string]string
	if err := p.reader.ReadData("functions", &functionTokenMappings); err != nil {
		return nil, nil, err
	}
	functionTokensList := make([]string, 0, len(functionTokenMappings))
	for token := range functionTokenMappings {
		functionTokensList = append(functionTokensList, token)
	}
	slices.Sort(functionTokensList)
	p.functionTokenMappings = functionTokenMappings
	p.functionTokensList = functionTokensList
	p.functionTokensRead = true
	return p.functionTokenMappings, p.functionTokensList, nil
}

type writer struct {
	basePath string
	format   string
	indent   string
}

func NewWriter(basePath, format, indent string) writer {
	return writer{basePath: basePath, format: format, indent: indent}
}

func (w *writer) WriteType(token string, spec schema.ComplexTypeSpec) (string, error) {
	var markdown string
	if strings.ContainsRune(spec.Description, '\n') {
		markdown = spec.Description
		spec.Description = ""
	}
	return w.WriteSpec(token, "types", spec, markdown)
}

func (w *writer) WriteResource(token string, spec schema.ResourceSpec) (string, error) {
	var markdown string
	if strings.ContainsRune(spec.Description, '\n') {
		markdown = spec.Description
		spec.Description = ""
	}
	return w.WriteSpec(token, "resources", spec, markdown)
}

func (w *writer) WriteFunction(token string, spec schema.FunctionSpec) (string, error) {
	var markdown string
	if strings.ContainsRune(spec.Description, '\n') {
		markdown = spec.Description
		spec.Description = ""
	}
	return w.WriteSpec(token, "functions", spec, markdown)
}

func (w *writer) WriteSpec(token string, kind string, data any, markdown string) (string, error) {
	// Handle module names which include type name after "/"
	path, err := getPath(token, kind)
	if err != nil {
		return "", err
	}
	if markdown != "" {
		if err := w.WriteFile(path+".md", []byte(markdown)); err != nil {
			return "", err
		}
	}

	if err := w.WriteData(path, data); err != nil {
		return "", err
	}
	return path, nil
}

func (w *writer) WriteData(pathExExt string, data any) error {
	var bytes []byte
	var err error
	var path string
	if w.format == "json" {
		path = pathExExt + ".json"
		if w.indent != "" {
			bytes, err = json.MarshalIndent(data, "", w.indent)
		} else {
			bytes, err = json.Marshal(data)
		}
	} else if w.format == "yaml" {
		path = pathExExt + ".yaml"
		bytes, err = yaml.Marshal(data)
	} else {
		return fmt.Errorf("unsupported format: %s", w.format)
	}
	if err != nil {
		return err
	}
	return w.WriteFile(path, bytes)
}

func (w *writer) WriteFile(path string, bytes []byte) error {
	if err := os.MkdirAll(filepath.Join(w.basePath, filepath.Dir(path)), 0755); err != nil {
		return err
	}
	return os.WriteFile(filepath.Join(w.basePath, path), bytes, fs.FileMode(0644))
}

type reader struct {
	fs       fs.FS
	basePath string
	format   string
}

func NewReader(fs fs.FS, basePath, format string) reader {
	return reader{fs: fs, basePath: basePath, format: format}
}

func (r *reader) ReadSpec(path string, data any) (description *string, err error) {
	descriptionBytes, err := r.ReadFile(path + ".md")
	if err != nil {
		if !os.IsNotExist(err) {
			return nil, err
		}
	} else {
		descriptionStr := string(descriptionBytes)
		description = &descriptionStr
	}
	err = r.ReadData(path, data)
	if err != nil {
		return nil, err
	}
	return description, nil
}

func (r *reader) ReadData(pathExExt string, data any) error {
	if r.format == "json" {
		bytes, err := r.ReadFile(pathExExt + ".json")
		if err != nil {
			return err
		}
		return json.Unmarshal(bytes, data)
	}
	if r.format == "yaml" {
		bytes, err := r.ReadFile(pathExExt + ".yaml")
		if err != nil {
			return err
		}
		return yaml.Unmarshal(bytes, data)
	}
	return fmt.Errorf("unsupported format: %s", r.format)
}

func (r *reader) ReadFile(path string) ([]byte, error) {
	return fs.ReadFile(r.fs, filepath.Join(r.basePath, path))
}

func getPath(token string, kind string) (string, error) {
	t, err := tokens.ParseTypeToken(token)
	if err != nil {
		return "", err
	}

	modName := t.Module().Name().String()
	typeName := t.Name().String()

	if parts := strings.Split(modName, "/"); len(parts) > 0 {
		modName = parts[0]
	}
	filename := makeFileName(typeName)
	path := filepath.Join(modName, kind, filename)
	return path, nil
}

func makeFileName(s string) string {
	return strings.ToLower(s) + "-" + shortHash(s)
}

func shortHash(s string) string {
	return fmt.Sprintf("%08x", crc32.Checksum([]byte(s), crcTable))
}

var crcTable = crc32.MakeTable(crc32.Castagnoli)
