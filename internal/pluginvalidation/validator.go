package pluginvalidation

import (
	"fmt"
	"path/filepath"
)

func Validate(repoRoot string) Report {
	var report Report

	root, err := filepath.Abs(repoRoot)
	if err != nil {
		report.add(repoRoot, fmt.Sprintf("cannot resolve repository root: %v", err))
		return report
	}

	v := &validator{root: root, report: &report}
	if !v.loadToolingConfig() {
		return report
	}

	entries := v.validateMarketplace()
	registered := v.registeredPlugins(entries)
	report.PluginCount = len(registered)

	for _, name := range sortedKeys(registered) {
		v.validatePlugin(name)
	}
	v.validateStrayPlugins(registered)

	return report
}

type validator struct {
	root   string
	config toolingConfig
	report *Report
}

func (v *validator) requireNonEmpty(where, field, value string) bool {
	if value == "" {
		v.report.add(where, fmt.Sprintf("`%s` must be a non-empty string", field))
		return false
	}
	return true
}
