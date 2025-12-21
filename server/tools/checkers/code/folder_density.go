// File: folder_density.go
// Purpose: Validates folder structure following SOLID principles - max 6 files
// per directory to enforce single-responsibility modules. Too many files in
// one folder indicates need for restructuring into sub-packages. Excludes only
// auto-generated and external directories.
// Path: server/tools/checkers/folder_density.go
// All Rights Reserved. Arc-Pub.

package code

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/arc-pub/server/tools/core"
)

const maxFilesPerFolder = 6

// Directories to exclude from folder density checks.
var excludedDirs = []string{
	"sqlc",
	"node_modules",
	"vendor",
	".git",
	".import",
}

// FolderDensity checks for too many files in a single directory.
type FolderDensity struct{}

// NewFolderDensity creates a FolderDensity checker.
func NewFolderDensity() *FolderDensity {
	return &FolderDensity{}
}

// Name returns the checker identifier.
func (f *FolderDensity) Name() string {
	return "folder-density"
}

// Check counts files per folder.
func (f *FolderDensity) Check(files []core.File) []core.Violation {
	folderCounts := make(map[string]int)

	for _, file := range files {
		dir := filepath.Dir(file.Path)
		if isExcludedDir(dir) {
			continue
		}
		folderCounts[dir]++
	}

	var violations []core.Violation
	for dir, count := range folderCounts {
		if count > maxFilesPerFolder {
			violations = append(violations, core.Violation{
				File:    dir,
				Line:    0,
				Rule:    "folder-density",
				Message: fmt.Sprintf("has %d files, max %d", count, maxFilesPerFolder),
			})
		}
	}

	return violations
}

func isExcludedDir(dir string) bool {
	lower := strings.ToLower(dir)
	for _, excluded := range excludedDirs {
		if strings.Contains(lower, excluded) {
			return true
		}
	}
	return false
}
