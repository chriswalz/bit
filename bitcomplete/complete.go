// Package main is complete tool for the go command line
package main

import (
	"github.com/c-bata/go-prompt"
	"github.com/chriswalz/bit/cmd"
	"github.com/posener/complete/v2"
	"github.com/posener/complete/v2/predict"
	"github.com/thoas/go-funk"
	"strings"
)


func main() {

	branchCompletion := &complete.Command{
		Args: complete.PredictFunc(func(prefix string) []string {
			branches := cmd.BranchListSuggestions()
			completion := make([]string, len(branches))
			for i, v := range branches {
				completion[i] = v.Text
			}
			return completion
		}),
	}

	cmds := cmd.AllBitAndGitSubCommands(cmd.ShellCmd)
	completionSubCmdMap := map[string]*complete.Command{}
	for _, v := range cmds {
		flagSuggestions := append(cmd.FlagSuggestionsForCommand(v.Name(), "--"), cmd.FlagSuggestionsForCommand(v.Name(), "-")...)
		flags := funk.Map(flagSuggestions, func(x prompt.Suggest) (string, complete.Predictor) {
			if strings.HasPrefix(x.Text, "--") {
				return x.Text, predict.Nothing
			} else if strings.HasPrefix(x.Text, "-") {
				return x.Text, predict.Nothing
			} else {
				return "", predict.Nothing
			}
		})
		completionSubCmdMap[v.Name()] = &complete.Command{
			Flags: flags.(map[string]complete.Predictor),
		}
		if v.Name() == "checkout" || v.Name() == "co" || v.Name() == "switch" {
			branchCompletion.Flags = flags.(map[string]complete.Predictor)
			completionSubCmdMap[v.Name()] = branchCompletion
		}
	}

	gogo := &complete.Command{
		Sub: completionSubCmdMap,
	}

	gogo.Complete("bit")
}
