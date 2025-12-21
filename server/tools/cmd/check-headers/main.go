// File: main.go
// Purpose: Entry point for check-headers quality tool. Validates file headers.
// Uses modular architecture. Exits with code 1 if violations found.
// Path: server/tools/cmd/check-headers/main.go
// All Rights Reserved. Arc-Pub.

package main

import (
	"os"

	"github.com/arc-pub/server/tools/checkers/code"
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

	checker := code.NewHeaders()
	violations := checker.Check(files)

	reporter := reporters.NewConsole()
	reporter.Report(os.Stdout, violations)

	if len(violations) > 0 {
		os.Exit(1)
	}
}
