// File: violation.go
// Purpose: Defines the Violation struct and Checker interface for quality
// tools. Violation captures rule violations with file path, line number, rule
// name, and descriptive message for actionable error reporting. Checker is
// the core interface all quality checks implement, enabling polymorphic
// execution of different rules. Name method identifies the checker for
// reporting, Check method performs the actual validation logic.
// Path: server/tools/core/violation.go
// All Rights Reserved. Arc-Pub.

package core

// Violation represents a single rule violation.
type Violation struct {
	File    string
	Line    int
	Rule    string
	Message string
}

// Checker defines a quality check operation.
type Checker interface {
	Name() string
	Check(files []File) []Violation
}
