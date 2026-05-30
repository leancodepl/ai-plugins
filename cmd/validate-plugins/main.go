package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/leancodepl/ai-plugins/internal/pluginvalidation"
)

func main() {
	repoRoot := flag.String("repo-root", ".", "path to the ai-plugins repository root")
	flag.Parse()

	report := pluginvalidation.Validate(*repoRoot)
	if report.OK() {
		fmt.Printf("OK: plugin structure is valid (%d plugin(s) checked).\n", report.PluginCount)
		return
	}

	fmt.Fprintln(os.Stderr, "FAIL: plugin structure has errors:")
	for _, err := range report.Errors {
		fmt.Fprintf(os.Stderr, "  - %s\n", err)
	}
	fmt.Fprintf(os.Stderr, "\n%d issue(s) found.\n", len(report.Errors))
	os.Exit(1)
}
