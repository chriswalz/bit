package cmd

import (
	"github.com/chriswalz/bit/gitextras"
	"github.com/spf13/cobra"
	"os"
	"path/filepath"
)

// releaseCmd represents the release command
var releaseCmd = &cobra.Command{
	Use:   "release",
	Short: "Generate a production release",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		arg := args[0]
		if arg == "bump" {
			arg = GenBumpedSemVersion()
		}
		os.MkdirAll(filepath.Dir("/tmp/bit/git-extras/"), os.ModePerm)
		RunScriptWithString("/tmp/bit/git-extras/git-release.sh", gitextras.GitRelease, arg)
	},
	Args: cobra.ExactArgs(1),
}

func init() {
	ShellCmd.AddCommand(releaseCmd)
}
