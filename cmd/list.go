package cmd

import "github.com/spf13/cobra"

// NewListCmd returns the list subcommand.
func NewListCmd(cfgPath string) *cobra.Command {
	return &cobra.Command{
		Use:   "list",
		Short: "List all notes in the vault",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
	}
}
