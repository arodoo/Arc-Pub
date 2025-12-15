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
