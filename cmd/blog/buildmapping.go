package main

import (
	"log/slog"

	"github.com/spf13/cobra"
	"gitlab.com/zerok/zerokspot.com/pkg/contentmapping"
	"go.opentelemetry.io/otel/codes"
)

var buildMappingCmd = &cobra.Command{
	Use: "build-mapping",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := findParentTrace(cmd.Context())
		ctx, span := tracer.Start(ctx, "cmd:build-mapping")
		defer span.End()
		outputFile := cmd.Flag("output").Value.String()
		slog.InfoContext(ctx, "Writing mapping file", slog.String("output", outputFile))
		mapping, err := contentmapping.BuildMapping("public")
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, "Failed to build mapping")
			return err
		}
		if err := contentmapping.SaveToFile(outputFile, mapping); err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, "Failed to save mapping")
			return err
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(buildMappingCmd)
	buildMappingCmd.Flags().String("output", "public/.mapping.json.xz", "Output file path")
}
