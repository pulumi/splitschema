// Copyright 2024, Pulumi Corporation.

package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/pulumi/splitschema"
	"github.com/spf13/cobra"
)

var (
	mergeSource  string
	mergeDest    string
	mergeCompact bool
)

var mergeCmd = &cobra.Command{
	Use:   "merge",
	Short: "Merge split schema files into a single schema file",
	RunE: func(cmd *cobra.Command, args []string) error {
		sourcePackage := splitschema.NewLocalPartialPackage(mergeSource)
		pkg, err := sourcePackage.ReadPackageSpec()
		if err != nil {
			return fmt.Errorf("read package spec: %w", err)
		}
		var pkgBytes []byte
		if mergeCompact {
			pkgBytes, err = json.Marshal(pkg)
		} else {
			pkgBytes, err = json.MarshalIndent(pkg, "", "  ")
		}
		if err != nil {
			return fmt.Errorf("marshal package spec: %w", err)
		}
		err = os.WriteFile(mergeDest, pkgBytes, 0644)
		if err != nil {
			return fmt.Errorf("write package spec: %w", err)
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(mergeCmd)
	mergeCmd.Flags().StringVarP(&mergeSource, "source", "s", ".", "Source directory containing split schema files")
	mergeCmd.Flags().StringVarP(&mergeDest, "dest", "d", "schema.json", "Destination file to write merged schema")
	mergeCmd.Flags().BoolVarP(&mergeCompact, "compact", "c", false, "Compact the merged schema")
}
