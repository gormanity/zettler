package cmd

import (
	"fmt"

	"github.com/gormanity/zettler/internal/config"
	"github.com/gormanity/zettler/internal/note"
	"github.com/spf13/cobra"
)

// NewListCmd returns the list subcommand.
func NewListCmd(cfgPath string) *cobra.Command {
	return &cobra.Command{
		Use:   "list",
		Short: "List all notes in the vault",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			cfg, err := config.Load(cfgPath)
			if err != nil {
				return err
			}

			notes, err := note.List(cfg.Vault)
			if err != nil {
				return err
			}

			for _, n := range notes {
				fmt.Fprintln(cmd.OutOrStdout(), n)
			}
			return nil
		},
	}
}
