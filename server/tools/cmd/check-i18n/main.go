// File: main.go
// Purpose: Entry point for check-i18n quality tool. Scans GDScript files for
// hardcoded UI strings that should use tr() translation function. Ensures
// all user-facing text can be localized. Exits with code 1 if violations
// found. Part of automated internationalization quality gates.
// Path: server/tools/cmd/check-i18n/main.go
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
	files, err := scanner.Scan(".", []string{".gd"})
	if err != nil {
		os.Stderr.WriteString("scan error: " + err.Error() + "\n")
		os.Exit(1)
	}

	checker := checkers.NewI18n()
	violations := checker.Check(files)

	reporter := reporters.NewConsole()
	reporter.Report(os.Stdout, violations)

	if len(violations) > 0 {
		os.Exit(1)
	}
}
