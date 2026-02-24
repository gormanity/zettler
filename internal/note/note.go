package note

import "time"

// Slugify converts a title into a kebab-case slug.
func Slugify(title string) string {
	return ""
}

// TimestampSlug returns a slug based on the given time.
func TimestampSlug(t time.Time) string {
	return ""
}

// Create creates a new note in dir with the given title and returns its path.
// If title is empty, a timestamp slug is used as the filename.
// Returns an error if the file already exists.
func Create(dir, title string, now time.Time) (string, error) {
	return "", nil
}
