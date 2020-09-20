package cmd

import (
	"fmt"
	"github.com/c-bata/go-prompt"
	"github.com/chriswalz/bit/util"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
	"os"
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
			branchName = selectBranchPrompt()
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


func completer(d prompt.Document) []prompt.Suggest {
	list := util.BranchList()
	var suggestions []prompt.Suggest
	for _, branch := range list {
		suggestions = append(suggestions, prompt.Suggest{
			Text:        branch,
			Description: "",
		})
	}

	return prompt.FilterHasPrefix(suggestions, d.GetWordBeforeCursor(), true)
}

func selectBranchPrompt() string {
	//p := NewPrompt()
	// select a branch
	result := prompt.Input("Select a branch: ", completer,
		prompt.OptionTitle("sql-prompt"),
		prompt.OptionHistory([]string{""}),
		prompt.OptionPrefixTextColor(prompt.Yellow),
		prompt.OptionPreviewSuggestionTextColor(prompt.Yellow),
		prompt.OptionSelectedSuggestionBGColor(prompt.Yellow),
		prompt.OptionSuggestionBGColor(prompt.Yellow),
		prompt.OptionSelectedSuggestionTextColor(prompt.Purple),
		prompt.OptionShowCompletionAtStart(),
		prompt.OptionCompletionOnDown(),
		prompt.OptionSwitchKeyBindMode(prompt.CommonKeyBind),
		prompt.OptionAddKeyBind(prompt.KeyBind{
			Key: prompt.ControlC,
			Fn: func(b *prompt.Buffer) {
				os.Exit(0)
			},
		}),
	)
	//result := p.Input()

	branchName := strings.TrimSpace(result)
	return branchName
}

var LivePrefixState struct {
	LivePrefix string
	IsEnable   bool
}

func executor(in string) {
	fmt.Println("Your input: " + in)
	if in == "" {
		LivePrefixState.IsEnable = false
		LivePrefixState.LivePrefix = in
		return
	}
	LivePrefixState.LivePrefix = in + "> "
	LivePrefixState.IsEnable = true
}

func changeLivePrefix() (string, bool) {
	return LivePrefixState.LivePrefix, LivePrefixState.IsEnable
}

func NewPrompt() *prompt.Prompt {
	p := prompt.New(
		executor,
		completer,
		prompt.OptionPrefix("Branch: "),
		prompt.OptionLivePrefix(changeLivePrefix),
		prompt.OptionTitle("live-prefix-example"),
	)
	return p
}