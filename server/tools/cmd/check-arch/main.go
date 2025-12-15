// check-arch: Verifies layer import restrictions.
package main

import (
	"os"

	"github.com/arc-pub/server/tools/checkers"
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

	checker := checkers.NewArch()
	violations := checker.Check(files)

	reporter := reporters.NewConsole()
	reporter.Report(os.Stdout, violations)

	if len(violations) > 0 {
		os.Exit(1)
	}
}
