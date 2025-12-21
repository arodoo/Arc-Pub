// File: main.go
// Purpose: Entry point for check-density tool. Validates folder file counts.
// Uses modular architecture. Exits with code 1 if violations found.
// Path: server/tools/cmd/check-density/main.go
// All Rights Reserved. Arc-Pub.

package main

import (
	"os"

	"github.com/arc-pub/server/tools/checkers/code"
	"github.com/arc-pub/server/tools/reporters"
)

func main() {
	checker := code.NewDensity(".")
	violations := checker.Check(nil)

	reporter := reporters.NewConsole()
	reporter.Report(os.Stdout, violations)

	if len(violations) > 0 {
		os.Exit(1)
	}
}
