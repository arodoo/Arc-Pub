// Arc-Pub - Metaverso 2D MMO Social
// Copyright (c) 2024. MIT License.

package checkers

import (
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/arc-pub/server/tools/core"
)

var extResourceRe = regexp.MustCompile(`ext_resource\s+.*path="(res://[^"]+)"`)

// Scenes checks for missing external resources in .tscn.
type Scenes struct {
	Root string
}

// NewScenes creates a Scenes checker.
func NewScenes(root string) *Scenes {
	return &Scenes{Root: root}
}

// Name returns the checker identifier.
func (c *Scenes) Name() string {
	return "scenes"
}

// Check verifies ext_resource paths exist.
func (c *Scenes) Check(files []core.File) []core.Violation {
	var violations []core.Violation

	for _, f := range files {
		if !strings.HasSuffix(f.Path, ".tscn") {
			continue
		}

		for i, line := range f.Lines {
			matches := extResourceRe.FindStringSubmatch(line)
			if len(matches) < 2 {
				continue
			}

			resPath := matches[1]
			diskPath := c.resPathToDisk(resPath)

			if _, err := os.Stat(diskPath); os.IsNotExist(err) {
				violations = append(violations, core.Violation{
					File:    f.Path,
					Line:    i + 1,
					Rule:    "scenes",
					Message: "missing: " + resPath,
				})
			}
		}
	}

	return violations
}

func (c *Scenes) resPathToDisk(resPath string) string {
	rel := strings.TrimPrefix(resPath, "res://")
	return filepath.Join(c.Root, rel)
}
