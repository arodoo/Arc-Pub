// File: main.go
// Purpose: Entry point for check-density quality tool. Scans directories and
// validates folder file counts do not exceed 10 files per folder. Follows
// cognitive load principles for navigable project structure. Exits with
// code 1 if violations found. Part of automated quality gates.
// Path: server/tools/cmd/check-density/main.go
// All Rights Reserved. Arc-Pub.

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
