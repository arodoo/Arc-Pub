// File: orphans.go
// Purpose: Detects orphaned files (dead code) not referenced by any scene.
// Builds a reference map from all .tscn files by extracting res:// paths,
// then scans for .gd scripts and .png images not in that map. Helps maintain
// a clean codebase by identifying unused assets and scripts for removal.
// Part of project quality gates for scene integrity. Reports each orphaned
// file path for developer review and potential cleanup decision.
// Path: server/tools/checkers/orphans.go
// All Rights Reserved. Arc-Pub.

package sync

import (
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/arc-pub/server/tools/core"
)

var resRefRe = regexp.MustCompile(`res://[^"'\s]+`)

// Orphans checks for unreferenced files.
type Orphans struct {
	Root string
}

// NewOrphans creates an Orphans checker.
func NewOrphans(root string) *Orphans {
	return &Orphans{Root: root}
}

// Name returns the checker identifier.
func (c *Orphans) Name() string {
	return "orphans"
}

// Check finds files not referenced by any .tscn.
func (c *Orphans) Check(files []core.File) []core.Violation {
	referenced := make(map[string]bool)
	var candidates []string

	for _, f := range files {
		if !strings.HasSuffix(f.Path, ".tscn") {
			continue
		}
		for _, line := range f.Lines {
			matches := resRefRe.FindAllString(line, -1)
			for _, m := range matches {
				rel := strings.TrimPrefix(m, "res://")
				referenced[rel] = true
			}
		}
	}

	exts := map[string]bool{".gd": true, ".png": true}
	filepath.Walk(c.Root, func(
		path string, info os.FileInfo, err error,
	) error {
		if err != nil || info.IsDir() {
			return nil
		}
		ext := filepath.Ext(path)
		if exts[ext] {
			candidates = append(candidates, path)
		}
		return nil
	})

	var violations []core.Violation
	for _, path := range candidates {
		relPath, _ := filepath.Rel(c.Root, path)
		relPath = filepath.ToSlash(relPath)

		if !referenced[relPath] {
			violations = append(violations, core.Violation{
				File:    path,
				Line:    0,
				Rule:    "orphans",
				Message: "not referenced by any .tscn",
			})
		}
	}

	return violations
}
