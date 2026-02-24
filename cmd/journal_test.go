package cmd_test

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/gormanity/zettler/cmd"
)

func TestJournalCommand(t *testing.T) {
	vault := t.TempDir()

	cfgDir := t.TempDir()
	cfgPath := filepath.Join(cfgDir, "config.toml")
	if err := os.WriteFile(cfgPath, []byte(`vault = "`+vault+`"`+"\n"), 0644); err != nil {
		t.Fatal(err)
	}

	// Use a no-op editor so the command doesn't block.
	t.Setenv("EDITOR", "true")

	root := cmd.NewRootCmd()
	journalCmd := cmd.NewJournalCmd(cfgPath)
	root.AddCommand(journalCmd)
	root.SetArgs([]string{"journal"})

	if err := root.Execute(); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Today's entry should exist in vault/journal/
	entries, err := os.ReadDir(filepath.Join(vault, "journal"))
	if err != nil {
		t.Fatalf("journal directory not created: %v", err)
	}
	if len(entries) != 1 {
		t.Fatalf("expected 1 journal entry, got %d", len(entries))
	}
	name := entries[0].Name()
	if !strings.HasSuffix(name, ".md") {
		t.Errorf("expected .md file, got %q", name)
	}
}
