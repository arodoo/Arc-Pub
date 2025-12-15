package checkers

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/arc-pub/server/tools/core"
)

const maxFilesPerDir = 10

// Density checks folder file count.
type Density struct {
	Root string
}

// NewDensity creates a Density checker.
func NewDensity(root string) *Density {
	return &Density{Root: root}
}

// Name returns the checker identifier.
func (c *Density) Name() string {
	return "density"
}

// Check scans directories for overcrowding.
func (c *Density) Check(_ []core.File) []core.Violation {
	var violations []core.Violation
	dirCounts := make(map[string]int)

	filepath.Walk(c.Root, func(
		path string, info os.FileInfo, err error,
	) error {
		if err != nil || info.IsDir() {
			return nil
		}
		if strings.HasPrefix(info.Name(), ".") {
			return nil
		}
		dir := filepath.Dir(path)
		dirCounts[dir]++
		return nil
	})

	for dir, count := range dirCounts {
		if count > maxFilesPerDir {
			violations = append(violations, core.Violation{
				File:    dir,
				Line:    0,
				Rule:    "density",
				Message: fmt.Sprintf("%d files; max %d", count, maxFilesPerDir),
			})
		}
	}

	return violations
}
