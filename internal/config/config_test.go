package config_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/gormanity/zettler/internal/config"
)

func TestLoad(t *testing.T) {
	dir := t.TempDir()
	cfgPath := filepath.Join(dir, "config.toml")
	content := `vault = "/my/notes"` + "\n" + `editor = "vim"` + "\n"
	if err := os.WriteFile(cfgPath, []byte(content), 0644); err != nil {
		t.Fatal(err)
	}

	cfg, err := config.Load(cfgPath)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if cfg.Vault != "/my/notes" {
		t.Errorf("vault: got %q, want %q", cfg.Vault, "/my/notes")
	}
	if cfg.Editor != "vim" {
		t.Errorf("editor: got %q, want %q", cfg.Editor, "vim")
	}
}

func TestLoadMissingFile(t *testing.T) {
	_, err := config.Load("/nonexistent/path/config.toml")
	if err == nil {
		t.Fatal("expected error for missing config file, got nil")
	}
}

func TestResolveEditor(t *testing.T) {
	t.Run("uses config editor when set", func(t *testing.T) {
		cfg := &config.Config{Editor: "nano"}
		if got := cfg.ResolveEditor(); got != "nano" {
			t.Errorf("got %q, want %q", got, "nano")
		}
	})

	t.Run("falls back to $EDITOR", func(t *testing.T) {
		t.Setenv("EDITOR", "emacs")
		cfg := &config.Config{Editor: ""}
		if got := cfg.ResolveEditor(); got != "emacs" {
			t.Errorf("got %q, want %q", got, "emacs")
		}
	})

	t.Run("returns empty string when neither is set", func(t *testing.T) {
		t.Setenv("EDITOR", "")
		cfg := &config.Config{Editor: ""}
		if got := cfg.ResolveEditor(); got != "" {
			t.Errorf("got %q, want empty string", got)
		}
	})
}
