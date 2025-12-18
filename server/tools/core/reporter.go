// Arc-Pub - Metaverso 2D MMO Social
// Copyright (c) 2024. MIT License.

package core

import "io"

// Reporter outputs violations to a destination.
type Reporter interface {
	Report(w io.Writer, violations []Violation) error
}
