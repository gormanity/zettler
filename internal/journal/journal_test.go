package journal_test

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/gormanity/zettler/internal/journal"
)

var referenceDate = time.Date(2026, 2, 24, 12, 0, 0, 0, time.UTC)

func TestEntryPath(t *testing.T) {
	got := journal.EntryPath("/my/vault", referenceDate)
	want := "/my/vault/journal/2026-02-24.md"
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestEnsureEntry(t *testing.T) {
	t.Run("creates file with frontmatter when missing", func(t *testing.T) {
		vault := t.TempDir()

		path, err := journal.EnsureEntry(vault, referenceDate)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		data, err := os.ReadFile(path)
		if err != nil {
			t.Fatalf("file not created: %v", err)
		}

		content := string(data)
		if !strings.Contains(content, `title: "Journal - February 24, 2026"`) {
			t.Errorf("missing title in frontmatter:\n%s", content)
		}
		if !strings.Contains(content, "created: 2026-02-24") {
			t.Errorf("missing created in frontmatter:\n%s", content)
		}
		if !strings.Contains(content, "tags: []") {
			t.Errorf("missing tags in frontmatter:\n%s", content)
		}
	})

	t.Run("creates parent journal directory if missing", func(t *testing.T) {
		vault := t.TempDir()

		_, err := journal.EnsureEntry(vault, referenceDate)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if _, err := os.Stat(filepath.Join(vault, "journal")); os.IsNotExist(err) {
			t.Error("journal directory was not created")
		}
	})

	t.Run("does not overwrite existing file", func(t *testing.T) {
		vault := t.TempDir()
		journalDir := filepath.Join(vault, "journal")
		if err := os.MkdirAll(journalDir, 0755); err != nil {
			t.Fatal(err)
		}
		existing := filepath.Join(journalDir, "2026-02-24.md")
		original := "my existing content"
		if err := os.WriteFile(existing, []byte(original), 0644); err != nil {
			t.Fatal(err)
		}

		_, err := journal.EnsureEntry(vault, referenceDate)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		data, _ := os.ReadFile(existing)
		if string(data) != original {
			t.Errorf("existing file was overwritten")
		}
	})
}

func TestParseDate(t *testing.T) {
	tests := []struct {
		input string
		want  time.Time
	}{
		{"yesterday", referenceDate.AddDate(0, 0, -1)},
		{"2026-02-20", time.Date(2026, 2, 20, 0, 0, 0, 0, time.UTC)},
	}

	for _, tt := range tests {
		got, err := journal.ParseDate(tt.input, referenceDate)
		if err != nil {
			t.Fatalf("ParseDate(%q): unexpected error: %v", tt.input, err)
		}
		if !got.Equal(tt.want) {
			t.Errorf("ParseDate(%q): got %v, want %v", tt.input, got, tt.want)
		}
	}
}

func TestParseDateInvalid(t *testing.T) {
	_, err := journal.ParseDate("not-a-date", referenceDate)
	if err == nil {
		t.Error("expected error for invalid date, got nil")
	}
}
