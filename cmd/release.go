package cmd

import (
	git_extras "github.com/chriswalz/bit/git-extras"
	"github.com/chriswalz/bit/util"
	"github.com/spf13/cobra"
	"os"
	"path/filepath"
)

// releaseCmd represents the release command
var releaseCmd = &cobra.Command{
	Use:   "release",
	Short: "Generate a production release",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		arg := args[0]
		if arg == "bump" {
			arg = util.GenBumpedSemVersion()
		}
		os.MkdirAll(filepath.Dir("/tmp/bit/git-extras/"), os.ModePerm)
		util.RunScriptWithString("/tmp/bit/git-extras/git-release.sh", git_extras.GitRelease, arg)
	},
	Args: cobra.ExactArgs(1),
}

func init() {
	shellCmd.AddCommand(releaseCmd)
}
