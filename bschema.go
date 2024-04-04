package bschema

import (
	"encoding/json"
	"io/fs"
	"os"

	"github.com/pulumi/pulumi/pkg/v3/codegen/schema"
	"go.mongodb.org/mongo-driver/bson"
)

func WritePackageSpec(path string, pkg schema.PackageSpec) error {
	bytes, err := bson.Marshal(pkg)
	if err != nil {
		return err
	}
	return os.WriteFile(path, bytes, fs.FileMode(0644))
}

func ReadPackageSpec(path string) (schema.PackageSpec, error) {
	bytes, err := os.ReadFile(path)
	if err != nil {
		return schema.PackageSpec{}, err
	}
	var pkg schema.PackageSpec
	err = bson.Unmarshal(bytes, &pkg)
	return pkg, err
}

type SizeStats struct {
	BsonBytes int
	JsonBytes int
}

func CalculateSizeStats(pkg schema.PackageSpec) (*SizeStats, error) {
	bsonBytes, err := bson.Marshal(pkg)
	if err != nil {
		return nil, err
	}
	jsonBytes, _ := json.Marshal(pkg)
	if err != nil {
		return nil, err
	}
	return &SizeStats{
		BsonBytes: len(bsonBytes),
		JsonBytes: len(jsonBytes),
	}, nil
}
