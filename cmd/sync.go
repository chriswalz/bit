package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

// syncCmd represents the sync command
var syncCmd = &cobra.Command{
	Use:   "sync",
	Short: "Synchronizes local changes with changes on origin or specified branch",
	Long: `sync
sync origin master
sync local-branch
`,
	Run: func(cmd *cobra.Command, args []string) {
		RunInTerminalWithColor("git", []string{"fetch"})

		// squash specific branch into current branch?
		if len(args) == 1 {
			branch := args[0]
			if branch == "master" {
				fmt.Println("Not supported")
				return
			}
			RunInTerminalWithColor("git", []string{"pull", "--ff-only"})
			RunInTerminalWithColor("git", []string{"merge", "--squash", branch})

		}
		// if possibly squashed
		if IsDiverged() {
			RunInTerminalWithColor("git", []string{"status", "-sb"})
			resp := PromptUser("Force (destructive) push to origin/" + CurrentBranch() + "? Y/n")
			if IsYes(resp) {
				RunInTerminalWithColor("git", []string{"push", "--force-with-lease"})
			}
			return
		}
		if !CloudBranchExists() {
			RunInTerminalWithColor("git", []string{"push", "--set-upstream", "origin", CurrentBranch()})
			save("")
			RunInTerminalWithColor("git", []string{"push"})
			return
		}
		save("")
		if !CloudBranchExists() {
			RunInTerminalWithColor("git", []string{"push", "--set-upstream", "origin", CurrentBranch()})
		}
		if IsAheadOfCurrent() {
			RunInTerminalWithColor("git", []string{"push"})
		} else {
			RunInTerminalWithColor("git", []string{"pull", "-r"})
			if len(args) > 0 {
				RunInTerminalWithColor("git", append([]string{"pull", "-r"}, args...))
			}
			RunInTerminalWithColor("git", []string{"push"})
		}

	},
	//Args: cobra.MaximumNArgs(1),
}

func init() {
	ShellCmd.AddCommand(syncCmd)
	// syncCmd.PersistentFlags().String("foo", "", "A help for foo")
	// syncCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
