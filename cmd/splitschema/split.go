// Copyright 2024, Pulumi Corporation.

package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/pulumi/pulumi/pkg/v3/codegen/schema"
	"github.com/pulumi/splitschema"
	"github.com/spf13/cobra"
)

var (
	splitSource string
	splitDest   string
)

var splitCmd = &cobra.Command{
	Use:   "split schema",
	Short: "Split a schema file into component files",
	Long: `Split a schema file into component files. This command will read a schema file
and split it into separate files for each resource, provider, and type.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		pkgBytes, err := os.ReadFile(splitSource)
		if err != nil {
			return fmt.Errorf("read source file: %w", err)
		}
		var pkg schema.PackageSpec
		err = json.Unmarshal(pkgBytes, &pkg)
		if err != nil {
			return fmt.Errorf("unmarshal source package spec: %w", err)
		}
		// Ensure destination exists
		err = os.MkdirAll(splitDest, 0755)
		if err != nil {
			return fmt.Errorf("create destination directory: %w", err)
		}
		err = splitschema.WritePackageSpec(splitDest, &pkg)
		if err != nil {
			return fmt.Errorf("write split package spec: %w", err)
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(splitCmd)
	splitCmd.Flags().StringVarP(&splitSource, "source", "s", ".", "Source schema file to split")
	splitCmd.Flags().StringVarP(&splitDest, "dest", "d", "schema.json", "Destination directory to write split schema")
}
