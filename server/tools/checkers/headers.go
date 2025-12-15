package checkers

import (
	"strings"

	"github.com/arc-pub/server/tools/core"
)

const licenseHeader = "Arc-Pub"

// Headers checks for license headers in files.
type Headers struct{}

// NewHeaders creates a Headers checker.
func NewHeaders() *Headers {
	return &Headers{}
}

// Name returns the checker identifier.
func (c *Headers) Name() string {
	return "headers"
}

// Check verifies files have license headers.
func (c *Headers) Check(files []core.File) []core.Violation {
	var violations []core.Violation

	for _, f := range files {
		if len(f.Lines) < 2 {
			continue
		}

		headerFound := false
		for i := 0; i < min(5, len(f.Lines)); i++ {
			if strings.Contains(f.Lines[i], licenseHeader) {
				headerFound = true
				break
			}
		}

		if !headerFound {
			violations = append(violations, core.Violation{
				File:    f.Path,
				Line:    1,
				Rule:    "headers",
				Message: "missing license header",
			})
		}
	}

	return violations
}
