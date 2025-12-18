// File: main.go
// Purpose: Entry point for check-assets quality tool. Scans Godot project for
// image and audio assets exceeding 2MB size limit. Large assets impact game
// loading times and repository size. Exits with code 1 if violations found.
// Part of automated asset management quality gates.
// Path: server/tools/cmd/check-assets/main.go
// All Rights Reserved. Arc-Pub.

package main

import (
	"flag"
	"os"

	"github.com/arc-pub/server/tools/checkers"
	"github.com/arc-pub/server/tools/reporters"
)

func main() {
	root := flag.String("root", "..", "Godot project root")
	flag.Parse()

	checker := checkers.NewAssets(*root)
	violations := checker.Check(nil)

	reporter := reporters.NewConsole()
	reporter.Report(os.Stdout, violations)

	if len(violations) > 0 {
		os.Exit(1)
	}
}
