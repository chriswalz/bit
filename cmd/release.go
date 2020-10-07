package cmd

import (
	"github.com/spf13/cobra"
)

// releaseCmd represents the release command
var releaseCmd = &cobra.Command{
	Use:   "release",
	Short: "Generate a production release",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		version := args[0]
		if version == "bump" {
			version = GenBumpedSemVersion()
		}
		save("")
		tagCurrentBranch(version)
		RunInTerminalWithColor("git", []string{"push", "--force-with-lease"})
	},
	Args: cobra.ExactArgs(1),
}

func init() {
	ShellCmd.AddCommand(releaseCmd)
}
