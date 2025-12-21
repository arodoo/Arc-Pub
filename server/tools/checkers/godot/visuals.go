// File: visuals.go
// Purpose: Enforces visual consistency rules by detecting hardcoded colors
// in GDScript and scene files. Colors should reference the Palette singleton
// for centralized theming control. Detects Color() constructor usage and
// hexadecimal color codes (#RRGGBB). Part of automated quality gates for
// visual consistency as defined in project global rules. Enables easy theme
// changes by ensuring all colors flow through a single source of truth.
// Path: server/tools/checkers/visuals.go
// All Rights Reserved. Arc-Pub.

package godot

import (
	"regexp"

	"github.com/arc-pub/server/tools/core"
)

var (
	colorConstructorRe = regexp.MustCompile(`Color\s*\(`)
	hexColorRe         = regexp.MustCompile(`#[0-9A-Fa-f]{6,8}`)
)

// Visuals checks for hardcoded colors.
type Visuals struct{}

// NewVisuals creates a Visuals checker.
func NewVisuals() *Visuals {
	return &Visuals{}
}

// Name returns the checker identifier.
func (c *Visuals) Name() string {
	return "visuals"
}

// Check scans for Color() or HEX codes.
func (c *Visuals) Check(files []core.File) []core.Violation {
	var violations []core.Violation

	for _, f := range files {
		for i, line := range f.Lines {
			if colorConstructorRe.MatchString(line) {
				violations = append(violations, core.Violation{
					File:    f.Path,
					Line:    i + 1,
					Rule:    "visuals",
					Message: "hardcoded Color(); use Palette",
				})
			}
			if hexColorRe.MatchString(line) {
				violations = append(violations, core.Violation{
					File:    f.Path,
					Line:    i + 1,
					Rule:    "visuals",
					Message: "hardcoded HEX color; use Palette",
				})
			}
		}
	}

	return violations
}
