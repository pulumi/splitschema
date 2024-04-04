package bschema

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/pulumi/pulumi/pkg/v3/codegen/schema"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestWriteTo(t *testing.T) {
	pkg := getPackage()
	dir := t.TempDir()
	path := filepath.Join(dir, "schema.bson")
	err := WritePackageSpec(path, pkg)
	if err != nil {
		t.Fatal(err)
	}

	actual, err := ReadPackageSpec(path)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, pkg, actual)
}

func TestStats(t *testing.T) {
	getStats := func(providerName string) *SizeStats {
		bytes, err := os.ReadFile(filepath.Join("testdata", providerName+".json"))
		require.NoError(t, err)

		pkg := schema.PackageSpec{}
		require.NoError(t, json.Unmarshal(bytes, &pkg))

		require.NoError(t, WritePackageSpec(filepath.Join("testdata", providerName+".bson"), pkg))

		stats, err := CalculateSizeStats(pkg)
		require.NoError(t, err)
		return stats
	}

	t.Log("aws", getStats("aws"))
	t.Log("azure-native", getStats("azure-native"))
}

func getPackage() schema.PackageSpec {
	pkg := schema.PackageSpec{
		Name: "test",
		Resources: map[string]schema.ResourceSpec{
			"test": {
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
