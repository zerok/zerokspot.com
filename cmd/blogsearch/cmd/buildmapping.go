package cmd

import (
	"encoding/json"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/ulikunitz/xz"
)

var buildMappingCmd = &cobra.Command{
	Use: "build-mapping",
	Run: func(cmd *cobra.Command, args []string) {
		mapping, err := buildMapping("public")
		if err != nil {
			logger.Fatal().Err(err).Msg("Failed to build mapping")
		}
		fp, err := os.OpenFile(filepath.Join("public", ".mapping.json.xz"), os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
		if err != nil {
			logger.Fatal().Err(err).Msgf("Failed to open %s for writing", ".mapping.json.xz")
		}
		defer fp.Close()
		w, err := xz.NewWriter(fp)
		if err != nil {
			logger.Fatal().Err(err).Msgf("Failed to create new xz writer")
		}
		defer w.Close()
		if err := json.NewEncoder(w).Encode(mapping); err != nil {
			logger.Fatal().Err(err).Msgf("Failed to write mapping file")
		}
	},
}

func init() {
	RootCmd.AddCommand(buildMappingCmd)
}
