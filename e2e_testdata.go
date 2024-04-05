// The following directive is necessary to make the package coherent:

//go:build ignore
// +build ignore

// This program generates contributors.go. It can be invoked by running
// go generate
package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/pulumi/pulumi/pkg/v3/codegen/schema"
	"github.com/pulumi/splitschema"
)

var awsVersion = "v6.29.0"

func main() {
	if err := os.MkdirAll("testdata", 0755); err != nil {
		panic(err)
	}
	bytes, err := downloadAwsSchema()
	if err != nil {
		panic(err)
	}

	var pkg schema.PackageSpec
	err = json.Unmarshal(bytes, &pkg)
	if err != nil {
		panic(err)
	}

	if err = os.WriteFile("testdata/aws.json", bytes, 0644); err != nil {
		panic(err)
	}

	if err := splitschema.WritePackageSpec("testdata/aws", &pkg); err != nil {
		panic(err)
	}
}

func downloadAwsSchema() ([]byte, error) {
	res, err := http.Get(fmt.Sprintf("https://raw.githubusercontent.com/pulumi/pulumi-aws/%s/provider/cmd/pulumi-resource-aws/schema.json", awsVersion))
	if err != nil {
		return nil, err
	}
	// Write response body to testdata/aws.json
	return io.ReadAll(res.Body)
}
