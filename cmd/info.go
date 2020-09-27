package cmd

import (
	"fmt"
	"github.com/chriswalz/bit/gitextras"
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
		RunScriptWithString("/tmp/bit/git-extras/git-info.sh", gitextras.GitInfo)

		fmt.Println("--- SUMMARY ---")
		RunScriptWithString("/tmp/bit/git-extras/git-summary.sh", gitextras.GitSummary)

		fmt.Println("--- EFFORT ---")
		RunScriptWithString("/tmp/bit/git-extras/git-effort.sh", gitextras.GitEffort)
	},
	//Args: cobra.MaximumNArgs(1),
}

func init() {
	ShellCmd.AddCommand(infoCmd)
}
