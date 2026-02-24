package cmd_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/gormanity/zettler/cmd"
)

func TestNewCommand(t *testing.T) {
	vault := t.TempDir()

	cfgPath := filepath.Join(t.TempDir(), "config.toml")
	if err := os.WriteFile(cfgPath, []byte(`vault = "`+vault+`"`+"\n"), 0644); err != nil {
		t.Fatal(err)
	}

	t.Setenv("EDITOR", "true")

	t.Run("creates note in inbox by default", func(t *testing.T) {
		root := cmd.NewRootCmd()
		root.AddCommand(cmd.NewNoteCmd(cfgPath))
		root.SetArgs([]string{"new", "My New Note"})

		if err := root.Execute(); err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		entries, err := os.ReadDir(filepath.Join(vault, "inbox"))
		if err != nil {
			t.Fatalf("inbox directory not created: %v", err)
		}
		if len(entries) != 1 {
			t.Fatalf("expected 1 note, got %d", len(entries))
		}
		if entries[0].Name() != "my-new-note.md" {
			t.Errorf("unexpected filename: %q", entries[0].Name())
		}
	})

	t.Run("creates note in specified --dir", func(t *testing.T) {
		targetDir := filepath.Join(vault, "projects")
		if err := os.MkdirAll(targetDir, 0755); err != nil {
			t.Fatal(err)
		}

		root := cmd.NewRootCmd()
		root.AddCommand(cmd.NewNoteCmd(cfgPath))
		root.SetArgs([]string{"new", "Project Note", "--dir", "projects"})

		if err := root.Execute(); err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		entries, err := os.ReadDir(targetDir)
		if err != nil {
			t.Fatalf("target directory not readable: %v", err)
		}
		if len(entries) != 1 {
			t.Fatalf("expected 1 note in projects/, got %d", len(entries))
		}
	})
}
