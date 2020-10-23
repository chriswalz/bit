package cmd

import (
	"fmt"
	"github.com/rs/zerolog/log"
	"os"
	"strings"
	"time"

	"github.com/c-bata/go-prompt"
	"github.com/spf13/cobra"
)

// ShellCmd represents the base command when called without any subcommands
var ShellCmd = &cobra.Command{
	Use:   "bit",
	Short: "Bit is a Git CLI that predicts what you want to do",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		completerSuggestionMap, bitCmdMap := CreateSuggestionMap(cmd)

		resp := SuggestionPrompt("> bit ", shellCommandCompleter(completerSuggestionMap))
		subCommand := resp
		if subCommand == "" {
			return
		}
		if strings.Index(resp, " ") > 0 {
			subCommand = subCommand[0:strings.Index(resp, " ")]
		}
		parsedArgs, err := parseCommandLine(resp)
		if err != nil {
			log.Debug().Err(err)
			return
		}
		if bitCmdMap[subCommand] == nil {
			yes := GitCommandsPromptUsed(parsedArgs, completerSuggestionMap, cmd.Version)
			if yes {
				return
			}
			RunGitCommandWithArgs(parsedArgs)
			return
		}

		cmd.SetArgs(parsedArgs)
		cmd.Execute()
	},
}

func init() {
	ShellCmd.PersistentFlags().Bool("debug", false, "Print debugging information")
}

func CreateSuggestionMap(cmd *cobra.Command) (map[string]func() []prompt.Suggest, map[string]*cobra.Command) {
	start := time.Now()
	_, bitCmdMap := AllBitSubCommands(cmd)
	log.Debug().Msg((time.Now().Sub(start)).String())
	start = time.Now()
	allBitCmds := AllBitAndGitSubCommands(cmd)
	log.Debug().Msg((time.Now().Sub(start)).String())
	//commonCommands := CobraCommandToSuggestions(CommonCommandsList())
	start = time.Now()
	branchListSuggestions := BranchListSuggestions()
	log.Debug().Msg((time.Now().Sub(start)).String())
	start = time.Now()
	combraCommandSuggestions := CobraCommandToSuggestions(allBitCmds)
	log.Debug().Msg((time.Now().Sub(start)).String())
	start = time.Now()
	gitAddSuggestions := GitAddSuggestions()
	log.Debug().Msg((time.Now().Sub(start)).String())
	start = time.Now()
	gitResetSuggestions := GitResetSuggestions()
	log.Debug().Msg((time.Now().Sub(start)).String())
	start = time.Now()

	completerSuggestionMap := map[string]func() []prompt.Suggest{
		"":         memoize([]prompt.Suggest{}),
		"shell":    memoize(combraCommandSuggestions),
		"checkout": memoize(branchListSuggestions),
		"switch":   memoize(branchListSuggestions),
		"co":       memoize(branchListSuggestions),
		"merge":    memoize(branchListSuggestions),
		"rebase":   memoize(branchListSuggestions),
		"log":      memoize(branchListSuggestions),
		"add":      memoize(gitAddSuggestions),
		"release": memoize([]prompt.Suggest{
			{Text: "bump", Description: "Increment SemVer from tags and release e.g. if latest is v0.1.2 it's bumped to v0.1.3 "},
			{Text: "<version>", Description: "Name of release version e.g. v0.1.2"},
		}),
		"reset": memoize(gitResetSuggestions),
		"pr": lazyLoad(GitHubPRSuggestions),
		//"_any": commonCommands,
	}
	return completerSuggestionMap, bitCmdMap

}

// Execute adds all child commands to the shell command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the ShellCmd.
func Execute() {
	if err := ShellCmd.Execute(); err != nil {
		log.Info().Err(err)
		os.Exit(1)
	}
}

func shellCommandCompleter(suggestionMap map[string]func() []prompt.Suggest) func(d prompt.Document) []prompt.Suggest {
	return func(d prompt.Document) []prompt.Suggest {
		return promptCompleter(suggestionMap, d.Text)
	}
}

func branchCommandCompleter(suggestionMap map[string]func() []prompt.Suggest) func(d prompt.Document) []prompt.Suggest {
	return func(d prompt.Document) []prompt.Suggest {
		return promptCompleter(suggestionMap, "checkout "+d.Text)
	}
}

func prCommandCompleter(suggestionMap map[string]func() []prompt.Suggest) func(d prompt.Document) []prompt.Suggest {
	return func(d prompt.Document) []prompt.Suggest {
		return promptCompleter(suggestionMap, "pr "+d.Text)
	}
}

func promptCompleter(suggestionMap map[string]func() []prompt.Suggest, text string) []prompt.Suggest {
	var suggestions []prompt.Suggest
	split := strings.Split(text, " ")
	filterFlags := make([]string, 0, len(split))
	for i, v := range split {
		if !strings.HasPrefix(v, "-") || i == len(split)-1 {
			filterFlags = append(filterFlags, v)
		}
	}
	prev := filterFlags[0] // in git commit -m "hello"  commit is prev
	if len(prev) == len(text) {
		suggestions = suggestionMap["shell"]()
		return prompt.FilterHasPrefix(suggestions, prev, true)
	}
	curr := filterFlags[1] // in git commit -m "hello"  "hello" is curr
	if strings.HasPrefix(curr, "--") {
		suggestions = FlagSuggestionsForCommand(prev, "--")
	} else if strings.HasPrefix(curr, "-") {
		suggestions = FlagSuggestionsForCommand(prev, "-")
	} else if suggestionMap[prev] != nil {
		suggestions = suggestionMap[prev]()
		if isBranchCompletionCommand(prev) {
			return prompt.FilterContains(suggestions, curr, true)
		}
	}
	return prompt.FilterHasPrefix(suggestions, curr, true)
}

func RunGitCommandWithArgs(args []string) {
	var err error
	err = RunInTerminalWithColor("git", args)
	if err != nil {
		log.Debug().Msg("Command may not exist: " + err.Error())
	}
	return
}

func GitCommandsPromptUsed(args []string, suggestionMap map[string]func() []prompt.Suggest, version string) bool {
	sub := args[0]
	// handle checkout,switch,co commands as checkout
	// if "-b" flag is not provided and branch does not exist
	// user would be prompted asking whether to create a branch or not
	// expected usage format
	//   bit (checkout|switch|co) [-b] branch-name
	if args[len(args)-1] == "--version" {
		fmt.Println("bit version " + version)
	}
	if isBranchChangeCommand(sub) {
		branchName := ""
		if sub == "pr" {
			runPr(suggestionMap)
			return true
		} else if len(args) < 2 {
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
	return ShellCmd.Version
}