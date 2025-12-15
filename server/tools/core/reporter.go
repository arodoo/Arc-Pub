package core

import "io"

// Reporter outputs violations to a destination.
type Reporter interface {
	Report(w io.Writer, violations []Violation) error
}
