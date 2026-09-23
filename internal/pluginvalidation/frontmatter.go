package pluginvalidation

import (
	"regexp"
	"strings"
)

var (
	yamlKeyRE              = regexp.MustCompile(`^([A-Za-z_][A-Za-z0-9_-]*)\s*:\s*(.*)$`)
	frontmatterCloseLineRE = regexp.MustCompile(`(?m)^---[ \t]*$`)
)

func parseFrontmatter(text string) (map[string]any, bool) {
	text = strings.ReplaceAll(text, "\r\n", "\n")
	if !strings.HasPrefix(text, frontmatterOpen) {
		return nil, false
	}

	rest := strings.TrimPrefix(text, frontmatterOpen)
	loc := frontmatterCloseLineRE.FindStringIndex(rest)
	if loc == nil {
		return nil, false
	}

	result := map[string]any{}
	var currentList []string
	var currentListKey string

	for _, rawLine := range strings.Split(rest[:loc[0]], "\n") {
		if strings.TrimSpace(rawLine) == "" {
			continue
		}

		stripped := strings.TrimLeft(rawLine, " \t")
		if strings.HasPrefix(stripped, "- ") {
			if currentListKey == "" {
				continue
			}
			item := stripYAMLQuotes(strings.TrimSpace(strings.TrimPrefix(stripped, "- ")))
			currentList = append(currentList, item)
			result[currentListKey] = currentList
			continue
		}

		matches := yamlKeyRE.FindStringSubmatch(rawLine)
		if matches == nil {
			continue
		}

		key := matches[1]
		value := strings.TrimSpace(matches[2])
		if value == "" {
			currentListKey = key
			currentList = []string{}
			result[key] = currentList
			continue
		}

		currentListKey = ""
		currentList = nil
		value = stripYAMLQuotes(value)
		switch strings.ToLower(value) {
		case "true":
			result[key] = true
		case "false":
			result[key] = false
		default:
			result[key] = value
		}
	}

	return result, true
}

func stripYAMLQuotes(value string) string {
	if len(value) < 2 {
		return value
	}
	first := value[0]
	last := value[len(value)-1]
	if (first == '"' && last == '"') || (first == '\'' && last == '\'') {
		return value[1 : len(value)-1]
	}
	return value
}
