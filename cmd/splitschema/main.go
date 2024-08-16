// Copyright 2024, Pulumi Corporation.

package main

import (
	"os"

	"github.com/spf13/cobra"
)

func main() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

var rootCmd = &cobra.Command{
	Use:   "splitschema",
	Short: "Splits schema json into component files",
	Long: `Makes Pulumi schema JSON files easier to work with by splitting them into
separate directories and files for each resource, provider, and type.`,
}

func init() {
}
