package main

import (
	"os"

	"dagger.io/dagger"
	"github.com/spf13/cobra"
)

var pipelineMainCmd = &cobra.Command{
	Use: "main",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := cmd.Context()
		ctx, span := tracer.Start(ctx, "main")
		defer span.End()
		dc, err := dagger.Connect(ctx, dagger.WithLogOutput(os.Stderr))
		if err != nil {
			failSpan(ctx, span, "Failed to connect to Dagger", err)
			return err
		}
		defer dc.Close()
		return build(ctx, dc, versions, true)
	},
}

func init() {
	pipelinesCmd.AddCommand(pipelineMainCmd)
}
