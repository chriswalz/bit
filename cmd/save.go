package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"strings"
)

// saveCmd represents the save command
var saveCmd = &cobra.Command{
	Use:   "save",
	Short: "Save your changes to your current branch",
	Long:  `E.g. bit save; bit save "commit message"`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) >= 1 && !strings.HasPrefix(args[0], "-") {
			args = append([]string{"-m", args[0]}, args[1:]...)
		}
		save(args)
	},
	Args:               cobra.MinimumNArgs(0),
	DisableFlagParsing: true,
}

// add comment

func init() {
	BitCmd.AddCommand(saveCmd)
	// saveCmd.PersistentFlags().String("foo", "", "A help for foo")
	// saveCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func getFlagValue(args []string, flagName string) *string {
	for i, arg := range args {
		if arg == flagName {
			return &args[i+1]
		}
	}
	return nil

}

func save(args []string) {
	if NothingToCommit() {
		fmt.Println("nothing to save or commit")
		return
	}
	msg := getFlagValue(args, "-m")
	if msg == nil {
		getFlagValue(args, "--message")
	}

	if msg == nil {
		// if ahead of master
		if IsAheadOfCurrent() || !CloudBranchExists() {
			// amend if already ahead
			extraArgs := []string{"-a", "--amend", "--no-edit"}
			RunInTerminalWithColor("git", append(append([]string{"commit"}, args...), extraArgs...))
		} else {
			RunInTerminalWithColor("git", []string{"status", "-sb", "--untracked-files=no"})
			resp := AskMultiLine("Please provide a description of your changes")
			extraArgs := []string{"-a", "-m " + resp}
			RunInTerminalWithColor("git", append(append([]string{"commit"}, args...), extraArgs...))
		}
	} else {
		extraArgs := []string{"-a"}
		RunInTerminalWithColor("git", append(append([]string{"commit"}, args...), extraArgs...))
	}
}
