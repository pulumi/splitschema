package bschema

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/pulumi/pulumi/pkg/v3/codegen/schema"
	"github.com/stretchr/testify/require"
)

func TestWriteTo(t *testing.T) {
	pkg := getPackage()
	dir := t.TempDir()
	err := WritePackageSpec(dir, &pkg)
	if err != nil {
		t.Fatal(err)
	}

	// actual, err := ReadPackageSpec(path)
	// if err != nil {
	// 	t.Fatal(err)
	// }

	// assert.Equal(t, pkg, actual)
}

func TestAws(t *testing.T) {
	pkg, err := readPackage(filepath.Join("testdata", "aws.json"))
	require.NoError(t, err)
	err = WritePackageSpec(filepath.Join("testdata", "aws"), pkg)
	require.NoError(t, err)
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

func getPackage() schema.PackageSpec {
	pkg := schema.PackageSpec{
		Name: "test",
		Resources: map[string]schema.ResourceSpec{
			"pkg:myMod:test": {
				ObjectTypeSpec: schema.ObjectTypeSpec{
					Description: "test",
					Type:        "object",
					Properties: map[string]schema.PropertySpec{
						"test": {
							Description: "test",
							TypeSpec: schema.TypeSpec{
								Type: "string",
							},
						},
					},
				},
			},
		},
	}
	return pkg
}
