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
		Runwithcolor("git", []string{"fetch"})

		// squash specific branch into current branch?
		if len(args) == 1 {
			branch := args[0]
			if branch == "master" {
				fmt.Println("Not supported")
				return
			}
			Runwithcolor("git", []string{"pull", "--ff-only"})
			Runwithcolor("git", []string{"merge", "--squash", branch})

		}
		// if possibly squashed
		if IsDiverged() {
			Runwithcolor("git", []string{"status", "-sb"})
			resp := PromptUser("Force (destructive) push to origin/" + CurrentBranch() + "? Y/n")
			if IsYes(resp) {
				Runwithcolor("git", []string{"push", "--force-with-lease"})
			}
			return
		}
		if !CloudBranchExists() {
			Runwithcolor("git", []string{"push", "--set-upstream", "origin", CurrentBranch()})
			save("")
			Runwithcolor("git", []string{"push"})
			return
		}
		save("")
		if !CloudBranchExists() {
			Runwithcolor("git", []string{"push", "--set-upstream", "origin", CurrentBranch()})
		}
		if IsAheadOfCurrent() {
			Runwithcolor("git", []string{"push"})
		} else {
			Runwithcolor("git", []string{"pull", "-r"})
			if len(args) > 0 {
				Runwithcolor("git", append([]string{"pull", "-r"}, args...))
			}
			Runwithcolor("git", []string{"push"})
		}

	},
	//Args: cobra.MaximumNArgs(1),
}

func init() {
	ShellCmd.AddCommand(syncCmd)
	// syncCmd.PersistentFlags().String("foo", "", "A help for foo")
	// syncCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
