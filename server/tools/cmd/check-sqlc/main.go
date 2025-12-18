// Arc-Pub - Metaverso 2D MMO Social
// Copyright (c) 2024. MIT License.

// check-sqlc: Verifies sqlc generated code is fresh.
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
