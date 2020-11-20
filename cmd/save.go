package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

// saveCmd represents the save command
var saveCmd = &cobra.Command{
	Use:   "save",
	Short: "Save your changes to your current branch",
	Long:  `E.g. bit save; bit save "commit message"`,
	Run: func(cmd *cobra.Command, args []string) {
		msg := ""
		if len(args) > 0 {
			msg = strings.Join(args, " ")
		}
		save(msg)
	},
	// Args: cobra.MaximumNArgs(1),
}

// add comment

func init() {
	BitCmd.AddCommand(saveCmd)
	// saveCmd.PersistentFlags().String("foo", "", "A help for foo")
	// saveCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func save(msg string) {
	if NothingToCommit() {
		fmt.Println("nothing to save or commit")
		return
	}

	if msg == "" {
		// if ahead of master
		if IsAheadOfCurrent() || !CloudBranchExists() {
			RunInTerminalWithColor("git", []string{"commit", "-a", "--amend", "--no-edit"}) // amend if already ahead
		} else {
			RunInTerminalWithColor("git", []string{"status", "-sb", "--untracked-files=no"})
			resp := AskMultiLine("Please provide a description of your changes")
			RunInTerminalWithColor("git", []string{"commit", "-a", "-m " + resp})
		}
	} else {
		RunInTerminalWithColor("git", []string{"commit", "-a", "-m" + msg})
	}
}
