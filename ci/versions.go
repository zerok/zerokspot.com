package main

import (
	"context"
	"fmt"
	"os"
	"strings"
)

type Versions struct {
	toolVersions map[string]string
}

func LoadVersions(ctx context.Context) (*Versions, error) {
	_, span := tracer.Start(ctx, "loadVersions")
	defer span.End()
	vs := &Versions{
		toolVersions: make(map[string]string),
	}

	rawToolVersions, err := os.ReadFile(".tool-versions")
	if err != nil {
		return nil, fmt.Errorf("failed to read .tool-versions: %w", err)
	}
	for _, line := range strings.Split(string(rawToolVersions), "\n") {
		elems := strings.SplitN(line, " ", 2)
		if len(elems) < 2 {
			continue
		}
		vs.toolVersions[elems[0]] = elems[1]
	}
	return vs, nil
}

func (vs *Versions) GoVersion() string {
	return vs.toolVersions["golang"]
}

func (vs *Versions) GoImage() string {
	return fmt.Sprintf("golang:%s", vs.GoVersion())
}

func (vs *Versions) HugoImage() string {
	tool := vs.toolVersions["hugo"]
	if strings.HasPrefix(tool, "extended_") {
		return fmt.Sprintf("klakegg/hugo:%s-ext-ubuntu", strings.TrimPrefix(tool, "extended_"))
	}
	return fmt.Sprintf("klakegg/hugo:%s-ext-ubuntu", tool)
}
