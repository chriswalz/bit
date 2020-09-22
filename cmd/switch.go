package cmd

import (
	"fmt"
	"github.com/c-bata/go-prompt"
	"github.com/chriswalz/bit/util"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
	"os/exec"
	"strings"
)

// switchCmd represents the switch command
var switchCmd = &cobra.Command{
	Use:   "switch",
	Short: "bit switch [branch-name]",
	Long: `For existing branches simply run bit switch [branch-name]. 

For creating a new branch it's the same command! You'll simply be prompted to confirm that you want to create a new branch
`,
	Args: cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		util.Runwithcolor([]string{"fetch"})

		branchName := ""
		if len(args) >= 1 {
			branchName = args[0]
		}
		if len(args) == 0 {
			branchName = util.SuggestionPrompt("bit switch ", branchCompleter(util.BranchListSuggestions()))
			if strings.Contains(branchName, "*") {
				return
			}
		}

		if util.StashableChanges() {
			util.Runwithcolor([]string{"stash", "save", util.CurrentBranch() + "-automaticBitStash"})
		}
		util.Runwithcolor([]string{"pull", "--ff-only"})
		branchExists := checkoutBranch(branchName)
		if !branchExists {
			prompt := promptui.Prompt{
				Label:     "Branch does not exist. Do you want to create it?",
				IsConfirm: true,
			}

			_, err := prompt.Run()

			if err != nil {
				fmt.Printf("Cancelling...")
				util.Runwithcolor([]string{"stash", "pop"})
				return
			}

			util.Runwithcolor([]string{"checkout", "-b", branchName})
			return
		}
		stashList := util.StashList()
		for _, stashLine := range stashList {
			if strings.Contains(stashLine, util.CurrentBranch()+"-automaticBitStash") {
				stashId := strings.Split(stashLine, ":")[0]
				util.Runwithcolor([]string{"stash", "pop", stashId})
				return

			}
		}
		util.Runwithcolor([]string{"pull", "--ff-only"})
	},
}

func init() {
	rootCmd.AddCommand(switchCmd)
	// switchCmd.PersistentFlags().String("foo", "", "A help for foo")
	// switchCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func checkoutBranch(branch string) bool {
	msg, err := exec.Command("git", "checkout", branch).CombinedOutput()
	if err != nil {
		//fmt.Println(err)
	}
	return !strings.Contains(string(msg), "did not match any file")
}

func branchCompleter(branches []prompt.Suggest) func(d prompt.Document) []prompt.Suggest {
	return func(d prompt.Document) []prompt.Suggest {
		return prompt.FilterHasPrefix(branches, d.GetWordBeforeCursor(), true)
	}
}

