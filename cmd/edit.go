package cmd

import (
	"fmt"
	"os/exec"

	"github.com/gormanity/zettler/internal/config"
	"github.com/gormanity/zettler/internal/note"
	"github.com/spf13/cobra"
)

// NewEditCmd returns the edit subcommand.
func NewEditCmd(cfgPath string) *cobra.Command {
	return &cobra.Command{
		Use:   "edit <slug>",
		Short: "Open a note by slug",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cfg, err := config.Load(cfgPath)
			if err != nil {
				return err
			}

			path, err := note.Find(cfg.Vault, args[0])
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
}
