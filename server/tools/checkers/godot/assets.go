// File: assets.go
// Purpose: Validates asset file size limits to prevent bloated game builds.
// Scans for image (png, jpg) and audio (wav) files exceeding 2MB threshold.
// Large assets should be compressed or split. This keeps repository size
// manageable and game loading times acceptable. Part of Godot-specific
// quality gates. Reports each oversized file with actual size for easy
// identification and remediation by artists and developers.
// Path: server/tools/checkers/assets.go
// All Rights Reserved. Arc-Pub.

package godot

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/arc-pub/server/tools/core"
)

const maxAssetSizeMB = 2

// Assets checks for oversized files.
type Assets struct {
	Root string
}

// NewAssets creates an Assets checker.
func NewAssets(root string) *Assets {
	return &Assets{Root: root}
}

// Name returns the checker identifier.
func (c *Assets) Name() string {
	return "assets"
}

// Check scans for assets exceeding size limit.
func (c *Assets) Check(_ []core.File) []core.Violation {
	var violations []core.Violation
	exts := map[string]bool{".png": true, ".jpg": true, ".wav": true}

	filepath.Walk(c.Root, func(
		path string, info os.FileInfo, err error,
	) error {
		if err != nil || info.IsDir() {
			return nil
		}

		ext := strings.ToLower(filepath.Ext(path))
		if !exts[ext] {
			return nil
		}

		sizeMB := float64(info.Size()) / (1024 * 1024)
		if sizeMB > maxAssetSizeMB {
			violations = append(violations, core.Violation{
				File:    path,
				Line:    0,
				Rule:    "assets",
				Message: fmt.Sprintf("%.1fMB > %dMB", sizeMB, maxAssetSizeMB),
			})
		}
		return nil
	})

	return violations
}
