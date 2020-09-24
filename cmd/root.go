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

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "bit",
	Short: "Bit is Git with a simple interface. Plus you can still use all the old git commands",
	Long:  `v0.3.3`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		gitCmds := util.AllGitSubCommands()
		bitCmds := cmd.Commands()
		bitCmdMap := map[string]*cobra.Command{}
		for _, bitCmd := range bitCmds {
			bitCmdMap[bitCmd.Name()] = bitCmd
		}
		completerSuggestionMap := map[string][]prompt.Suggest{
			"":         {},
			"root":     util.CobraCommandToSuggestions(append(gitCmds, bitCmds...)),
			"checkout": util.BranchListSuggestions(),
			"switch":   util.BranchListSuggestions(),
			"add":      util.GitAddSuggestions(),
			"release": {
				{Text: "bump", Description: "Increment SemVer from tags and release"},
				{Text: "<version>", Description: "Name of release version e.g. v0.1.2"},
						},
		}
		resp := util.SuggestionPrompt("bit ", rootCommandCompleter(completerSuggestionMap))
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
			err = util.Runwithcolor("git",parsedArgs)
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

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func rootCommandCompleter(suggestionMap map[string][]prompt.Suggest) func(d prompt.Document) []prompt.Suggest {
	return func(d prompt.Document) []prompt.Suggest {
		//fmt.Println(d.GetWordBeforeCursor())
		// only 1 command
		var suggestions []prompt.Suggest
		if len(d.GetWordBeforeCursor()) == len(d.Text) {
			//fmt.Println("same")
			suggestions = suggestionMap["root"]
		} else {
			split := strings.Split(d.Text, " ")
			prev := split[0]
			curr := split[1]
			if strings.HasPrefix(curr, "--") {
				suggestions = util.FlagSuggestions(prev, "--")
			} else if strings.HasPrefix(curr, "-") {
				suggestions = util.FlagSuggestions(prev, "-")
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
