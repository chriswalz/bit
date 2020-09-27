package cmd

import (
	"fmt"
	"github.com/chriswalz/bit/gitextras"
	"github.com/chriswalz/bit/util"
	"github.com/spf13/cobra"
	"os"
	"path/filepath"
)

// infoCmd represents the info command
var infoCmd = &cobra.Command{
	Use:   "info",
	Short: "Get general information about the status of your repository",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		os.MkdirAll(filepath.Dir("/tmp/bit/git-extras/"), os.ModePerm)
		fmt.Println("--- INFO ---")
		util.RunScriptWithString("/tmp/bit/git-extras/git-info.sh", gitextras.GitInfo)

		fmt.Println("--- SUMMARY ---")
		util.RunScriptWithString("/tmp/bit/git-extras/git-summary.sh", gitextras.GitSummary)

		fmt.Println("--- EFFORT ---")
		util.RunScriptWithString("/tmp/bit/git-extras/git-effort.sh", gitextras.GitEffort)
	},
	//Args: cobra.MaximumNArgs(1),
}

func init() {
	shellCmd.AddCommand(infoCmd)
}
