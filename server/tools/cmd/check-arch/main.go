// File: main.go
// Purpose: Entry point for check-arch quality tool. Validates architecture
// layering rules. Exits with code 1 if violations found.
// Path: server/tools/cmd/check-arch/main.go
// All Rights Reserved. Arc-Pub.

package main

import (
	"os"

	"github.com/arc-pub/server/tools/checkers/sync"
	"github.com/arc-pub/server/tools/reporters"
	"github.com/arc-pub/server/tools/scanners"
)

func main() {
	scanner := scanners.NewFileScanner()
	files, err := scanner.Scan(".", []string{".go"})
	if err != nil {
		os.Stderr.WriteString("scan error: " + err.Error() + "\n")
		os.Exit(1)
	}

	checker := sync.NewArch()
	violations := checker.Check(files)

	reporter := reporters.NewConsole()
	reporter.Report(os.Stdout, violations)

	if len(violations) > 0 {
		os.Exit(1)
	}
}
