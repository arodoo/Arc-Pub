// Arc-Pub - Metaverso 2D MMO Social
// Copyright (c) 2024. MIT License.

// check-orphans: Finds unreferenced .gd/.png files.
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
