// File: main.go
// Purpose: Entry point for check-sqlc quality tool. Verifies sqlc generated
// code is synchronized with database schema source files. Compares file
// timestamps to detect stale code. Exits with code 1 if regeneration needed.
// Part of automated contract sync quality gates.
// Path: server/tools/cmd/check-sqlc/main.go
// All Rights Reserved. Arc-Pub.

package main

import (
	"os"

	"github.com/arc-pub/server/tools/checkers"
	"github.com/arc-pub/server/tools/reporters"
)

func main() {
	checker := checkers.NewSQLCSync(
		"db/schema.sql",
		"internal/infra/postgres/sqlc",
	)
	violations := checker.Check(nil)

	reporter := reporters.NewConsole()
	reporter.Report(os.Stdout, violations)

	if len(violations) > 0 {
		os.Exit(1)
	}
}
