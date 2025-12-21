// File: headers.go
// Purpose: Validates proper file headers in Go and GDScript files. Headers must
// include filename with extension, purpose description (minimum 100 characters
// explaining functionality), file path, and copyright notice with all rights
// reserved. This ensures consistent documentation across the entire codebase.
// Path: server/tools/checkers/headers.go
// All Rights Reserved. Arc-Pub.

package code

import (
	"regexp"
	"strings"

	"github.com/arc-pub/server/tools/core"
)

var (
	fileRe   = regexp.MustCompile(`(?i)File:\s*\S+`)
	pathRe   = regexp.MustCompile(`(?i)Path:\s*\S+`)
	rightsRe = regexp.MustCompile(`(?i)All Rights Reserved`)
)

// Headers checks for complete file headers.
type Headers struct{}

// NewHeaders creates a Headers checker.
func NewHeaders() *Headers {
	return &Headers{}
}

// Name returns the checker identifier.
func (c *Headers) Name() string {
	return "headers"
}

// Check verifies files have complete headers.
func (c *Headers) Check(files []core.File) []core.Violation {
	var violations []core.Violation

	for _, f := range files {
		if len(f.Lines) < 5 {
			continue
		}

		header := strings.Join(f.Lines[:min(10, len(f.Lines))], " ")

		// Skip generated files (sqlc, protobuf, etc)
		if strings.Contains(header, "DO NOT EDIT") {
			continue
		}

		missing := c.checkHeader(header)

		if len(missing) > 0 {
			violations = append(violations, core.Violation{
				File:    f.Path,
				Line:    1,
				Rule:    "headers",
				Message: "missing: " + strings.Join(missing, ", "),
			})
		}
	}

	return violations
}

func (c *Headers) checkHeader(header string) []string {
	var missing []string

	if !fileRe.MatchString(header) {
		missing = append(missing, "File:")
	}
	if !c.hasPurpose(header) {
		missing = append(missing, "Purpose (100+ chars)")
	}
	if !pathRe.MatchString(header) {
		missing = append(missing, "Path:")
	}
	if !rightsRe.MatchString(header) {
		missing = append(missing, "All Rights Reserved")
	}

	return missing
}

func (c *Headers) hasPurpose(header string) bool {
	idx := strings.Index(strings.ToLower(header), "purpose:")
	if idx == -1 {
		return false
	}

	rest := header[idx+8:]
	pathIdx := strings.Index(strings.ToLower(rest), "path:")
	if pathIdx == -1 {
		return len(rest) >= 100
	}
	return pathIdx >= 100
}
