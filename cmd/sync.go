package cmd

import (
	"fmt"
	"github.com/chriswalz/bit/util"
	"github.com/spf13/cobra"
)

// syncCmd represents the sync command
var syncCmd = &cobra.Command{
	Use:   "sync",
	Short: "Catches up to another branch (Rebasing) and updates cloud branch with your changes",
	Long: `sync
sync origin master
sync local-branch
`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("sync")
		util.Runwithcolor([]string{"fetch"})
		// if possibly squashed
		if util.IsDiverged() {
			resp := util.PromptUser("Force (destructive) push to origin/" + util.CurrentBranch() + "? Y/n")
			if util.IsYes(resp) {
				fmt.Println("[implement force push]")
			}

			return
		}
		save("")
		if util.CloudBranchExists() {
			util.Runwithcolor([]string{"pull", "-r"})
			if len(args) > 0 {
				util.Runwithcolor(append([]string{"pull", "-r"}, args...))
			}
			util.Runwithcolor([]string{"push"})
		} else {
			util.Runwithcolor([]string{"push", "--set-upstream", "origin", util.CurrentBranch()})
		}

	},
	//Args: cobra.MaximumNArgs(1),
}

func init() {
	rootCmd.AddCommand(syncCmd)
	// syncCmd.PersistentFlags().String("foo", "", "A help for foo")
	// syncCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
