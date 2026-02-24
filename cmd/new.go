package cmd

import (
	"fmt"
	"os/exec"
	"path/filepath"
	"time"

	"github.com/gormanity/zettler/internal/config"
	"github.com/gormanity/zettler/internal/note"
	"github.com/spf13/cobra"
)

// NewNoteCmd returns the new subcommand.
func NewNoteCmd(cfgPath string) *cobra.Command {
	var dir string

	c := &cobra.Command{
		Use:   "new [title]",
		Short: "Create a new note",
		Args:  cobra.MaximumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cfg, err := config.Load(cfgPath)
			if err != nil {
				return err
			}

			var title string
			if len(args) == 1 {
				title = args[0]
			}

			dest := filepath.Join(cfg.Vault, "inbox")
			if dir == "." {
				dest = "."
			} else if dir != "" {
				dest = filepath.Join(cfg.Vault, dir)
			}

			path, err := note.Create(dest, title, time.Now())
			if err != nil {
				return err
			}

			editor := cfg.ResolveEditor()
			if editor == "" {
				return fmt.Errorf("no editor configured: set $EDITOR or editor in config")
			}

			return exec.Command(editor, path).Run()
		},
	}

	c.Flags().StringVar(&dir, "dir", "", "destination directory relative to vault root (or '.' for current directory)")
	return c
}
