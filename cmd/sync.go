package cmd

import (
	"fmt"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
	"strings"
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

		// After syncing with current branch and user wants to sync with another branch

		if CurrentBranch() == "master" && len(args) == 1 && strings.HasSuffix(args[0], "master"){
			prompt := promptui.Prompt{
				Label:     "Squash & merge this branch into master",
				IsConfirm: true,
			}

			_, err := prompt.Run()

			if err != nil {
				fmt.Printf("Cancelling...")
				RunInTerminalWithColor("git", []string{"stash", "pop"})
				return
			}
			RunInTerminalWithColor("git", []string{"merge", "--squash"})
			return
		}

		if len(args) == 1 {
			branch := args[0]
			refreshOnBranch(branch)
		}

	},
	//Args: cobra.MaximumNArgs(1),
}

func init() {
	ShellCmd.AddCommand(syncCmd)
	// syncCmd.PersistentFlags().String("foo", "", "A help for foo")
	// syncCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
