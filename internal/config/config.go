package config

// Config holds the zettler configuration.
type Config struct {
	Vault  string `toml:"vault"`
	Editor string `toml:"editor"`
}

// Load reads a Config from the TOML file at path.
func Load(path string) (*Config, error) {
	return nil, nil
}

// ResolveEditor returns the configured editor, falling back to $EDITOR.
func (c *Config) ResolveEditor() string {
	return ""
}
