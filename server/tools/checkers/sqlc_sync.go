package checkers

import (
	"os"
	"path/filepath"

	"github.com/arc-pub/server/tools/core"
)

// SQLCSync checks if generated code is newer than schema.
type SQLCSync struct {
	SchemaPath    string
	GeneratedPath string
}

// NewSQLCSync creates a SQLCSync checker.
func NewSQLCSync(schema, generated string) *SQLCSync {
	return &SQLCSync{
		SchemaPath:    schema,
		GeneratedPath: generated,
	}
}

// Name returns the checker identifier.
func (c *SQLCSync) Name() string {
	return "sqlc"
}

// Check verifies generated code is up to date.
func (c *SQLCSync) Check(_ []core.File) []core.Violation {
	var violations []core.Violation

	schemaInfo, err := os.Stat(c.SchemaPath)
	if err != nil {
		return violations
	}

	genFiles, err := filepath.Glob(
		filepath.Join(c.GeneratedPath, "*.go"),
	)
	if err != nil || len(genFiles) == 0 {
		violations = append(violations, core.Violation{
			File:    c.GeneratedPath,
			Line:    0,
			Rule:    "sqlc",
			Message: "no generated files; run sqlc generate",
		})
		return violations
	}

	for _, gf := range genFiles {
		info, err := os.Stat(gf)
		if err != nil {
			continue
		}
		if info.ModTime().Before(schemaInfo.ModTime()) {
			violations = append(violations, core.Violation{
				File:    gf,
				Line:    0,
				Rule:    "sqlc",
				Message: "outdated; run sqlc generate",
			})
		}
	}

	return violations
}
