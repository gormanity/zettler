package cmd_test

import (
	"bytes"
	"testing"

	"github.com/gormanity/zettler/cmd"
	"github.com/stretchr/testify/assert"
)

func TestRootCmd_NoOutput(t *testing.T) {
	// Create the root command
	rootCmd := cmd.NewRootCmd()

	// Capture command output
	var out bytes.Buffer
	rootCmd.SetOut(&out)
	rootCmd.SetErr(&out)
	rootCmd.SetArgs([]string{}) // No arguments

	// Execute the command
	err := rootCmd.Execute()

	// Ensure no error occurs
	assert.NoError(t, err)

	// Ensure that no output is returned
	assert.Equal(
		t,
		"",
		out.String(),
		"Expected no output from running `zettler` with no arguments",
	)
}

func TestRootCmd_HelpFlag(t *testing.T) {
	rootCmd := cmd.NewRootCmd()

	expectedHelpText := rootCmd.Long

	testCases := []struct {
		name string
		args []string
	}{
		{"help flag", []string{"--help"}},
		{"short help flag", []string{"-h"}},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Capture command output
			var out bytes.Buffer
			rootCmd.SetOut(&out)
			rootCmd.SetErr(&out)
			rootCmd.SetArgs(tc.args)

			// Execute the command
			err := rootCmd.Execute()

			// Ensure no error occurs
			assert.NoError(t, err)

			// Verify that the output contains the expected help text
			output := out.String()
			assert.Contains(t, output, expectedHelpText)
		})
	}
}
