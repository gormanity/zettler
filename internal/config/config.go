package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/BurntSushi/toml"
)

// DefaultPath returns the default config file path.
func DefaultPath() (string, error) {
	dir, err := os.UserConfigDir()
	if err != nil {
		return "", fmt.Errorf("finding config directory: %w", err)
	}
	return filepath.Join(dir, "zettler", "config.toml"), nil
}

// Config holds the zettler configuration.
type Config struct {
	Vault  string `toml:"vault"`
	Editor string `toml:"editor"`
}

// Load reads a Config from the TOML file at path.
func Load(path string) (*Config, error) {
	var cfg Config
	if _, err := toml.DecodeFile(path, &cfg); err != nil {
		return nil, fmt.Errorf("loading config from %s: %w", path, err)
	}
	return &cfg, nil
}

// ResolveEditor returns the configured editor, falling back to $EDITOR.
func (c *Config) ResolveEditor() string {
	if c.Editor != "" {
		return c.Editor
	}
	return os.Getenv("EDITOR")
}
