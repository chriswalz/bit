package cmd

import (
	"fmt"
	"strings"

	"github.com/AlecAivazis/survey/v2"
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
		upstream := "origin"
		currentBranch := CurrentBranch()
		tbranch := currentBranch
		if len(args) == 1 {
			tbranch = args[0]
		}
		if len(args) == 2 {
			upstream = args[0]
			tbranch = args[1]
		}
		RunInTerminalWithColor("git", []string{"fetch", upstream, currentBranch})

		// if possibly squashed or rebase needed
		if IsDiverged() {
			RunInTerminalWithColor("git", []string{"status", "-sb", "--untracked-files=no"})

			ans := ""
			optionMap := map[string]string{
				"rebase": "Rebase on origin/upstream",
				"force":  "Force (destructive) push to " + upstream + "/" + currentBranch,
				"cancel": "Cancel",
			}
			prompt := &survey.Select{
				Message: "Branch is diverged from origin/upstream – handle by...",
				Options: []string{
					optionMap["rebase"],
					optionMap["force"],
					optionMap["cancel"],
				},
			}
			fmt.Println()
			survey.AskOne(prompt, &ans)
			if ans == optionMap["rebase"] {
				RunInTerminalWithColor("git", []string{"pull", "-r", upstream, currentBranch})
				return
			} else if ans == optionMap["force"] {
				RunInTerminalWithColor("git", []string{"push", upstream, currentBranch, "--force-with-lease"})
				// dont return user may have additional changes to save
			} else {
				fmt.Println("Canceling...")
				return
			}
		}
		if !CloudBranchExists() {
			RunInTerminalWithColor("git", []string{"push", "--set-upstream", upstream, currentBranch})
			save([]string{})
			RunInTerminalWithColor("git", []string{"push", upstream, currentBranch})
			return
		}
		save([]string{})
		if IsAheadOfCurrent() {
			RunInTerminalWithColor("git", []string{"push", upstream, currentBranch})
		} else {
			RunInTerminalWithColor("git", []string{"pull", "-r", upstream, currentBranch})
			RunInTerminalWithColor("git", []string{"push", upstream, currentBranch})
		}

		// After syncing with current branch and user wants to sync with another tbranch

		if currentBranch == "master" && !strings.HasSuffix(tbranch, "master") {
			yes := AskConfirm("Squash & merge " + args[0] + " into master?")

			if yes {
				RunInTerminalWithColor("git", []string{"merge", upstream, tbranch, "--squash"})
				return
			}
			fmt.Println("Cancelling...")
			// RunInTerminalWithColor("git", []string{"stash", "pop"}) deprecated switch stashing
			return
		}

		if tbranch == "master" {
			RunInTerminalWithColor("git", []string{"pull", "--rebase", upstream, tbranch})
			return
		}

		//TODO sync with another branch when not on master
		//
		//

		if len(args) == 1 {
			branch := args[0]
			refreshOnBranch(branch)
		}
	},
	// Args: cobra.MaximumNArgs(1),
}

func init() {
	BitCmd.AddCommand(syncCmd)
	// syncCmd.PersistentFlags().String("foo", "", "A help for foo")
	// syncCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
