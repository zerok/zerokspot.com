package resizer

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

type Resizer interface {
	Resize(ctx context.Context, path string) error
	GetProfile(profileName string) *Profile
	GetProfiles() *Profiles
}

type MagickResizer struct {
	dataPath string
	profiles *Profiles
}

func NewMagickResizer(dataPath string, profiles *Profiles) *MagickResizer {
	return &MagickResizer{
		dataPath: dataPath,
		profiles: profiles,
	}
}

func (mr *MagickResizer) Resize(ctx context.Context, path string) error {
	sourcePath := filepath.Join(mr.dataPath, "photos", path)
	outputFolder := filepath.Join(mr.dataPath, "photos", "resized", path)
	if err := os.MkdirAll(outputFolder, 0700); err != nil {
		return fmt.Errorf("failed to created folder for resized files: %w", err)
	}
	for name, prof := range mr.profiles.p {
		targetPath := filepath.Join(outputFolder, fmt.Sprintf("%s.%s", name, prof.Format))
		if err := exec.CommandContext(ctx, "magick", "convert", sourcePath, "-auto-orient", "-strip", "-resize", fmt.Sprintf("%dx", prof.Width)+">", targetPath).Run(); err != nil {
			return err
		}
	}
	return nil
}

func (mr *MagickResizer) GetProfile(pname string) *Profile {
	p, ok := mr.GetProfiles().Get(pname)
	if ok {
		return &p
	}
	return nil
}

func (mr *MagickResizer) GetProfiles() *Profiles {
	return mr.profiles
}
