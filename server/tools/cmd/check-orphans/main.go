// File: main.go
// Purpose: Entry point for check-orphans quality tool. Detects unreferenced
// GDScript and PNG files not used by any scene. Helps identify dead code and
// unused assets for cleanup. Exits with code 1 if violations found. Part of
// automated project hygiene quality gates.
// Path: server/tools/cmd/check-orphans/main.go
// All Rights Reserved. Arc-Pub.

package main

import (
	"flag"
	"os"

	"github.com/arc-pub/server/tools/checkers"
	"github.com/arc-pub/server/tools/reporters"
	"github.com/arc-pub/server/tools/scanners"
)

func main() {
	root := flag.String("root", "..", "Godot project root")
	flag.Parse()

	scanner := scanners.NewFileScanner()
	files, err := scanner.Scan(*root, []string{".tscn"})
	if err != nil {
		os.Stderr.WriteString("scan error: " + err.Error() + "\n")
		os.Exit(1)
	}

	checker := checkers.NewOrphans(*root)
	violations := checker.Check(files)

	reporter := reporters.NewConsole()
	reporter.Report(os.Stdout, violations)

	if len(violations) > 0 {
		os.Exit(1)
	}
}
