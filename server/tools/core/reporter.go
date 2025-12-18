// File: reporter.go
// Purpose: Defines the Reporter interface for outputting quality check
// violations. Abstracts the output format enabling different implementations
// like console output, JSON for CI systems, or HTML reports. Uses io.Writer
// for flexible output destinations (stdout, files, buffers). Follows Open/
// Closed Principle allowing new output formats without modifying existing
// checkers. Returns error for handling write failures gracefully.
// Path: server/tools/core/reporter.go
// All Rights Reserved. Arc-Pub.

package core

import "io"

// Reporter outputs violations to a destination.
type Reporter interface {
	Report(w io.Writer, violations []Violation) error
}
