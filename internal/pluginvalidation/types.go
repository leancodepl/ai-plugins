package pluginvalidation

import "encoding/json"

type marketplace struct {
	Metadata map[string]any      `json:"metadata"`
	Plugins  []marketplacePlugin `json:"plugins"`
}

type marketplacePlugin struct {
	Name        string
	Source      string
	External    bool
	Description string
}

// UnmarshalJSON accepts both source forms: a local path string and an object
// pointing at an external repository ({"source": "github", ...}). External
// plugins have no local directory, so the validator skips them.
func (p *marketplacePlugin) UnmarshalJSON(data []byte) error {
	var raw struct {
		Name        string          `json:"name"`
		Source      json.RawMessage `json:"source"`
		Description string          `json:"description"`
	}
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}
	p.Name = raw.Name
	p.Description = raw.Description
	if len(raw.Source) == 0 {
		return nil
	}
	if raw.Source[0] == '{' {
		p.External = true
		return nil
	}
	return json.Unmarshal(raw.Source, &p.Source)
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
