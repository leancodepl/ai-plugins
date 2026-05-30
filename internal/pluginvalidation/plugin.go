package pluginvalidation

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func (v *validator) validatePlugin(name string) {
	pluginDir := filepath.Join(v.root, filepath.FromSlash(v.config.PluginRoot), name)
	where := filepath.ToSlash(filepath.Join(v.config.PluginRoot, name))

	if !isDir(pluginDir) {
		v.report.add(where, "directory is missing but the plugin is registered in the marketplace")
		return
	}

	manifestPath := filepath.Join(pluginDir, filepath.FromSlash(pluginManifestPath))
	manifestRelPath := filepath.ToSlash(filepath.Join(v.config.PluginRoot, name, pluginManifestPath))
	if manifest, ok := v.loadManifest(manifestPath, manifestRelPath); ok {
		v.validateManifest(pluginDir, manifestPath, name, manifest)
	}

	if !isFile(filepath.Join(pluginDir, readmeFile)) {
		v.report.add(where, fmt.Sprintf("missing %s", readmeFile))
	}
	v.validateSkills(pluginDir, where)
}

func (v *validator) loadManifest(path, relPath string) (pluginManifest, bool) {
	data, ok := readFile(path, relPath, v.report)
	if !ok {
		return pluginManifest{}, false
	}

	var manifest pluginManifest
	if err := json.Unmarshal(data, &manifest); err != nil {
		v.report.add(relPath, fmt.Sprintf("invalid JSON: %v", err))
		return pluginManifest{}, false
	}
	return manifest, true
}

func (v *validator) validateManifest(pluginDir, manifestPath, pluginName string, manifest pluginManifest) {
	where := relative(v.root, manifestPath)
	expectedSkillsPtr := manifestDirectoryPointer(skillsDir)

	if manifest.Name != pluginName {
		v.report.add(where, fmt.Sprintf("`name` must match directory (%q), got %q", pluginName, manifest.Name))
	}
	v.requireNonEmpty(where, "version", manifest.Version)
	v.validateManifestDirPointer(pluginDir, where, "skills", manifest.Skills, expectedSkillsPtr)
	v.validateManifestFilePointer(pluginDir, where, "logo", manifest.Logo)
	v.validateManifestFilePointer(pluginDir, where, "mcpServers", manifest.MCPServers)
}

func manifestDirectoryPointer(dir string) string {
	return "./" + filepath.ToSlash(dir) + "/"
}

func (v *validator) validateManifestDirPointer(pluginDir, where, key, pointer, expected string) {
	if pointer == "" {
		v.report.add(where, fmt.Sprintf("`%s` must be %q", key, expected))
		return
	}
	if pointer != expected {
		v.report.add(where, fmt.Sprintf("`%s` must be %q (got %q)", key, expected, pointer))
	}

	target, ok := resolvePluginRelative(pluginDir, pointer)
	if !ok {
		v.report.add(where, fmt.Sprintf("`%s` must be a relative path inside the plugin directory (got %q)", key, pointer))
		return
	}
	if !isDir(target) {
		v.report.add(where, fmt.Sprintf("`%s` points to missing directory: %s", key, pointer))
	}
}

func (v *validator) validateManifestFilePointer(pluginDir, where, key, pointer string) {
	if pointer == "" {
		return
	}

	target, ok := resolvePluginRelative(pluginDir, pointer)
	if !ok {
		v.report.add(where, fmt.Sprintf("`%s` must be a relative path inside the plugin directory (got %q)", key, pointer))
		return
	}
	if !isFile(target) {
		v.report.add(where, fmt.Sprintf("`%s` points to missing file: %s", key, pointer))
	}
}

func (v *validator) validateSkills(pluginDir string, where string) {
	skillsPath := filepath.Join(pluginDir, skillsDir)
	if !isDir(skillsPath) {
		v.report.add(where, fmt.Sprintf("missing %s/ directory", skillsDir))
		return
	}

	matches, err := filepath.Glob(filepath.Join(skillsPath, "*", skillFile))
	if err != nil {
		v.report.add(where, fmt.Sprintf("cannot scan %s/: %v", skillsDir, err))
		return
	}
	if len(matches) == 0 {
		v.report.add(where, fmt.Sprintf("%s/ must contain at least one `<skill-name>/%s`", skillsDir, skillFile))
		return
	}

	hasUsage := false
	for _, match := range matches {
		name := strings.ToLower(filepath.Base(filepath.Dir(match)))
		if isUsageSkillName(name) {
			hasUsage = true
			break
		}
	}
	if !hasUsage {
		v.report.add(where, "a usage skill is required")
	}
}

func isUsageSkillName(name string) bool {
	return name == "usage" ||
		strings.HasSuffix(name, "-usage") ||
		strings.HasSuffix(name, "_usage")
}

func (v *validator) validateStrayPlugins(registered map[string]struct{}) {
	pluginsDir := filepath.Join(v.root, filepath.FromSlash(v.config.PluginRoot))
	entries, err := os.ReadDir(pluginsDir)
	if err != nil {
		v.report.add(v.config.PluginRoot+"/", fmt.Sprintf("directory is missing or unreadable: %v", err))
		return
	}

	for _, entry := range entries {
		if !entry.IsDir() || strings.HasPrefix(entry.Name(), ".") {
			continue
		}
		if _, ok := registered[entry.Name()]; !ok {
			v.report.add(
				filepath.ToSlash(filepath.Join(v.config.PluginRoot, entry.Name())),
				"directory is not registered in the marketplace (stray plugin, likely rename leftover)",
			)
		}
	}
}
