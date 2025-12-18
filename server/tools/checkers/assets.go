// Arc-Pub - Metaverso 2D MMO Social
// Copyright (c) 2024. MIT License.

package checkers

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
