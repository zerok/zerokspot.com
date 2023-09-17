package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/ulikunitz/xz"
)

var buildMappingCmd = &cobra.Command{
	Use: "build-mapping",
	RunE: func(cmd *cobra.Command, args []string) error {
		mapping, err := buildMapping("public")
		if err != nil {
			return fmt.Errorf("failed to build mapping: %w", err)
		}
		fp, err := os.OpenFile(filepath.Join("public", ".mapping.json.xz"), os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
		if err != nil {
			return fmt.Errorf("failed to open %s for writing: %w", ".mapping.json.xz", err)
		}
		defer fp.Close()
		w, err := xz.NewWriter(fp)
		if err != nil {
			return fmt.Errorf("failed to create new xz writer: %w", err)
		}
		defer w.Close()
		if err := json.NewEncoder(w).Encode(mapping); err != nil {
			return fmt.Errorf("failed to write mapping file: %w", err)
		}
		return nil
	},
}

func init() {
	RootCmd.AddCommand(buildMappingCmd)
}
