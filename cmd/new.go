package cmd

import "github.com/spf13/cobra"

// NewNoteCmd returns the new subcommand.
func NewNoteCmd(cfgPath string) *cobra.Command {
	return &cobra.Command{
		Use:   "new [title]",
		Short: "Create a new note",
		Args:  cobra.MaximumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
	}
}
