// File: limits.go
// Purpose: Validates file and line length limits following cognitive load
// principles. Enforces maximum 120 lines per file to encourage focused,
// single-responsibility modules. Enforces maximum 120 characters per line
// for readability without horizontal scrolling. These limits are slightly
// under CI thresholds to provide buffer. Reports violations for each file
// exceeding line count and each individual line exceeding character count.
// Path: server/tools/checkers/limits.go
// All Rights Reserved. Arc-Pub.

package checkers

import (
	"fmt"

	"github.com/arc-pub/server/tools/core"
)

const (
	maxLines      = 120
	maxLineLength = 120
)

// Limits checks file and line length constraints.
type Limits struct{}

// NewLimits creates a Limits checker.
func NewLimits() *Limits {
	return &Limits{}
}

// Name returns the checker identifier.
func (l *Limits) Name() string {
	return "limits"
}

// Check scans files for length violations.
func (l *Limits) Check(files []core.File) []core.Violation {
	var violations []core.Violation

	for _, f := range files {
		if len(f.Lines) > maxLines {
			violations = append(violations, core.Violation{
				File:    f.Path,
				Line:    len(f.Lines),
				Rule:    "file-length",
				Message: fmt.Sprintf("exceeds %d lines", maxLines),
			})
		}

		for i, line := range f.Lines {
			if len(line) > maxLineLength {
				violations = append(violations, core.Violation{
					File:    f.Path,
					Line:    i + 1,
					Rule:    "line-length",
					Message: fmt.Sprintf("exceeds %d chars", maxLineLength),
				})
			}
		}
	}

	return violations
}
