package cmd_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/gormanity/zettler/cmd"
)

func TestScratchCommand(t *testing.T) {
	cfgPath := filepath.Join(t.TempDir(), "config.toml")
	if err := os.WriteFile(cfgPath, []byte(`vault = "/unused"`+"\n"), 0644); err != nil {
		t.Fatal(err)
	}

	tmpDir := t.TempDir()
	t.Setenv("TMPDIR", tmpDir)
	t.Setenv("EDITOR", "true")

	root := cmd.NewRootCmd()
	root.AddCommand(cmd.NewScratchCmd(cfgPath))
	root.SetArgs([]string{"scratch"})

	if err := root.Execute(); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	entries, err := os.ReadDir(tmpDir)
	if err != nil {
		t.Fatal(err)
	}
	if len(entries) != 1 {
		t.Fatalf("expected 1 scratch file in TMPDIR, got %d", len(entries))
	}
	if filepath.Ext(entries[0].Name()) != ".md" {
		t.Errorf("expected .md file, got %q", entries[0].Name())
	}
}
