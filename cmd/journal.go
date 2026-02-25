package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"time"

	"github.com/gormanity/zettler/internal/config"
	"github.com/gormanity/zettler/internal/journal"
	"github.com/spf13/cobra"
)

// NewJournalCmd returns the journal subcommand.
func NewJournalCmd(cfgPath string) *cobra.Command {
	return &cobra.Command{
		Use:   "journal [date]",
		Short: "Open today's journal entry, or a past entry by date",
		Args:  cobra.MaximumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cfg, err := config.Load(cfgPath)
			if err != nil {
				return err
			}

			date := time.Now()
			if len(args) == 1 {
				date, err = journal.ParseDate(args[0], time.Now())
				if err != nil {
					return err
				}
			}

			path, err := journal.EnsureEntry(cfg.Vault, date)
			if err != nil {
				return err
			}

			editor := cfg.ResolveEditor()
			if editor == "" {
				return fmt.Errorf("no editor configured: set $EDITOR or editor in config")
			}

			c := exec.Command(editor, path)
			c.Stdin = os.Stdin
			c.Stdout = os.Stdout
			c.Stderr = os.Stderr
			return c.Run()
		},
	}
}
