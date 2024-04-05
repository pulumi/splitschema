package splitschema

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"slices"
	"sync"

	ccmap "github.com/orcaman/concurrent-map/v2"
	"github.com/pulumi/pulumi/pkg/v3/codegen/schema"
	yaml "gopkg.in/yaml.v3"
)

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

	var waitGroup sync.WaitGroup
	var resourceLoadErr, functionLoadErr, typesLoadErr error

	waitGroup.Add(1)
	go func() {
		core.Resources, resourceLoadErr = p.GetResources()
		waitGroup.Done()
	}()

	waitGroup.Add(1)
	go func() {
		core.Functions, functionLoadErr = p.GetFunctions()
		waitGroup.Done()
	}()

	waitGroup.Add(1)
	go func() {
		core.Types, typesLoadErr = p.GetTypes()
		waitGroup.Done()
	}()

	waitGroup.Wait()

	if resourceLoadErr != nil {
		return nil, resourceLoadErr
	}
	if functionLoadErr != nil {
		return nil, functionLoadErr
	}
	if typesLoadErr != nil {
		return nil, typesLoadErr
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

func (p *partialPackage) GetResources() (map[string]schema.ResourceSpec, error) {
	resourceTokens, err := p.GetResourceTokens()
	if err != nil {
		return nil, err
	}
	var lastErr error
	resources := make([]*schema.ResourceSpec, len(resourceTokens))
	var waitGroup sync.WaitGroup
	for i, v := range resourceTokens {
		waitGroup.Add(1)
		go func(index int, token string) {
			res, err := p.GetResource(token)
			if err != nil {
				lastErr = err
			}
			resources[index] = res
			waitGroup.Done()
		}(i, v)
	}
	waitGroup.Wait()
	if lastErr != nil {
		return nil, lastErr
	}

	resourceMap := make(map[string]schema.ResourceSpec, len(resourceTokens))
	for i, v := range resourceTokens {
		resourceMap[v] = *resources[i]
	}
	return resourceMap, nil
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

func (p *partialPackage) GetFunctions() (map[string]schema.FunctionSpec, error) {
	functionTokens, err := p.GetFunctionTokens()
	if err != nil {
		return nil, err
	}
	var lastErr error
	functions := make([]*schema.FunctionSpec, len(functionTokens))
	var waitGroup sync.WaitGroup
	for i, v := range functionTokens {
		waitGroup.Add(1)
		go func(index int, token string) {
			fn, err := p.GetFunction(token)
			if err != nil {
				lastErr = err
			}
			functions[index] = fn
			waitGroup.Done()
		}(i, v)
	}
	waitGroup.Wait()
	if lastErr != nil {
		return nil, lastErr
	}

	functionMap := make(map[string]schema.FunctionSpec, len(functionTokens))
	for i, v := range functionTokens {
		functionMap[v] = *functions[i]
	}
	return functionMap, nil
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

func (p *partialPackage) GetTypes() (map[string]schema.ComplexTypeSpec, error) {
	typeTokens, err := p.GetTypeTokens()
	if err != nil {
		return nil, err
	}
	var lastErr error
	types := make([]*schema.ComplexTypeSpec, len(typeTokens))
	var waitGroup sync.WaitGroup
	for i, v := range typeTokens {
		waitGroup.Add(1)
		go func(index int, token string) {
			typ, err := p.GetType(token)
			if err != nil {
				lastErr = err
			}
			types[index] = typ
			waitGroup.Done()
		}(i, v)
	}
	waitGroup.Wait()
	if lastErr != nil {
		return nil, lastErr
	}

	typeMap := make(map[string]schema.ComplexTypeSpec, len(typeTokens))
	for i, v := range typeTokens {
		typeMap[v] = *types[i]
	}
	return typeMap, nil
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
