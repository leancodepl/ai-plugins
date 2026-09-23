package pluginvalidation

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
)

// TODO: Consider refactoring to accept a single path and to move the report
// update elsewhere, deriving the root path instead of taking it separately.
func readFile(path, relativePath string, report *Report) ([]byte, bool) {
	data, err := os.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			report.add(relativePath, "file is missing")
		} else {
			report.add(relativePath, fmt.Sprintf("cannot read file: %v", err))
		}
		return nil, false
	}
	return data, true
}

func isFile(path string) bool {
	info, err := os.Stat(path)
	return err == nil && !info.IsDir()
}

func isDir(path string) bool {
	info, err := os.Stat(path)
	return err == nil && info.IsDir()
}

func relative(root, path string) string {
	relative, err := filepath.Rel(root, path)
	if err != nil {
		return filepath.ToSlash(path)
	}
	return filepath.ToSlash(relative)
}

func resolvePluginRelative(pluginDir, pointer string) (string, bool) {
	pointer = filepath.FromSlash(pointer)

	if !filepath.IsLocal(pointer) || filepath.Clean(pointer) == "." {
		return "", false
	}
	return filepath.Join(pluginDir, pointer), true
}

func sortedKeys[V any](m map[string]V) []string {
	keys := make([]string, 0, len(m))
	for key := range m {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	return keys
}
