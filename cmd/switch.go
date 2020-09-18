package cmd

import (
	"fmt"
	"github.com/chriswalz/bit/util"
	"github.com/spf13/cobra"
	"os/exec"
	"strings"
	"github.com/manifoldco/promptui"
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
			// select a branch
			prompt := promptui.Select{
				Label: "Select branch",
				Items: util.BranchList(),
			}

			_, result, err := prompt.Run()

			if err != nil {
				fmt.Printf("Prompt failed %v\n", err)
				return
			}

			branchName = strings.TrimSpace(result)
			fmt.Println(branchName+branchName)
		}

		if util.StashableChanges() {
			util.Runwithcolor([]string{"stash", "save", util.CurrentBranch() + "-automaticBitStash"})
		}
		util.Runwithcolor([]string{"pull", "--ff-only"})
		branchExists := checkoutBranch(branchName)
		if !branchExists {
			resp := util.PromptUser("Branch does not exist. Do you want to create it? Y/n")
			if util.IsYes(resp) {
				util.Runwithcolor([]string{"checkout", "-b", branchName})
				return
			}
			util.Runwithcolor([]string{"stash", "pop"})
			return
		}
		stashList := util.StashList()
		for _, stashLine := range stashList {
			if strings.Contains(stashLine, util.CurrentBranch() + "-automaticBitStash") {
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