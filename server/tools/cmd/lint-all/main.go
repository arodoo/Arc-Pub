// File: main.go
// Purpose: Entry point for lint-all orchestration tool. Runs all core quality
// checks sequentially: limits, density, arch, and sqlc. Aggregates results
// and exits with code 1 if any check fails. Provides single command for CI
// pipelines. Part of automated quality gate infrastructure.
// Path: server/tools/cmd/lint-all/main.go
// All Rights Reserved. Arc-Pub.

package main

import (
	"fmt"
	"os"
	"os/exec"
)

var checks = []string{
	"./tools/cmd/check-limits",
	"./tools/cmd/check-density",
	"./tools/cmd/check-arch",
	"./tools/cmd/check-sqlc",
}

func main() {
	failed := false

	for _, check := range checks {
		fmt.Printf("Running %s...\n", check)
		cmd := exec.Command("go", "run", check)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		if err := cmd.Run(); err != nil {
			failed = true
		}
	}

	if failed {
		os.Exit(1)
	}
	fmt.Println("✓ All checks passed")
}
