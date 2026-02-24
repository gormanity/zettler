package note

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"
)

var nonAlphanumeric = regexp.MustCompile(`[^a-z0-9]+`)

// Slugify converts a title into a kebab-case slug.
func Slugify(title string) string {
	s := strings.ToLower(title)
	s = strings.ReplaceAll(s, "'", "")
	s = nonAlphanumeric.ReplaceAllString(s, "-")
	s = strings.Trim(s, "-")
	return s
}

// TimestampSlug returns a slug based on the given time.
func TimestampSlug(t time.Time) string {
	return t.Format("20060102150405")
}

// Create creates a new note in dir with the given title and returns its path.
// If title is empty, a timestamp slug is used as the filename.
// Returns an error if the file already exists.
func Create(dir, title string, now time.Time) (string, error) {
	if err := os.MkdirAll(dir, 0755); err != nil {
		return "", fmt.Errorf("creating directory: %w", err)
	}

	var slug string
	if title == "" {
		slug = TimestampSlug(now)
	} else {
		slug = Slugify(title)
	}

	path := filepath.Join(dir, slug+".md")

	if _, err := os.Stat(path); err == nil {
		return "", fmt.Errorf("note already exists: %s", path)
	}

	content := fmt.Sprintf("---\ntitle: %q\ncreated: %s\ntags: []\n---\n",
		title,
		now.Format("2006-01-02"),
	)
	if err := os.WriteFile(path, []byte(content), 0644); err != nil {
		return "", fmt.Errorf("creating note: %w", err)
	}

	return path, nil
}

// Find searches vault recursively for a note with the given slug (filename without .md).
// Returns the full path if found, or an error if not found.
func Find(vault, slug string) (string, error) {
	return "", nil
}
