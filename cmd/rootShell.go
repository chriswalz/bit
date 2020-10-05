package cmd

import (
	"fmt"
	"github.com/c-bata/go-prompt"
	"github.com/spf13/cobra"
	"os"
	"strings"
)

var cfgFile string

// ShellCmd represents the base command when called without any subcommands
var ShellCmd = &cobra.Command{
	Use:   "bit",
	Short: "Bit is a Git CLI that predicts what you want to do",
	Long:  `v0.3.16`,
	Run: func(cmd *cobra.Command, args []string) {
		_, bitCmdMap := AllBitSubCommands(cmd)
		allBitCmds := AllBitAndGitSubCommands(cmd)
		//commonCommands := CobraCommandToSuggestions(CommonCommandsList())
		branchListSuggestions := BranchListSuggestions()
		completerSuggestionMap := map[string][]prompt.Suggest{
			"":         {},
			"shell":    CobraCommandToSuggestions(allBitCmds),
			"checkout": branchListSuggestions,
			"switch":   branchListSuggestions,
			"co":       branchListSuggestions,
			"merge":    branchListSuggestions,
			"add":      GitAddSuggestions(),
			"release": {
				{Text: "bump", Description: "Increment SemVer from tags and release"},
				{Text: "<version>", Description: "Name of release version e.g. v0.1.2"},
			},
			"reset": GitResetSuggestions(),
			//"_any": commonCommands,
		}
		//if isGitRepo() {
		//	fmt.Println("fatal: not a git repository (or any of the parent directories): .git")
		//	return
		//}

		resp := SuggestionPrompt("> bit ", shellCommandCompleter(completerSuggestionMap))
		subCommand := resp
		if subCommand == "" {
			return
		}
		if strings.Index(resp, " ") > 0 {
			subCommand = subCommand[0:strings.Index(resp, " ")]
		}
		if bitCmdMap[subCommand] == nil {
			parsedArgs, err := parseCommandLine(resp)
			if err != nil {
				fmt.Println(err)
				return
			}
			sub := parsedArgs[0]
			if sub == "checkout" || sub == "switch" || sub == "co" {
				branchName := strings.TrimSpace(parsedArgs[len(parsedArgs)-1])
				if strings.HasPrefix(branchName, "origin/") {
					branchName = branchName[7:]
				}
				parsedArgs[len(parsedArgs)-1] = branchName
			}
			err = Runwithcolor("git", parsedArgs)
			if err != nil {
				fmt.Println("DEBUG: CMD may not be allow listed")
			}
			return
		}
		parsedArgs, err := parseCommandLine(resp)
		if err != nil {
			fmt.Println(err)
			return
		}
		cmd.SetArgs(parsedArgs)
		cmd.Execute()
	},
}

// Execute adds all child commands to the shell command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the ShellCmd.
func Execute() {
	if err := ShellCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func shellCommandCompleter(suggestionMap map[string][]prompt.Suggest) func(d prompt.Document) []prompt.Suggest {
	return func(d prompt.Document) []prompt.Suggest {
		//fmt.Println(d.GetWordBeforeCursor())
		// only 1 command
		var suggestions []prompt.Suggest
		if len(d.GetWordBeforeCursor()) == len(d.Text) {
			//fmt.Println("same")
			suggestions = suggestionMap["shell"]
		} else {
			split := strings.Split(d.Text, " ")
			filterFlags := make([]string, 0, len(split))
			for i, v := range split {
				if !strings.HasPrefix(v, "-") || i == len(split)-1 {
					filterFlags = append(filterFlags, v)
				}
			}
			prev := filterFlags[0] // in git commit -m "hello"  commit is prev
			curr := filterFlags[1] // in git commit -m "hello"  "hello" is curr
			if strings.HasPrefix(curr, "--") {
				suggestions = FlagSuggestionsForCommand(prev, "--")
			} else if strings.HasPrefix(curr, "-") {
				suggestions = FlagSuggestionsForCommand(prev, "-")
			} else if suggestionMap[prev] != nil {
				suggestions = suggestionMap[prev]
			}
		}
		//suggestions = append(suggestionMap["_any"], suggestions...)
		return prompt.FilterContains(suggestions, d.GetWordBeforeCursor(), true)
	}
}

func parseCommandLine(command string) ([]string, error) {
	var args []string
	state := "start"
	current := ""
	quote := "\""
	escapeNext := true
	for i := 0; i < len(command); i++ {
		c := command[i]

		if state == "quotes" {
			if string(c) != quote {
				current += string(c)
			} else {
				args = append(args, current)
				current = ""
				state = "start"
			}
			continue
		}

		if escapeNext {
			current += string(c)
			escapeNext = false
			continue
		}

		if c == '\\' {
			escapeNext = true
			continue
		}

		if c == '"' || c == '\'' {
			state = "quotes"
			quote = string(c)
			continue
		}

		if state == "arg" {
			if c == ' ' || c == '\t' {
				args = append(args, current)
				current = ""
				state = "start"
			} else {
				current += string(c)
			}
			continue
		}

		if c != ' ' && c != '\t' {
			state = "arg"
			current += string(c)
		}
	}

	if state == "quotes" {
		return []string{}, fmt.Errorf("Unclosed quote in command line: %s", command)
	}

	if current != "" {
		args = append(args, current)
	}

	return args, nil
}
