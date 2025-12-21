// File: main.go
// Purpose: Entry point for check-sqlc tool. Validates sqlc sync.
// Uses modular architecture. Exits with code 1 if violations found.
// Path: server/tools/cmd/check-sqlc/main.go
// All Rights Reserved. Arc-Pub.

package main

import (
	"os"

	"github.com/arc-pub/server/tools/checkers/sync"
	"github.com/arc-pub/server/tools/reporters"
)

func main() {
	checker := sync.NewSQLCSync(
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
