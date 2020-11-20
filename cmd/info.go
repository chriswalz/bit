package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/chriswalz/bit/gitextras"
	"github.com/spf13/cobra"
)

// infoCmd represents the info command
var infoCmd = &cobra.Command{
	Use:   "info",
	Short: "Get general information about the status of your repository",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		os.MkdirAll(filepath.Dir("/tmp/bit/git-extras/"), os.ModePerm)
		fmt.Println("--- INFO ---")
		RunInTerminalWithColor("/bin/sh", []string{`-c`, gitextras.GitInfo})

		fmt.Println("--- SUMMARY ---")
		RunInTerminalWithColor("/bin/sh", []string{`-c`, gitextras.GitSummary})

		fmt.Println("\n--- EFFORT ---")
		fmt.Println("\nCommits | Files")
		RunInTerminalWithColor("/bin/sh", []string{`-c`, `git log --pretty=format: --name-only | sort | uniq -c | sort -rg | awk 'NR > 1 { print }' | head -15`})
	},
	// Args: cobra.MaximumNArgs(1),
}

func init() {
	BitCmd.AddCommand(infoCmd)
}
