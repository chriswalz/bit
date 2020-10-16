package cmd

import (
	"fmt"
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
		upstream := "origin"
		tbranch := CurrentBranch()
		if len(args) == 1 {
			tbranch = args[0]
		}
		if len(args) == 2 {
			upstream = args[0]
			tbranch = args[1]
		}
		RunInTerminalWithColor("git", []string{"fetch", upstream, tbranch})

		// if possibly squashed
		if IsDiverged() {
			RunInTerminalWithColor("git", []string{"status", "-sb", "--untracked-files=no"})
			yes := AskConfirm("Force (destructive) push to " + upstream + "/" + tbranch + "?")
			if yes {
				RunInTerminalWithColor("git", []string{"push", upstream, tbranch, "--force-with-lease"})
			}
			return
		}
		if !CloudBranchExists() {
			RunInTerminalWithColor("git", []string{"push", "--set-upstream", upstream, tbranch})
			save("")
			RunInTerminalWithColor("git", []string{"push", upstream, tbranch})
			return
		}
		save("")
		if IsAheadOfCurrent() {
			RunInTerminalWithColor("git", []string{"push", upstream, tbranch})
		} else {
			RunInTerminalWithColor("git", []string{"pull", "-r", upstream, tbranch})
			RunInTerminalWithColor("git", []string{"push", upstream, tbranch})
		}

		// After syncing with current tbranch and user wants to sync with another tbranch

		if CurrentBranch() == "master" && !strings.HasSuffix(tbranch, "master") {
			yes := AskConfirm("Squash & merge " + args[0] +" into master?")

			if yes {
				RunInTerminalWithColor("git", []string{"merge", upstream, tbranch, "--squash"})
				return
			}
			fmt.Printf("Cancelling...")
			//RunInTerminalWithColor("git", []string{"stash", "pop"}) deprecated switch stashing
			return
		}

		if len(args) == 1 {
			branch := args[0]
			refreshOnBranch(branch)
		}

	},
	Args: cobra.MaximumNArgs(2),
}

func init() {
	ShellCmd.AddCommand(syncCmd)
	// syncCmd.PersistentFlags().String("foo", "", "A help for foo")
	// syncCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
