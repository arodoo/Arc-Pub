// check-assets: Verifies asset file sizes.
package main

import (
	"flag"
	"os"

	"github.com/arc-pub/server/tools/checkers"
	"github.com/arc-pub/server/tools/reporters"
)

func main() {
	root := flag.String("root", "..", "Godot project root")
	flag.Parse()

	checker := checkers.NewAssets(*root)
	violations := checker.Check(nil)

	reporter := reporters.NewConsole()
	reporter.Report(os.Stdout, violations)

	if len(violations) > 0 {
		os.Exit(1)
	}
}
