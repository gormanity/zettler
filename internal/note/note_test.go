package note_test

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/gormanity/zettler/internal/note"
)

var referenceDate = time.Date(2026, 2, 24, 15, 4, 5, 0, time.UTC)

func TestSlugify(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{"Hello World", "hello-world"},
		{"  leading and trailing  ", "leading-and-trailing"},
		{"multiple   spaces", "multiple-spaces"},
		{"Special! Characters?", "special-characters"},
		{"already-kebab-case", "already-kebab-case"},
		{"Mixed CAPS and lower", "mixed-caps-and-lower"},
		{"apostrophe's and \"quotes\"", "apostrophes-and-quotes"},
	}

	for _, tt := range tests {
		got := note.Slugify(tt.input)
		if got != tt.want {
			t.Errorf("Slugify(%q): got %q, want %q", tt.input, got, tt.want)
		}
	}
}

func TestTimestampSlug(t *testing.T) {
	got := note.TimestampSlug(referenceDate)
	want := "20260224150405"
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestCreate(t *testing.T) {
	t.Run("creates note with title in specified directory", func(t *testing.T) {
		dir := t.TempDir()

		path, err := note.Create(dir, "My Test Note", referenceDate)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if filepath.Base(path) != "my-test-note.md" {
			t.Errorf("expected filename my-test-note.md, got %q", filepath.Base(path))
		}

		data, err := os.ReadFile(path)
		if err != nil {
			t.Fatalf("file not created: %v", err)
		}

		content := string(data)
		if !strings.Contains(content, `title: "My Test Note"`) {
			t.Errorf("missing title in frontmatter:\n%s", content)
		}
		if !strings.Contains(content, "created: 2026-02-24") {
			t.Errorf("missing created in frontmatter:\n%s", content)
		}
		if !strings.Contains(content, "tags: []") {
			t.Errorf("missing tags in frontmatter:\n%s", content)
		}
	})

	t.Run("uses timestamp slug when title is empty", func(t *testing.T) {
		dir := t.TempDir()

		path, err := note.Create(dir, "", referenceDate)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if filepath.Base(path) != "20260224150405.md" {
			t.Errorf("expected timestamp filename, got %q", filepath.Base(path))
		}
	})

	t.Run("creates parent directory if missing", func(t *testing.T) {
		vault := t.TempDir()
		inbox := filepath.Join(vault, "inbox")

		_, err := note.Create(inbox, "New Note", referenceDate)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if _, err := os.Stat(inbox); os.IsNotExist(err) {
			t.Error("inbox directory was not created")
		}
	})

	t.Run("returns error if file already exists", func(t *testing.T) {
		dir := t.TempDir()

		if _, err := note.Create(dir, "Duplicate", referenceDate); err != nil {
			t.Fatal(err)
		}
		_, err := note.Create(dir, "Duplicate", referenceDate)
		if err == nil {
			t.Error("expected error for duplicate note, got nil")
		}
	})
}
