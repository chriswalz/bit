package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

// releaseCmd represents the release command
var releaseCmd = &cobra.Command{
	Use:   "release",
	Short: "Commit unstaged changes, bump minor tag, push",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		version := args[0]
		if version == "bump" {
			version = GenBumpedSemVersion()
		}
		save("")
		err := tagCurrentBranch(version)
		if err != nil {
			log.Println(err)
			return
		}
		RunInTerminalWithColor("git", []string{"push", "--force-with-lease"})
		RunInTerminalWithColor("git", []string{"push", "--tags"})
	},
	Args: cobra.ExactArgs(1),
}

func init() {
	BitCmd.AddCommand(releaseCmd)
}
