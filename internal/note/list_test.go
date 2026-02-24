package note_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/gormanity/zettler/internal/note"
)

func TestList(t *testing.T) {
	vault := t.TempDir()

	// Create notes in various subdirectories.
	files := []string{
		"inbox/capture.md",
		"projects/work/meeting-notes.md",
		"journal/2026-02-24.md",
	}
	for _, f := range files {
		path := filepath.Join(vault, filepath.FromSlash(f))
		if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
			t.Fatal(err)
		}
		if err := os.WriteFile(path, []byte("# Note"), 0644); err != nil {
			t.Fatal(err)
		}
	}

	t.Run("returns all markdown files as relative paths", func(t *testing.T) {
		got, err := note.List(vault)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if len(got) != len(files) {
			t.Fatalf("got %d notes, want %d: %v", len(got), len(files), got)
		}

		// Results should be sorted.
		want := []string{
			"inbox/capture.md",
			"journal/2026-02-24.md",
			"projects/work/meeting-notes.md",
		}
		for i, w := range want {
			if got[i] != w {
				t.Errorf("got[%d] = %q, want %q", i, got[i], w)
			}
		}
	})

	t.Run("returns empty slice for empty vault", func(t *testing.T) {
		empty := t.TempDir()
		got, err := note.List(empty)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if len(got) != 0 {
			t.Errorf("expected empty slice, got %v", got)
		}
	})
}
