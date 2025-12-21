// File: main.go
// Purpose: Entry point for check-assets tool. Validates asset file sizes.
// Uses modular architecture. Exits with code 1 if violations found.
// Path: server/tools/cmd/check-assets/main.go
// All Rights Reserved. Arc-Pub.

package main

import (
	"flag"
	"os"

	"github.com/arc-pub/server/tools/checkers/godot"
	"github.com/arc-pub/server/tools/reporters"
)

func main() {
	flag.Parse()
	root := "."
	if flag.NArg() > 0 {
		root = flag.Arg(0)
	}

	checker := godot.NewAssets(root)
	violations := checker.Check(nil)

	reporter := reporters.NewConsole()
	reporter.Report(os.Stdout, violations)

	if len(violations) > 0 {
		os.Exit(1)
	}
}
