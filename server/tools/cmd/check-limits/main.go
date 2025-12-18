// File: main.go
// Purpose: Entry point for check-limits quality tool. Scans Go and GDScript
// files for line count (max 120) and line length (max 120 chars) violations.
// Uses modular scanner/checker/reporter pattern. Exits with code 1 if any
// violations found for CI integration. Part of automated quality gates.
// Path: server/tools/cmd/check-limits/main.go
// All Rights Reserved. Arc-Pub.

package main

import (
	"os"

	"github.com/arc-pub/server/tools/checkers"
	"github.com/arc-pub/server/tools/reporters"
	"github.com/arc-pub/server/tools/scanners"
)

func main() {
	scanner := scanners.NewFileScanner()
	files, err := scanner.Scan(".", []string{".go", ".gd"})
	if err != nil {
		os.Stderr.WriteString("scan error: " + err.Error() + "\n")
		os.Exit(1)
	}

	checker := checkers.NewLimits()
	violations := checker.Check(files)

	reporter := reporters.NewConsole()
	reporter.Report(os.Stdout, violations)

	if len(violations) > 0 {
		os.Exit(1)
	}
}
