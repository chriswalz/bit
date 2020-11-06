package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// completeCmd represents the complete command
var completeCmd = &cobra.Command{
	Use:   "complete",
	Short: "Add classical tab completion to bit",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		os.Setenv("COMP_INSTALL", "1")
		Bitcomplete()
	},
	Args: cobra.NoArgs,
}

func init() {
	ShellCmd.AddCommand(completeCmd)
}
