package splitschema

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/pulumi/pulumi/pkg/v3/codegen/schema"
	yaml "gopkg.in/yaml.v3"
)

func WritePackageSpec(path string, pkg *schema.PackageSpec, opts ...WriteOption) error {
	return WritePackageSpecWithMetadata(path, pkg, nil, opts...)
}

func WritePackageSpecWithMetadata(path string, pkg *schema.PackageSpec, metadata *PackageMetadata, opts ...WriteOption) error {
	return WritePackageSpecWithTypedMetadata(path, pkg, metadata, opts...)
}

func WritePackageSpecWithTypedMetadata[Resource, Function, Type any](path string, pkg *schema.PackageSpec, metadata *TypedPackageMetadata[Resource, Function, Type], opts ...WriteOption) error {
	options := &WriteOptions{}
	for _, opt := range opts {
		opt.Apply(options)
	}

	pkgCopy := *pkg
	functions := pkg.Functions
	pkgCopy.Functions = nil
	types := pkg.Types
	pkgCopy.Types = nil
	resources := pkg.Resources
	pkgCopy.Resources = nil

	indent := "    "
	if options.Compact {
		indent = ""
	}
	writer := NewWriter(path, "json", indent)
	if err := writer.WriteData("core", pkgCopy, ""); err != nil {
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
	if metadata != nil {
		for token, resourceMetadata := range metadata.Resources {
			path, err := writer.WriteMetadata(token, "resources", resourceMetadata)
			if err != nil {
				return err
			}
			resourceTokens[token] = path
		}
	}
	if err := writer.WriteData("resources", resourceTokens, ""); err != nil {
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
	if metadata != nil {
		for token, functionMetadata := range metadata.Functions {
			path, err := writer.WriteSpec(token, "functions", functionMetadata, "")
			if err != nil {
				return err
			}
			functionTokens[token] = path
		}
	}
	if err := writer.WriteData("functions", functionTokens, ""); err != nil {
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
	if metadata != nil {
		for token, typeMetadata := range metadata.Types {
			path, err := writer.WriteSpec(token, "types", typeMetadata, "")
			if err != nil {
				return err
			}
			typeTokens[token] = path
		}
	}
	if err := writer.WriteData("types", typeTokens, ""); err != nil {
		return err
	}

	return nil
}

func WriteOptionCompact() WriteOption {
	return optionFunc(func(opts *WriteOptions) {
		opts.Compact = true
	})
}

type WriteOption interface {
	Apply(*WriteOptions)
}

type WriteOptions struct {
	Compact bool
}

type optionFunc func(*WriteOptions)

func (o optionFunc) Apply(opts *WriteOptions) {
	o(opts)
}

type PackageMetadata = TypedPackageMetadata[any, any, any]

type TypedPackageMetadata[Resource, Function, Type any] struct {
	Resources map[string]Resource
	Functions map[string]Function
	Types     map[string]Type
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

	if err := w.WriteData(path, data, "        "); err != nil {
		return "", err
	}
	return path, nil
}

func (w *writer) WriteMetadata(token string, kind string, data any) (string, error) {
	// Handle module names which include type name after "/"
	path, err := getPath(token, kind)
	if err != nil {
		return "", err
	}

	if err := w.WriteData(path+".meta", data, ""); err != nil {
		return "", err
	}
	return path, nil
}

func (w *writer) WriteData(pathExExt string, data any, prefix string) error {
	var bytes []byte
	var err error
	var path string
	if w.format == "json" {
		path = pathExExt + ".json"
		if w.indent != "" {
			bytes, err = json.MarshalIndent(data, prefix, w.indent)
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
