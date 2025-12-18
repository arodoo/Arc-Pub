// Arc-Pub - Metaverso 2D MMO Social
// Copyright (c) 2024. MIT License.

// check-density: Enforces max files per folder.
package main

import (
	"os"

	"github.com/arc-pub/server/tools/checkers"
	"github.com/arc-pub/server/tools/reporters"
)

func main() {
	checker := checkers.NewDensity(".")
	violations := checker.Check(nil)

	reporter := reporters.NewConsole()
	reporter.Report(os.Stdout, violations)

	if len(violations) > 0 {
		os.Exit(1)
	}
}
