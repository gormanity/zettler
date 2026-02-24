package cmd

import "github.com/spf13/cobra"

// NewJournalCmd returns the journal subcommand.
func NewJournalCmd(cfgPath string) *cobra.Command {
	return &cobra.Command{
		Use:   "journal [date]",
		Short: "Open today's journal entry, or a past entry by date",
		Args:  cobra.MaximumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
	}
}
