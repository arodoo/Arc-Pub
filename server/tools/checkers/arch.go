// File: arch.go
// Purpose: Enforces architectural layer boundaries by detecting forbidden
// imports. Validates that cmd/ layer does not directly import infrastructure
// packages, maintaining proper dependency direction per hexagonal architecture.
// Parses Go source files to extract imports and checks against forbidden list.
// Part of architecture quality gates preventing coupling violations that make
// code harder to test and maintain. Reports each violation with line number.
// Path: server/tools/checkers/arch.go
// All Rights Reserved. Arc-Pub.

package checkers

import (
	"go/parser"
	"go/token"
	"strings"

	"github.com/arc-pub/server/tools/core"
)

var forbiddenImports = []string{
	"internal/infra/postgres",
	"internal/infra/http",
}

// Arch checks for layer import violations.
type Arch struct{}

// NewArch creates an Arch checker.
func NewArch() *Arch {
	return &Arch{}
}

// Name returns the checker identifier.
func (c *Arch) Name() string {
	return "arch"
}

// Check verifies cmd doesn't import forbidden packages.
func (c *Arch) Check(files []core.File) []core.Violation {
	var violations []core.Violation

	for _, f := range files {
		if !strings.Contains(f.Path, "/cmd/") {
			continue
		}
		if !strings.HasSuffix(f.Path, ".go") {
			continue
		}

		fset := token.NewFileSet()
		node, err := parser.ParseFile(fset, f.Path, f.Content, parser.ImportsOnly)
		if err != nil {
			continue
		}

		for _, imp := range node.Imports {
			impPath := strings.Trim(imp.Path.Value, `"`)
			for _, forbidden := range forbiddenImports {
				if strings.Contains(impPath, forbidden) {
					violations = append(violations, core.Violation{
						File:    f.Path,
						Line:    fset.Position(imp.Pos()).Line,
						Rule:    "arch",
						Message: "cmd cannot import " + forbidden,
					})
				}
			}
		}
	}

	return violations
}
