// File: i18n.go
// Purpose: Validates internationalization compliance in GDScript files by
// detecting hardcoded UI strings that should use tr() translation function.
// Scans for string assignments to .text and .tooltip properties not wrapped
// in tr() calls. This ensures all user-facing text can be localized. Part
// of the automated quality gates defined in project global rules for visual
// consistency and internationalization compliance across the Godot client.
// Path: server/tools/checkers/i18n.go
// All Rights Reserved. Arc-Pub.

package checkers

import (
	"regexp"
	"strings"

	"github.com/arc-pub/server/tools/core"
)

var (
	stringLiteralRe = regexp.MustCompile(`\.(text|tooltip)\s*=\s*"[^"]+`)
	trWrappedRe     = regexp.MustCompile(`tr\s*\(\s*"`)
)

// I18n checks for hardcoded strings in GDScript.
type I18n struct{}

// NewI18n creates an I18n checker.
func NewI18n() *I18n {
	return &I18n{}
}

// Name returns the checker identifier.
func (c *I18n) Name() string {
	return "i18n"
}

// Check scans GDScript files for non-translated strings.
func (c *I18n) Check(files []core.File) []core.Violation {
	var violations []core.Violation

	for _, f := range files {
		if !strings.HasSuffix(f.Path, ".gd") {
			continue
		}

		for i, line := range f.Lines {
			if !stringLiteralRe.MatchString(line) {
				continue
			}
			if trWrappedRe.MatchString(line) {
				continue
			}

			violations = append(violations, core.Violation{
				File:    f.Path,
				Line:    i + 1,
				Rule:    "i18n",
				Message: "hardcoded string; use tr()",
			})
		}
	}

	return violations
}
