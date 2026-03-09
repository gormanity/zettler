package journal

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

// EntryPath returns the file path for a journal entry.
func EntryPath(vault string, date time.Time) string {
	return filepath.Join(vault, "journal", date.Format("2006-01-02")+".md")
}

// EnsureEntry creates the journal entry file if it does not exist.
// Returns the path to the entry.
func EnsureEntry(vault string, date time.Time) (string, error) {
	path := EntryPath(vault, date)

	if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
		return "", fmt.Errorf("creating journal directory: %w", err)
	}

	if _, err := os.Stat(path); err == nil {
		return path, nil
	}

	content := fmt.Sprintf("---\ntitle: %q\ncreated: %s\nlocation: \ntags: []\n---\n",
		"Journal - "+date.Format("January 2, 2006"),
		date.Format("2006-01-02"),
	)
	if err := os.WriteFile(path, []byte(content), 0644); err != nil {
		return "", fmt.Errorf("creating journal entry: %w", err)
	}

	return path, nil
}

// ParseDate parses a date argument. Supports "yesterday" and "YYYY-MM-DD".
// now is used as the reference time.
func ParseDate(s string, now time.Time) (time.Time, error) {
	if s == "yesterday" {
		return now.AddDate(0, 0, -1), nil
	}
	t, err := time.Parse("2006-01-02", s)
	if err != nil {
		return time.Time{}, fmt.Errorf("invalid date %q: expected YYYY-MM-DD or \"yesterday\"", s)
	}
	return t, nil
}
