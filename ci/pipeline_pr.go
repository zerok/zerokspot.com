package main

import (
	"os"

	"dagger.io/dagger"
	"github.com/spf13/cobra"
)

var pipelinePRCmd = &cobra.Command{
	Use: "pr",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := cmd.Context()
		ctx, span := tracer.Start(ctx, "pr")
		defer span.End()
		dc, err := dagger.Connect(ctx, dagger.WithLogOutput(os.Stderr))
		if err != nil {
			failSpan(ctx, span, "Failed to connect to Dagger", err)
			return err
		}
		defer dc.Close()
		return build(ctx, dc, versions, false)
	},
}

func init() {
	pipelinesCmd.AddCommand(pipelinePRCmd)
}
