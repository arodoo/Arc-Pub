// Arc-Pub - Metaverso 2D MMO Social
// Copyright (c) 2024. MIT License.

// Package core provides shared interfaces for quality tools.
package core

// File represents a scanned file with metadata.
type File struct {
	Path    string
	Content []byte
	Lines   []string
}

// Scanner defines file discovery operations.
type Scanner interface {
	Scan(root string, extensions []string) ([]File, error)
}
