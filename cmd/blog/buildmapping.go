package main

import (
	"github.com/spf13/cobra"
	"gitlab.com/zerok/zerokspot.com/pkg/contentmapping"
)

var outputFile string

var buildMappingCmd = &cobra.Command{
	Use: "build-mapping",
	RunE: func(cmd *cobra.Command, args []string) error {
		mapping, err := contentmapping.BuildMapping("public")
		if err != nil {
			return err
		}
		return contentmapping.SaveToFile(outputFile, mapping)
	},
}

func init() {
	rootCmd.AddCommand(buildMappingCmd)
	buildMappingCmd.Flags().StringVar(&outputFile, "output", "public/.mapping.json.xz", "Output file path")
}
