package journal

import (
	"path/filepath"
	"time"
)

// EntryPath returns the file path for a journal entry.
func EntryPath(vault string, date time.Time) string {
	return ""
}

// EnsureEntry creates the journal entry file if it does not exist.
// Returns the path to the entry.
func EnsureEntry(vault string, date time.Time) (string, error) {
	return "", nil
}

// ParseDate parses a date argument. Supports "yesterday" and "YYYY-MM-DD".
// now is used as the reference time.
func ParseDate(s string, now time.Time) (time.Time, error) {
	return time.Time{}, nil
}

// suppress unused import error until implemented
var _ = filepath.Join
