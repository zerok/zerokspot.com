package main

import (
	"context"
	"fmt"
	"os"
	"strings"
)

type Versions struct {
	images map[string]string
}

func LoadVersions(ctx context.Context) (*Versions, error) {
	vs := &Versions{
		images: make(map[string]string),
	}

	dockerfileContent, err := os.ReadFile("Dockerfile")
	if err != nil {
		return nil, fmt.Errorf("failed to read Dockerfile: %w", err)
	}
	for _, line := range strings.Split(string(dockerfileContent), "\n") {
		if !strings.HasPrefix(line, "FROM ") {
			continue
		}
		elems := strings.SplitN(line, " ", 3)
		if len(elems) < 2 {
			continue
		}
		image := elems[1]
		elems = strings.SplitN(image, ":", 2)
		vs.images[elems[0]] = elems[1]
	}
	return vs, nil
}

func (vs *Versions) GoImage() string {
	return fmt.Sprintf("golang:%s", vs.images["golang"])
}

func (vs *Versions) AlpineImage() string {
	return fmt.Sprintf("alpine:%s", vs.images["alpine"])
}

func (vs *Versions) UbuntuImage() string {
	return fmt.Sprintf("ubuntu:%s", vs.images["ubuntu"])
}

func (vs *Versions) WebmentiondImage() string {
	return fmt.Sprintf("zerok/webmentiond:%s", vs.images["zerok/webmentiond"])
}
