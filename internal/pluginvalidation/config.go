package pluginvalidation

import (
	"encoding/json"
	"fmt"
	"path/filepath"
)

func (v *validator) loadToolingConfig() bool {
	data, ok := readFile(filepath.Join(v.root, toolingConfigPath), toolingConfigPath, v.report)
	if !ok {
		return false
	}

	if err := json.Unmarshal(data, &v.config); err != nil {
		v.report.add(toolingConfigPath, fmt.Sprintf("invalid JSON: %v", err))
		return false
	}
	v.requireNonEmpty(toolingConfigPath, "pluginRoot", v.config.PluginRoot)
	return v.report.OK()
}
