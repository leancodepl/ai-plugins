package pluginvalidation

type marketplace struct {
	Metadata map[string]any      `json:"metadata"`
	Plugins  []marketplacePlugin `json:"plugins"`
}

type marketplacePlugin struct {
	Name        string `json:"name"`
	Source      string `json:"source"`
	Description string `json:"description"`
}

type pluginManifest struct {
	Name       string `json:"name"`
	Version    string `json:"version"`
	Skills     string `json:"skills"`
	Logo       string `json:"logo"`
	MCPServers string `json:"mcpServers"`
}

type toolingConfig struct {
	PluginRoot string `json:"pluginRoot"`
}
