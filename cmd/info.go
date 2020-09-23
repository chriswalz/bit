package cmd

import (
	"fmt"
	git_extras "github.com/chriswalz/bit/git-extras"
	"github.com/chriswalz/bit/util"
	"github.com/spf13/cobra"
	"os"
	"path/filepath"
)

// infoCmd represents the info command
var infoCmd = &cobra.Command{
	Use:   "info",
	Short: "Get general information about the status of your repository",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		os.MkdirAll(filepath.Dir("/tmp/bit/git-extras/"), os.ModePerm)
		fmt.Println("--- INFO ---")
		util.RunScriptWithString("/tmp/bit/git-extras/git-info.sh", git_extras.GitInfo)

		fmt.Println("--- SUMMARY ---")
		util.RunScriptWithString("/tmp/bit/git-extras/git-summary.sh", git_extras.GitSummary)

		fmt.Println("--- EFFORT ---")
		util.RunScriptWithString("/tmp/bit/git-extras/git-effort.sh", git_extras.GitEffort)
	},
	//Args: cobra.MaximumNArgs(1),
}

func init() {
	rootCmd.AddCommand(infoCmd)
}
