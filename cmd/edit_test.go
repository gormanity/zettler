package cmd_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/gormanity/zettler/cmd"
)

func TestEditCommand(t *testing.T) {
	vault := t.TempDir()

	cfgPath := filepath.Join(t.TempDir(), "config.toml")
	if err := os.WriteFile(cfgPath, []byte(`vault = "`+vault+`"`+"\n"), 0644); err != nil {
		t.Fatal(err)
	}

	// Create a note in the vault to find.
	inbox := filepath.Join(vault, "inbox")
	if err := os.MkdirAll(inbox, 0755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(inbox, "my-note.md"), []byte("# My Note"), 0644); err != nil {
		t.Fatal(err)
	}

	t.Setenv("EDITOR", "true")

	t.Run("opens note by exact slug", func(t *testing.T) {
		root := cmd.NewRootCmd()
		root.AddCommand(cmd.NewEditCmd(cfgPath))
		root.SetArgs([]string{"edit", "my-note"})

		if err := root.Execute(); err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
	})

	t.Run("returns error for unknown slug", func(t *testing.T) {
		root := cmd.NewRootCmd()
		root.AddCommand(cmd.NewEditCmd(cfgPath))
		root.SetArgs([]string{"edit", "nonexistent"})
		root.SilenceErrors = true

		if err := root.Execute(); err == nil {
			t.Error("expected error for unknown slug, got nil")
		}
	})
}
