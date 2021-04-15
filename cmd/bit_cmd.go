package cmd

import (
	"fmt"
	"github.com/chriswalz/complete/v3"
	"github.com/google/shlex"
	"github.com/lithammer/fuzzysearch/fuzzy"
	"os"
	"sort"
	"strings"

	"github.com/rs/zerolog/log"

	"github.com/c-bata/go-prompt"
	"github.com/spf13/cobra"
)

// BitCmd represents the base command when called without any subcommands
var BitCmd = &cobra.Command{
	Use:   "bit",
	Short: "Bit is a Git CLI that predicts what you want to do",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		suggestionTree, bitCmdMap := CreateSuggestionMap(cmd)

		repeat := strings.ToLower(os.Getenv("BIT_INTERACTIVE")) == "true"
		repeatAmount := 1
		if repeat {
			repeatAmount = 5000
		}

		for i := repeatAmount; i > 0; i-- {
			resp := SuggestionPrompt("> bit ", shellCommandCompleter(suggestionTree))
			subCommand := resp
			if subCommand == "" {
				return
			}
			if strings.Index(resp, " ") > 0 {
				subCommand = subCommand[0:strings.Index(resp, " ")]
			}
			parsedArgs, err := parseCommandLine(resp)
			if err != nil {
				log.Debug().Err(err).Send()
				continue
			}
			if bitCmdMap[subCommand] == nil {
				yes := HijackGitCommandOccurred(parsedArgs, suggestionTree, cmd.Version)
				if yes {
					continue
				}
				RunGitCommandWithArgs(parsedArgs)
				continue
			}

			cmd.SetArgs(parsedArgs)
			cmd.Execute()
		}
	},
}

func init() {
	BitCmd.PersistentFlags().Bool("debug", false, "Print debugging information")
}

// Execute adds all child commands to the shell command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the BitCmd.
func Execute() {
	if err := BitCmd.Execute(); err != nil {
		log.Info().Err(err)
		os.Exit(1)
	}
}

func shellCommandCompleter(suggestionTree *complete.CompTree) func(d prompt.Document) []prompt.Suggest {
	return func(d prompt.Document) []prompt.Suggest {
		return promptCompleter(suggestionTree, d.Text)
	}
}

func branchCommandCompleter(suggestionMap *complete.CompTree) func(d prompt.Document) []prompt.Suggest {
	return func(d prompt.Document) []prompt.Suggest {
		return promptCompleter(suggestionMap, "checkout "+d.Text)
	}
}

func specificCommandCompleter(subCmd string, suggestionMap *complete.CompTree) func(d prompt.Document) []prompt.Suggest {
	return func(d prompt.Document) []prompt.Suggest {
		return promptCompleter(suggestionMap, subCmd+" "+d.Text)
	}
}

var fuzzyQuery = func(s, prefix string) bool {
	return fuzzy.Match(prefix, s)
}

func promptCompleter(suggestionTree *complete.CompTree, text string) []prompt.Suggest {
	text = "bit " + text

	var sugg []prompt.Suggest

	queryFunc := strings.HasPrefix

	split, err := shlex.Split(strings.TrimSpace(text))
	if err != nil {
		log.Debug().Err(err).Send()
		return sugg
	}
	lastToken := split[len(split)-1]
	lastCommand := lastToken
	if !strings.HasSuffix(text, " ") && len(split) >= 2 {
		lastCommand = split[len(split)-2]
	}

	// use fuzzy search completion when querying branch names
	if isBranchCompletionCommand(lastCommand) {
		queryFunc = fuzzyQuery
	}

	suggestions, err := complete.CompleteLine(text, suggestionTree, queryFunc)
	if err != nil {
		log.Debug().Err(err).Send()
		return sugg
	}

	// for branches dont undo most recent sorts with alphabetical sort
	if !isBranchCompletionCommand(lastCommand) {
		sort.Slice(suggestions, func(i, j int) bool {
			return suggestions[i].Name < suggestions[j].Name
		})
	}

	for _, suggestion := range suggestions {
		name := suggestion.Name
		if strings.HasPrefix(lastToken, "-") && !strings.HasSuffix(text, " ") {
			name = "-" + suggestion.Name
			if len(name) > 2 {
				name = "-" + name
			}
		}
		sugg = append(sugg, prompt.Suggest{
			Text:        name,
			Description: suggestion.Desc,
		})
	}

	if text == "bit " {
		sugg = append(CobraCommandToSuggestions(CommonCommandsList()), sugg...)
	}

	return prompt.FilterHasPrefix(sugg, "", true)
}

func RunGitCommandWithArgs(args []string) {
	var err error
	err = RunInTerminalWithColor("git", args)
	if err != nil {
		log.Debug().Msg("Command may not exist: " + err.Error())
	}
	return
}

func HijackGitCommandOccurred(args []string, suggestionMap *complete.CompTree, version string) bool {
	sub := args[0]
	// handle checkout,switch,co commands as checkout
	// if "-b" flag is not provided and branch does not exist
	// user would be prompted asking whether to create a branch or not
	// expected usage format
	//   bit (checkout|switch|co) [-b] branch-name
	if args[len(args)-1] == "--version" || args[len(args)-1] == "version" {
		fmt.Println("bit version " + version)
		return false
	}
	if sub == "pr" {
		runPr(suggestionMap)
		return true
	}
	if sub == "merge" && len(args) == 1 {
		branchName := SuggestionPrompt("> bit "+sub+" ", specificCommandCompleter("merge", suggestionMap))
		RunInTerminalWithColor("git", []string{"merge", branchName})
		return true
	}
	if isBranchChangeCommand(sub) {
		branchName := ""
		if len(args) < 2 {
			branchName = SuggestionPrompt("> bit "+sub+" ", branchCommandCompleter(suggestionMap))
		} else {
			branchName = strings.TrimSpace(args[len(args)-1])
		}

		if strings.HasPrefix(branchName, "origin/") {
			branchName = branchName[7:]
		}
		args[len(args)-1] = branchName
		var createBranch bool
		if len(args) == 3 && args[len(args)-2] == "-b" {
			createBranch = true
		}
		branchExists := checkoutBranch(branchName)
		if branchExists {
			refreshBranch()
			return true
		}

		if !createBranch && !AskConfirm("Branch does not exist. Do you want to create it?") {
			fmt.Printf("Cancelling...")
			return true
		}

		RunInTerminalWithColor("git", []string{"checkout", "-b", branchName})
		return true
	}
	return false
}

func GetVersion() string {
	return BitCmd.Version
}
