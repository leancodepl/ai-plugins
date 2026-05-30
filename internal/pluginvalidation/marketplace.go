package pluginvalidation

import (
	"encoding/json"
	"fmt"
	"path/filepath"
)

func (v *validator) loadMarketplace() *marketplace {
	data, ok := readFile(filepath.Join(v.root, filepath.FromSlash(marketplacePath)), marketplacePath, v.report)
	if !ok {
		return nil
	}

	var mp marketplace
	if err := json.Unmarshal(data, &mp); err != nil {
		v.report.add(marketplacePath, fmt.Sprintf("invalid JSON: %v", err))
		return nil
	}
	return &mp
}

func (v *validator) validateMarketplace() []marketplacePlugin {
	mp := v.loadMarketplace()
	if mp == nil {
		return nil
	}

	v.validateMarketplaceEntries(marketplacePath, mp.Plugins)
	for i, entry := range mp.Plugins {
		if entry.Source == "" {
			continue
		}
		v.validateSource(marketplacePath, i, entry)
	}

	return mp.Plugins
}

func (v *validator) validateSource(where string, i int, entry marketplacePlugin) {
	if entry.Name == "" {
		return
	}
	expected := "./" + filepath.ToSlash(filepath.Join(v.config.PluginRoot, entry.Name))
	if entry.Source != expected {
		v.report.add(
			where,
			fmt.Sprintf("plugins[%d].source must be %q (got %q)", i, expected, entry.Source),
		)
	}
}

func (v *validator) validateMarketplaceEntries(where string, entries []marketplacePlugin) {
	if entries == nil {
		v.report.add(where, "`plugins` must be a list")
		return
	}

	seen := map[string]int{}
	for i, entry := range entries {
		if v.requireNonEmpty(where, fmt.Sprintf("plugins[%d].name", i), entry.Name) {
			if previous, ok := seen[entry.Name]; ok {
				v.report.add(where, fmt.Sprintf("plugins[%d].name duplicates plugins[%d].name (%q)", i, previous, entry.Name))
			} else {
				seen[entry.Name] = i
			}
		}
		v.requireNonEmpty(where, fmt.Sprintf("plugins[%d].source", i), entry.Source)
	}
}

func (v *validator) registeredPlugins(entries []marketplacePlugin) map[string]struct{} {
	registered := map[string]struct{}{}
	for _, entry := range entries {
		if entry.Name != "" {
			registered[entry.Name] = struct{}{}
		}
	}
	return registered
}
