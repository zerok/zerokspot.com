package main

import (
	"io/fs"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
	"gitlab.com/zerok/zerokspot.com/pkg/resizer"
)

func generatePhotoProfiles() *resizer.Profiles {
	profiles := resizer.NewProfiles()
	profiles.Add("1024", resizer.Profile{
		Width:  1024,
		Format: resizer.FormatJPEG,
	})
	profiles.Add("800", resizer.Profile{
		Width:  800,
		Format: resizer.FormatJPEG,
	})
	profiles.Add("400", resizer.Profile{
		Width:  400,
		Format: resizer.FormatJPEG,
	})
	return profiles
}

func generateResizePhotosCmd() *cobra.Command {
	var dataFolder string
	cmd := &cobra.Command{
		Use: "resize-photos",
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := findParentTrace(cmd.Context())
			ctx, span := tracer.Start(ctx, "resize-photos")
			defer span.End()
			profiles := generatePhotoProfiles()
			r := resizer.NewMagickResizer(dataFolder, profiles)
			photosFolder := filepath.Join(dataFolder, "photos")
			filepath.Walk(photosFolder, func(path string, finfo fs.FileInfo, err error) error {
				if path == filepath.Join(photosFolder, "resized") {
					return filepath.SkipDir
				}
				if finfo.IsDir() {
					return nil
				}
				relPath := strings.TrimPrefix(path, photosFolder+"/")
				return r.Resize(ctx, relPath)
			})
			return nil
		},
	}
	cmd.Flags().StringVar(&dataFolder, "data-folder", "./data", "Folder where files etc. are stored")
	return cmd
}
