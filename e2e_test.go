package splitschema

import (
	"embed"
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/pulumi/pulumi/pkg/v3/codegen/schema"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

//go:generate go run e2e_testdata.go

//go:embed testdata/aws/*
var awsEmbeddedSplit embed.FS

//go:embed testdata/aws.json
var awsEmbedded []byte

func TestEmbedded(t *testing.T) {
	pkg := NewPartialPackage(awsEmbeddedSplit, "testdata/aws")
	resources, err := pkg.GetResourceTokens()
	require.NoError(t, err)
	assert.NotEmpty(t, resources)
}

func BenchmarkGetResource(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		pkg := NewPartialPackage(awsEmbeddedSplit, "testdata/aws")
		_, err := pkg.GetResource("aws:ec2/instance:Instance")
		require.NoError(b, err)
		b.Log(i, b.Elapsed())
	}
}

func BenchmarkReadResourcesEmbedded(b *testing.B) {
	for i := 0; i < b.N; i++ {
		pkg := NewPartialPackage(awsEmbeddedSplit, "testdata/aws")
		_, err := pkg.ReadPackageSpec()
		if err != nil {
			b.Fatal(err)
		}
		b.Log(i, b.Elapsed())
	}
}

func BenchmarkOldEmbedded(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var pkg schema.PackageSpec
		err := json.Unmarshal(awsEmbedded, &pkg)
		require.NoError(b, err)
		b.Log(i, b.Elapsed())
	}
}

func TestWriteTo(t *testing.T) {
	pkg, err := readPackage(filepath.Join("testdata", "aws.json"))
	require.NoError(t, err)
	dir := t.TempDir()
	err = WritePackageSpec(dir, pkg)
	if err != nil {
		t.Fatal(err)
	}

	// actual, err := ReadPackageSpec(path)
	// if err != nil {
	// 	t.Fatal(err)
	// }

	// assert.Equal(t, pkg, actual)
}

func BenchmarkWriteTo(b *testing.B) {
	pkg, err := readPackage(filepath.Join("testdata", "aws.json"))
	require.NoError(b, err)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		dir := b.TempDir()
		err := WritePackageSpec(dir, pkg)
		if err != nil {
			b.Fatal(err)
		}
		b.Log(b.Elapsed())
	}
}

func BenchmarkReadFrom(b *testing.B) {
	pkg, err := readPackage(filepath.Join("testdata", "aws.json"))
	require.NoError(b, err)
	b.Log("combined read", b.Elapsed())
	dir := b.TempDir()
	err = WritePackageSpec(dir, pkg)
	require.NoError(b, err)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := ReadPackageSpec(dir)
		if err != nil {
			b.Fatal(err)
		}
		b.Log("distributed read", b.Elapsed())
	}
}

func TestAwsWrite(t *testing.T) {
	awsPath := filepath.Join("testdata", "aws")
	pkg, err := readPackage(filepath.Join("testdata", "aws.json"))
	require.NoError(t, err)
	assert.NotNil(t, pkg)
	err = WritePackageSpec(awsPath, pkg, WriteOptionCompact())
	require.NoError(t, err)
}

func TestAwsRoundTrip(t *testing.T) {
	pkg, err := readPackage(filepath.Join("testdata", "aws.json"))
	require.NoError(t, err)
	assert.NotNil(t, pkg)

	dir := t.TempDir()
	err = WritePackageSpec(dir, pkg)
	require.NoError(t, err)

	readSpec, err := ReadPackageSpec(dir)
	require.NoError(t, err)
	require.NotNil(t, readSpec)
	assert.Equal(t, len(pkg.Resources), len(readSpec.Resources))
	assert.Equal(t, len(pkg.Types), len(readSpec.Types))
	assert.Equal(t, len(pkg.Functions), len(readSpec.Functions))

	assert.Equal(t, pkg.AllowedPackageNames, readSpec.AllowedPackageNames)
	assert.Equal(t, pkg.Attribution, readSpec.Attribution)
	assert.Equal(t, pkg.Config, readSpec.Config)
	assert.Equal(t, pkg.Description, readSpec.Description)
	assert.Equal(t, pkg.DisplayName, readSpec.DisplayName)
	assert.Equal(t, pkg.Homepage, readSpec.Homepage)
	assert.Equal(t, pkg.Keywords, readSpec.Keywords)
	assert.Equal(t, pkg.Language, readSpec.Language)
	assert.Equal(t, pkg.License, readSpec.License)
	assert.Equal(t, pkg.LogoURL, readSpec.LogoURL)
	assert.Equal(t, pkg.Meta, readSpec.Meta)
	assert.Equal(t, pkg.Name, readSpec.Name)
	assert.Equal(t, pkg.PluginDownloadURL, readSpec.PluginDownloadURL)
	assert.Equal(t, pkg.Provider, readSpec.Provider)
	assert.Equal(t, pkg.Publisher, readSpec.Publisher)
	assert.Equal(t, pkg.Repository, readSpec.Repository)
	assert.Equal(t, pkg.Version, readSpec.Version)

	for tok, readResource := range readSpec.Resources {
		assert.Equal(t, pkg.Resources[tok], readResource, "resource %s", tok)
	}

	for tok, readType := range readSpec.Types {
		assert.Equal(t, pkg.Types[tok], readType, "type %s", tok)
	}

	for tok, readFunction := range readSpec.Functions {
		assert.Equal(t, pkg.Functions[tok], readFunction, "function %s", tok)
	}

	// When this full equality check fails, it's often slow so we use the above specific comparisons which are faster, and comment this out.
	assert.Equal(t, pkg, readSpec)
}

func readPackage(path string) (*schema.PackageSpec, error) {
	bytes, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var pkg schema.PackageSpec
	err = json.Unmarshal(bytes, &pkg)
	if err != nil {
		return nil, err
	}
	return &pkg, nil
}

func TestMetadataRoundTrip(t *testing.T) {
	token := "test:index:resource"
	pkg := schema.PackageSpec{
		Name: "test",
		Resources: map[string]schema.ResourceSpec{
			token: {
				InputProperties: map[string]schema.PropertySpec{
					"foo": {
						TypeSpec: schema.TypeSpec{
							Type: "string",
						},
					},
				},
			},
		},
	}
	metadata := TypedPackageMetadata[map[string]any, any, any]{
		Resources: map[string]map[string]any{
			token: map[string]interface{}{
				"metaKey": "metaValue",
			},
		},
	}
	dir := t.TempDir()
	require.NoError(t, WritePackageSpecWithTypedMetadata(dir, &pkg, &metadata))

	readSpec := NewLocalPartialPackageWithMetadata[map[string]interface{}, any, any](dir)
	actual, err := readSpec.GetResourceMeta(token)
	require.NoError(t, err)
	assert.Equal(t, metadata.Resources[token], *actual)
}
