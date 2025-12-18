// File: scanner.go
// Purpose: Defines the core Scanner interface and File struct for quality
// tool file discovery operations. Scanner abstracts file system traversal
// enabling different implementations (filesystem, git, embedded). File struct
// holds path, raw content bytes, and pre-split lines for efficient processing
// by checkers. Follows Interface Segregation Principle from SOLID for clean
// abstractions that checkers depend on without coupling to implementation.
// Path: server/tools/core/scanner.go
// All Rights Reserved. Arc-Pub.

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
