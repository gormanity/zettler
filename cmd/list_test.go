package cmd_test

import (
	"bytes"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/gormanity/zettler/cmd"
)

func TestListCommand(t *testing.T) {
	vault := t.TempDir()

	cfgPath := filepath.Join(t.TempDir(), "config.toml")
	if err := os.WriteFile(cfgPath, []byte(`vault = "`+vault+`"`+"\n"), 0644); err != nil {
		t.Fatal(err)
	}

	// Create a couple of notes.
	inbox := filepath.Join(vault, "inbox")
	if err := os.MkdirAll(inbox, 0755); err != nil {
		t.Fatal(err)
	}
	for _, name := range []string{"alpha.md", "beta.md"} {
		if err := os.WriteFile(filepath.Join(inbox, name), []byte("# Note"), 0644); err != nil {
			t.Fatal(err)
		}
	}

	var out bytes.Buffer
	root := cmd.NewRootCmd()
	listCmd := cmd.NewListCmd(cfgPath)
	root.AddCommand(listCmd)
	root.SetOut(&out)
	root.SetArgs([]string{"list"})

	if err := root.Execute(); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	output := out.String()
	if !strings.Contains(output, "inbox/alpha.md") {
		t.Errorf("expected alpha.md in output:\n%s", output)
	}
	if !strings.Contains(output, "inbox/beta.md") {
		t.Errorf("expected beta.md in output:\n%s", output)
	}
}
