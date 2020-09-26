package cmd

import (
	"fmt"
	"github.com/chriswalz/bit/util"
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
		util.Runwithcolor("git",[]string{"fetch"})

		// squash specific branch into current branch?
		if len(args) == 1 {
			branch := args[0]
			if branch == "master" {
				fmt.Println("Not supported")
				return
			}
			util.Runwithcolor("git",[]string{"pull", "--ff-only"})
			util.Runwithcolor("git",[]string{"merge", "--squash", branch})

		}
		// if possibly squashed
		if util.IsDiverged() {
			util.Runwithcolor("git",[]string{"status", "-sb"})
			resp := util.PromptUser("Force (destructive) push to origin/" + util.CurrentBranch() + "? Y/n")
			if util.IsYes(resp) {
				util.Runwithcolor("git",[]string{"push", "--force-with-lease"})
			}
			return
		}
		if !util.CloudBranchExists() {
			util.Runwithcolor("git",[]string{"push", "--set-upstream", "origin", util.CurrentBranch()})
			save("")
			util.Runwithcolor("git",[]string{"push"})
			return
		}
		save("")
		if !util.CloudBranchExists() {
			util.Runwithcolor("git",[]string{"push", "--set-upstream", "origin", util.CurrentBranch()})
		}
		if util.IsAheadOfCurrent() {
			util.Runwithcolor("git",[]string{"push"})
		} else {
			util.Runwithcolor("git",[]string{"pull", "-r"})
			if len(args) > 0 {
				util.Runwithcolor("git",append([]string{"pull", "-r"}, args...))
			}
			util.Runwithcolor("git",[]string{"push"})
		}

	},
	//Args: cobra.MaximumNArgs(1),
}

func init() {
	shellCmd.AddCommand(syncCmd)
	// syncCmd.PersistentFlags().String("foo", "", "A help for foo")
	// syncCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
