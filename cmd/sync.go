package cmd

import (
	"fmt"
	"github.com/chriswalz/bit/util"
	"github.com/spf13/cobra"
	"os/exec"
	"strings"
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
		// if possibly squashed
		// do nothing
		if !isBehindCurrent() && !isAheadOfCurrent() {
			return
		}
		save("")
		if cloudBranchExists() {
			util.Runwithcolor([]string{"pull", "-r"})
			if len(args) > 0 {
				util.Runwithcolor(append([]string{"pull", "-r"}, args...))
			}
			util.Runwithcolor([]string{"push"})
		} else {
			util.Runwithcolor([]string{"push", "--set-upstream", "origin", currentBranch()})
		}

	},
	//Args: cobra.MaximumNArgs(1),
}

func init() {
	rootCmd.AddCommand(syncCmd)
	// syncCmd.PersistentFlags().String("foo", "", "A help for foo")
	// syncCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func cloudBranchExists() bool {
	msg, err := exec.Command("git", "pull").CombinedOutput()
	if err != nil {
		fmt.Println(err)
	}
	//log.Println("msg:", string(msg))
	//log.Println("err:", err)
	return !strings.Contains(string(msg), "There is no tracking information for the current branch")
}

func currentBranch() string {
	msg, err := exec.Command("git", "branch", "--show-current").CombinedOutput()
	if err != nil {
		fmt.Println(err)
	}
	return strings.TrimSpace(string(msg))
}