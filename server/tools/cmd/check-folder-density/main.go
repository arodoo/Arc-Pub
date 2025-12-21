// File: main.go
// Purpose: Entry point for check-folder-density quality tool. Validates max 6
// files per directory. Uses modular architecture. Exits with code 1 if found.
// Path: server/tools/cmd/check-folder-density/main.go
// All Rights Reserved. Arc-Pub.

package main

import (
	"flag"
	"os"

	"github.com/arc-pub/server/tools/checkers/code"
	"github.com/arc-pub/server/tools/reporters"
	"github.com/arc-pub/server/tools/scanners"
)

func main() {
	root := flag.String("root", ".", "Root directory to scan")
	flag.Parse()

	scanner := scanners.NewFileScanner()
	files, err := scanner.Scan(*root, []string{".go", ".gd"})
	if err != nil {
		os.Stderr.WriteString("scan error: " + err.Error() + "\n")
		os.Exit(1)
	}

	checker := code.NewFolderDensity()
	violations := checker.Check(files)

	reporter := reporters.NewConsole()
	reporter.Report(os.Stdout, violations)

	if len(violations) > 0 {
		os.Exit(1)
	}
}
