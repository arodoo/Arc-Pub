// File: main.go
// Purpose: Entry point for check-visuals quality tool. Scans GDScript and
// scene files for hardcoded Color() constructors and HEX codes. Colors should
// use Palette singleton for consistent theming. Exits with code 1 if violations
// found. Part of automated visual consistency quality gates.
// Path: server/tools/cmd/check-visuals/main.go
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
	files, err := scanner.Scan(".", []string{".gd", ".tscn"})
	if err != nil {
		os.Stderr.WriteString("scan error: " + err.Error() + "\n")
		os.Exit(1)
	}

	checker := checkers.NewVisuals()
	violations := checker.Check(files)

	reporter := reporters.NewConsole()
	reporter.Report(os.Stdout, violations)

	if len(violations) > 0 {
		os.Exit(1)
	}
}
