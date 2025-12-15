// Package scanners provides file scanning implementations.
package scanners

import (
	"bufio"
	"bytes"
	"os"
	"path/filepath"
	"strings"

	"github.com/arc-pub/server/tools/core"
)

// FileScanner scans directories for files by extension.
type FileScanner struct{}

// NewFileScanner creates a FileScanner instance.
func NewFileScanner() *FileScanner {
	return &FileScanner{}
}

// Scan walks root and returns files matching extensions.
func (s *FileScanner) Scan(
	root string,
	extensions []string,
) ([]core.File, error) {
	var files []core.File

	extMap := make(map[string]bool)
	for _, ext := range extensions {
		extMap[ext] = true
	}

	err := filepath.Walk(root, func(
		path string, info os.FileInfo, err error,
	) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}

		ext := filepath.Ext(path)
		if !extMap[ext] {
			return nil
		}

		content, err := os.ReadFile(path)
		if err != nil {
			return err
		}

		lines := splitLines(content)
		files = append(files, core.File{
			Path:    path,
			Content: content,
			Lines:   lines,
		})
		return nil
	})

	return files, err
}

func splitLines(content []byte) []string {
	var lines []string
	scanner := bufio.NewScanner(bytes.NewReader(content))
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}

// ScanDirs returns directory paths from root.
func ScanDirs(root string) ([]string, error) {
	var dirs []string
	err := filepath.Walk(root, func(
		path string, info os.FileInfo, err error,
	) error {
		if err != nil {
			return err
		}
		if info.IsDir() && !strings.HasPrefix(info.Name(), ".") {
			dirs = append(dirs, path)
		}
		return nil
	})
	return dirs, err
}
