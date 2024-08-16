// Copyright 2024, Pulumi Corporation.

// The following directive is necessary to make the package coherent:

//go:build ignore
// +build ignore

// This program generates contributors.go. It can be invoked by running
// go generate ./testdata from the parent directory.
package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/pulumi/pulumi/pkg/v3/codegen/schema"
	"github.com/pulumi/splitschema"
)

func main() {
	if err := prepareProviderSchema("aws", "v6.29.0"); err != nil {
		panic(err)
	}
	if err := prepareProviderSchema("azure-native", "v2.35.0"); err != nil {
		panic(err)
	}
}

func prepareProviderSchema(provider, version string) error {
	bytes, err := downloadSchema(provider, version)
	if err != nil {
		return err
	}

	var pkg schema.PackageSpec
	err = json.Unmarshal(bytes, &pkg)
	if err != nil {
		return err
	}

	if err = os.WriteFile(filepath.Join("testdata", provider+".json"), bytes, 0644); err != nil {
		return err
	}

	return splitschema.WritePackageSpec(filepath.Join("testdata", provider), &pkg)
}

func downloadSchema(provider, version string) ([]byte, error) {
	res, err := http.Get(fmt.Sprintf("https://raw.githubusercontent.com/pulumi/pulumi-%s/%s/provider/cmd/pulumi-resource-%s/schema.json", provider, version, provider))
	if err != nil {
		return nil, err
	}
	// Write response body to testdata/aws.json
	return io.ReadAll(res.Body)
}
