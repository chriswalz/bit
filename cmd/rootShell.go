package cmd

import (
	"errors"
	"fmt"
	"github.com/c-bata/go-prompt"
	"github.com/chriswalz/bit/util"
	"github.com/spf13/cobra"
	"os"
	"strings"
)

var cfgFile string

// shellCmd represents the base command when called without any subcommands
var shellCmd = &cobra.Command{
	Use:   "bit",
	Short: "Bit is Git with a simple interface. Plus you can still use all the old git commands",
	Long:  `v0.3.11`,
	ValidArgs: []string{"checkout", "status", "stash"},
	//ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	//	//gitCmds := util.AllGitSubCommands()
	//	return []string{"checkout", "status", "stash"}, cobra.ShellCompDirectiveDefault
	//},
	Run: func(cmd *cobra.Command, args []string) {
		gitCmds := util.AllGitSubCommands()
		bitCmds := cmd.Commands()
		bitCmdMap := map[string]*cobra.Command{}
		for _, bitCmd := range bitCmds {
			bitCmdMap[bitCmd.Name()] = bitCmd
		}
		completerSuggestionMap := map[string][]prompt.Suggest{
			"":         {},
			"shell":    util.CobraCommandToSuggestions(append(gitCmds, bitCmds...)),
			"checkout": util.BranchListSuggestions(),
			"switch":   util.BranchListSuggestions(),
			"add":      util.GitAddSuggestions(),
			"release": {
				{Text: "bump", Description: "Increment SemVer from tags and release"},
				{Text: "<version>", Description: "Name of release version e.g. v0.1.2"},
			},
		}
		resp := util.SuggestionPrompt("bit ", shellCommandCompleter(completerSuggestionMap))
		subCommand := resp
		if strings.Index(resp, " ") > 0 {
			subCommand = subCommand[0:strings.Index(resp, " ")]
		}
		if bitCmdMap[subCommand] == nil {
			parsedArgs, err := parseCommandLine(resp)
			if err != nil {
				fmt.Println(err)
				return
			}
			err = util.Runwithcolor("git", parsedArgs)
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
// This is called by main.main(). It only needs to happen once to the shellCmd.
func Execute() {
	if err := shellCmd.Execute(); err != nil {
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
			prev := filterFlags[0] // git command or sub command (not a flag)
			curr := filterFlags[1] // current argument or flag
			if strings.HasPrefix(curr, "--") {
				suggestions = util.FlagSuggestionsForCommand(prev, "--")
			} else if strings.HasPrefix(curr, "-") {
				suggestions = util.FlagSuggestionsForCommand(prev, "-")
			} else if suggestionMap[prev] != nil {
				suggestions = suggestionMap[prev]
			}

		}
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
		return []string{}, errors.New(fmt.Sprintf("Unclosed quote in command line: %s", command))
	}

	if current != "" {
		args = append(args, current)
	}

	return args, nil
}
