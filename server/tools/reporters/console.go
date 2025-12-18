// File: console.go
// Purpose: Implements the Reporter interface for human-readable console
// output. Formats violations with file path, line number, rule name, and
// message using Unicode symbols for visual clarity (checkmarks, crosses).
// Outputs summary count at the end. Designed for developer terminal usage
// during local development. Returns nil error as stdout writes rarely fail.
// Path: server/tools/reporters/console.go
// All Rights Reserved. Arc-Pub.

package reporters

import (
	"fmt"
	"io"

	"github.com/arc-pub/server/tools/core"
)

// Console outputs violations to terminal with colors.
type Console struct{}

// NewConsole creates a Console reporter.
func NewConsole() *Console {
	return &Console{}
}

// Report writes violations to w in human-readable format.
func (c *Console) Report(
	w io.Writer,
	violations []core.Violation,
) error {
	for _, v := range violations {
		fmt.Fprintf(w, "✗ %s:%d [%s]\n", v.File, v.Line, v.Rule)
		fmt.Fprintf(w, "  %s\n\n", v.Message)
	}

	if len(violations) > 0 {
		fmt.Fprintf(w, "Found %d violation(s)\n", len(violations))
	} else {
		fmt.Fprintln(w, "✓ No violations found")
	}
	return nil
}
