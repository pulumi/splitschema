package bschema

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
	err = WritePackageSpec(awsPath, pkg)
	require.NoError(t, err)
}

func TestAws(t *testing.T) {
	awsPath := filepath.Join("testdata", "aws")
	pkg, err := readPackage(filepath.Join("testdata", "aws.json"))
	require.NoError(t, err)
	assert.NotNil(t, pkg)
	// err = WritePackageSpec(awsPath, pkg)
	// require.NoError(t, err)

	readSpec, err := ReadPackageSpec(awsPath)
	require.NoError(t, err)
	require.NotNil(t, readSpec)
	assert.Equal(t, len(pkg.Resources), len(readSpec.Resources))
	assert.Equal(t, len(pkg.Types), len(readSpec.Types))
	assert.Equal(t, len(pkg.Functions), len(readSpec.Functions))
	assert.Equal(t, pkg.Provider, readSpec.Provider)

	pp := NewPartialFsPackage(awsPath)
	resourceTokens, err := pp.GetResourceTokens()
	require.NoError(t, err)
	for _, tok := range resourceTokens {
		res, err := pp.GetResource(tok)
		require.NoError(t, err)
		assert.Equal(t, pkg.Resources[tok], *res, "resource %s", tok)
	}
}

func readPackage(path string) (*schema.PackageSpec, error) {
	var pkg schema.PackageSpec
	bytes, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(bytes, &pkg)
	if err != nil {
		return nil, err
	}
	return &pkg, nil
}
