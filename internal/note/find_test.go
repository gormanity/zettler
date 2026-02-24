package note_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/gormanity/zettler/internal/note"
)

func TestFind(t *testing.T) {
	vault := t.TempDir()

	// Create a note nested in a subdirectory.
	subdir := filepath.Join(vault, "projects", "my-project")
	if err := os.MkdirAll(subdir, 0755); err != nil {
		t.Fatal(err)
	}
	notePath := filepath.Join(subdir, "my-note.md")
	if err := os.WriteFile(notePath, []byte("# My Note"), 0644); err != nil {
		t.Fatal(err)
	}

	t.Run("finds note by exact slug", func(t *testing.T) {
		got, err := note.Find(vault, "my-note")
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if got != notePath {
			t.Errorf("got %q, want %q", got, notePath)
		}
	})

	t.Run("returns error when slug not found", func(t *testing.T) {
		_, err := note.Find(vault, "nonexistent")
		if err == nil {
			t.Error("expected error for missing slug, got nil")
		}
	})
}
