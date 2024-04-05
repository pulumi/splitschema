package splitschema

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"slices"
	"sync"
	"sync/atomic"

	ccmap "github.com/orcaman/concurrent-map/v2"
	"github.com/pulumi/pulumi/pkg/v3/codegen/schema"
	yaml "gopkg.in/yaml.v3"
)

func NewLocalPartialPackage(basePath string) partialPackage {
	return NewPartialPackage(os.DirFS(basePath), ".")
}

func NewPartialPackage(fs fs.FS, basePath string) partialPackage {
	return partialPackage{
		reader:    newReader(fs, basePath, "json"),
		resources: ccmap.New[*schema.ResourceSpec](),
		functions: ccmap.New[*schema.FunctionSpec](),
		types:     ccmap.New[*schema.ComplexTypeSpec](),
	}
}

func ReadPackageSpec(path string) (*schema.PackageSpec, error) {
	partialPkg := NewLocalPartialPackage(path)
	return partialPkg.ReadPackageSpec()
}

type partialPackage struct {
	core *schema.PackageSpec

	reader         reader
	resourceTokens atomic.Pointer[tokenMappings]
	resources      ccmap.ConcurrentMap[string, *schema.ResourceSpec]

	functionTokens atomic.Pointer[tokenMappings]
	functions      ccmap.ConcurrentMap[string, *schema.FunctionSpec]

	typeTokens atomic.Pointer[tokenMappings]
	types      ccmap.ConcurrentMap[string, *schema.ComplexTypeSpec]
}

func NewLocalPartialPackageWithMetadata[ResourceMeta any, FunctionMeta any, TypeMeta any](basePath string) partialPackageWithMetadata[ResourceMeta, FunctionMeta, TypeMeta] {
	return NewPartialPackageWithMetadata[ResourceMeta, FunctionMeta, TypeMeta](os.DirFS(basePath), ".")
}

func NewPartialPackageWithMetadata[ResourceMeta any, FunctionMeta any, TypeMeta any](fs fs.FS, basePath string) partialPackageWithMetadata[ResourceMeta, FunctionMeta, TypeMeta] {
	return partialPackageWithMetadata[ResourceMeta, FunctionMeta, TypeMeta]{
		partialPackage: NewPartialPackage(fs, basePath),
		resourceMeta:   ccmap.New[*ResourceMeta](),
		functionMeta:   ccmap.New[*FunctionMeta](),
		typeMeta:       ccmap.New[*TypeMeta](),
	}
}

type partialPackageWithMetadata[ResourceMeta any, FunctionMeta any, TypeMeta any] struct {
	partialPackage
	resourceMeta ccmap.ConcurrentMap[string, *ResourceMeta]
	functionMeta ccmap.ConcurrentMap[string, *FunctionMeta]
	typeMeta     ccmap.ConcurrentMap[string, *TypeMeta]
}

type tokenMappings struct {
	mapping map[string]string
	list    []string
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
	if err := p.reader.readData("core", &core); err != nil {
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
	path, err := getPath(token, "resources")
	if err != nil {
		return nil, err
	}
	var spec schema.ResourceSpec
	description, err := p.reader.readSpec(path, &spec)
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
	mappings, err := p.getResourceTokenMappings()
	return mappings.list, err
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
	path, err := getPath(token, "functions")
	if err != nil {
		return nil, err
	}
	var spec schema.FunctionSpec
	description, err := p.reader.readSpec(path, &spec)
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
	mappings, err := p.getFunctionTokenMappings()
	return mappings.list, err
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
	path, err := getPath(token, "types")
	if err != nil {
		return nil, err
	}
	var spec schema.ComplexTypeSpec
	description, err := p.reader.readSpec(path, &spec)
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
	mappings, err := p.getTypeTokenMappings()
	return mappings.list, err
}

// getResourceTokenMappings returns the resource token mappings and a sorted list of resource tokens.
func (p *partialPackage) getResourceTokenMappings() (*tokenMappings, error) {
	return p.getTokenMappings(&p.resourceTokens, "resources")
}

func (p *partialPackage) getTokenMappings(tokenMappingsPtr *atomic.Pointer[tokenMappings], pathExExt string) (*tokenMappings, error) {
	if mappings := tokenMappingsPtr.Load(); mappings != nil {
		return mappings, nil
	}

	mappings := tokenMappings{}

	if err := p.reader.readData(pathExExt, &mappings.mapping); err != nil {
		return nil, err
	}
	mappings.list = make([]string, 0, len(mappings.mapping))
	for token := range mappings.mapping {
		mappings.list = append(mappings.list, token)
	}
	slices.Sort(mappings.list)

	if !tokenMappingsPtr.CompareAndSwap(nil, &mappings) {
		return tokenMappingsPtr.Load(), nil
	}

	return &mappings, nil
}

func (p *partialPackage) getTypeTokenMappings() (*tokenMappings, error) {
	return p.getTokenMappings(&p.typeTokens, "types")
}

func (p *partialPackage) getFunctionTokenMappings() (*tokenMappings, error) {
	return p.getTokenMappings(&p.functionTokens, "functions")
}

func (p *partialPackageWithMetadata[ResourceMeta, FunctionMeta, TypeMeta]) GetResourceMeta(token string) (*ResourceMeta, error) {
	return getMetadata(&p.resourceMeta, &p.reader, "resources", token)
}

func (p *partialPackageWithMetadata[ResourceMeta, FunctionMeta, TypeMeta]) GetFunctionMeta(token string) (*FunctionMeta, error) {
	return getMetadata(&p.functionMeta, &p.reader, "functions", token)
}

func (p *partialPackageWithMetadata[ResourceMeta, FunctionMeta, TypeMeta]) GetTypeMeta(token string) (*TypeMeta, error) {
	return getMetadata(&p.typeMeta, &p.reader, "types", token)
}

type reader struct {
	fs       fs.FS
	basePath string
	format   string
}

func newReader(fs fs.FS, basePath, format string) reader {
	return reader{fs: fs, basePath: basePath, format: format}
}

func (r *reader) readSpec(path string, data any) (description *string, err error) {
	descriptionBytes, err := r.ReadFile(path + ".md")
	if err != nil {
		if !os.IsNotExist(err) {
			return nil, err
		}
	} else {
		descriptionStr := string(descriptionBytes)
		description = &descriptionStr
	}
	err = r.readData(path, data)
	if err != nil {
		return nil, err
	}
	return description, nil
}

func (r *reader) readData(pathExExt string, data any) error {
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

func getMetadata[T any](cache *ccmap.ConcurrentMap[string, *T], reader *reader, kind, token string) (*T, error) {
	if spec, ok := cache.Get(token); ok {
		return spec, nil
	}
	path, err := getPath(token, kind)
	if err != nil {
		return nil, err
	}

	var spec T
	err = reader.readData(path+".meta", &spec)
	if err != nil {
		return nil, err
	}
	if !cache.SetIfAbsent(token, &spec) {
		// Use the first written spec if another goroutine wrote the spec first.
		if spec, ok := cache.Get(token); ok {
			return spec, nil
		}
	}
	return &spec, nil
}
