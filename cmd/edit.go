package cmd

import "github.com/spf13/cobra"

// NewEditCmd returns the edit subcommand.
func NewEditCmd(cfgPath string) *cobra.Command {
	return &cobra.Command{
		Use:   "edit <slug>",
		Short: "Open a note by slug",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
	}
}
