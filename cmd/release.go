package cmd

import (
	"fmt"
	"os/exec"

	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

// releaseCmd represents the release command
var releaseCmd = &cobra.Command{
	Use:   "release",
	Short: "Commit unstaged changes, bump minor tag, push",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		var err error
		version := args[0]
		if version == "bump" {
			rawGitTagOutput, err := exec.Command("/bin/sh", "-c", `git for-each-ref --format="%(refname:short)" --sort=-authordate --count=1 refs/tags`).CombinedOutput()
			if err != nil {
				log.Debug().Err(err).Send()
			}
			version, err = GenBumpedSemVersion(string(rawGitTagOutput))
			if err != nil {
				fmt.Println(err)
				return
			}
		}
		save([]string{})
		err = tagCurrentBranch(version)
		if err != nil {
			fmt.Println(err)
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
