package cmd

import "github.com/spf13/cobra"

// NewScratchCmd returns the scratch subcommand.
func NewScratchCmd(cfgPath string) *cobra.Command {
	return &cobra.Command{
		Use:   "scratch",
		Short: "Open a throwaway note in $TMPDIR",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
	}
}
