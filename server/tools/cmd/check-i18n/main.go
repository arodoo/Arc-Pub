// File: main.go
// Purpose: Entry point for check-i18n quality tool. Validates no hardcoded
// text in Godot scripts. Exits with code 1 if violations found.
// Path: server/tools/cmd/check-i18n/main.go
// All Rights Reserved. Arc-Pub.

package main

import (
	"flag"
	"os"

	"github.com/arc-pub/server/tools/checkers/godot"
	"github.com/arc-pub/server/tools/reporters"
	"github.com/arc-pub/server/tools/scanners"
)

func main() {
	flag.Parse()
	root := "."
	if flag.NArg() > 0 {
		root = flag.Arg(0)
	}

	scanner := scanners.NewFileScanner()
	files, err := scanner.Scan(root, []string{".gd"})
	if err != nil {
		os.Stderr.WriteString("scan error: " + err.Error() + "\n")
		os.Exit(1)
	}

	checker := godot.NewI18n()
	violations := checker.Check(files)

	reporter := reporters.NewConsole()
	reporter.Report(os.Stdout, violations)

	if len(violations) > 0 {
		os.Exit(1)
	}
}
