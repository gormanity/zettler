package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"time"

	"github.com/gormanity/zettler/internal/config"
	"github.com/spf13/cobra"
)

// NewScratchCmd returns the scratch subcommand.
func NewScratchCmd(cfgPath string) *cobra.Command {
	return &cobra.Command{
		Use:   "scratch",
		Short: "Open a throwaway note in $TMPDIR",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			cfg, err := config.Load(cfgPath)
			if err != nil {
				return err
			}

			name := "scratch-" + time.Now().Format("20060102150405") + ".md"
			path := filepath.Join(os.TempDir(), name)

			if err := os.WriteFile(path, []byte{}, 0644); err != nil {
				return fmt.Errorf("creating scratch file: %w", err)
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
