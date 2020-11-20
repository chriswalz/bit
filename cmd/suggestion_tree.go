package cmd

import (
	"github.com/c-bata/go-prompt"
	"github.com/chriswalz/complete/v2"
	"github.com/chriswalz/complete/v2/predict"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"github.com/thoas/go-funk"
	"time"
)

func CreateSuggestionMap(cmd *cobra.Command) (*complete.Command, map[string]*cobra.Command) {
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
	cobraCmdNames := CobraCommandToName(allBitCmds)
	log.Debug().Msg((time.Now().Sub(start)).String())
	start = time.Now()
	gitAddSuggestions := GitAddSuggestions()
	log.Debug().Msg((time.Now().Sub(start)).String())
	start = time.Now()
	//gitResetSuggestions := GitResetSuggestions()
	log.Debug().Msg((time.Now().Sub(start)).String())
	start = time.Now()
	gitmojiSuggestions := GitmojiSuggestions()
	log.Debug().Msg((time.Now().Sub(start)).String())

	branchListText := funk.Map(branchListSuggestions, func(s prompt.Suggest) string {
		return s.Text
	}).([]string)

	gitAddList := funk.Map(gitAddSuggestions, func(s prompt.Suggest) string {
		return s.Text
	}).([]string)

	gitmojiList := funk.Map(gitmojiSuggestions, func(s prompt.Suggest) string {
		return s.Text
	}).([]string)

	suggestionTree := b
	// add dynamic predictions and bit specific commands
	b.Sub["add"].Args = predict.Set(gitAddList)
	b.Sub["checkout"].Args = predict.Set(branchListText)
	b.Sub["co"].Args = predict.Set(branchListText)
	b.Sub["log"].Args = predict.Set(branchListText)
	b.Sub["merge"].Args = predict.Set(branchListText)
	b.Sub["rebase"].Args = predict.Set(branchListText)
	b.Sub["release"] = &complete.Command{Description: "Commit unstaged changes, bump minor tag, push", Args: predict.Set{"bump", "<version>"}}
	b.Sub["pr"] = &complete.Command{Description: "Check out a pull request from Github (requires GH CLI)", Args: complete.PredictFunc(lazyLoad(GitHubPRSuggestions))}
	b.Sub["gitmoji"] = &complete.Command{Description: "(Pre-alpha) Commit using gitmojis", Args: predict.Set(gitmojiList)}
	b.Sub["save"] = &complete.Command{Description: "Save your changes to your current branch"}
	b.Sub["update"] = &complete.Command{Description: "Updates bit to the latest or specified version"}
	b.Sub["complete"] = &complete.Command{Description: "Add classical tab completion to bit"}
	b.Sub["sync"] = &complete.Command{Description: "Synchronizes local changes with changes on origin or specified branch"}
	b.Sub["reset"].Args = predict.Set{"HEAD~1"}
	b.Sub["status"] = &complete.Command{
		Description: "Show the working tree status",
		Flags: map[string]complete.Predictor{
			"porcelain": predict.Set{"v1", "v2"},
		},
	}
	b.Sub["submodule"] = &complete.Command{Description: "Initialize, update or inspect submodules"}
	b.Sub["switch"] = &complete.Command{Description: "Switch branches", Args: predict.Set(branchListText)}

	// dynamically add "Common Commands" & "Git aliases"
	for _, name := range cobraCmdNames {
		if suggestionTree.Sub[name] != nil {
			continue
		}
		suggestionTree.Sub[name] = &complete.Command{}
	}

	funk.ForEach(branchListSuggestions, func(s prompt.Suggest) {
		if descriptionMap[s.Text] != "" {
			return
		}
		descriptionMap[s.Text] = s.Description
	})

	funk.ForEach(gitmojiSuggestions, func(s prompt.Suggest) {
		if descriptionMap[s.Text] != "" {
			return
		}
		descriptionMap[s.Text] = s.Description
	})

	// command
	// flags
	// commands
	// value

	//completerSuggestionMap := map[string]func() []prompt.Suggest{
	//	"":         memoize([]prompt.Suggest{}),
	//	"shell":    memoize(combraCommandSuggestions),
	//	"checkout": memoize(branchListSuggestions),
	//	"switch":   memoize(branchListSuggestions),
	//	"co":       memoize(branchListSuggestions),
	//	"merge":    memoize(branchListSuggestions),
	//	"rebase":   memoize(branchListSuggestions),
	//	"log":      memoize(branchListSuggestions),
	//	"add":      memoize(gitAddSuggestions),
	//	"release": memoize([]prompt.Suggest{
	//		{Text: "bump", Description: "Increment SemVer from tags and release e.g. if latest is v0.1.2 it's bumped to v0.1.3 "},
	//		{Text: "<version>", Description: "Name of release version e.g. v0.1.2"},
	//	}),
	//	"reset":   memoize(gitResetSuggestions),
	//	"pr":      lazyLoad(GitHubPRSuggestions),
	//	"gitmoji": memoize(gitmoji),
	//	"save":    memoize(gitmoji),
	//	//"_any": commonCommands,
	//}
	return suggestionTree, bitCmdMap
}

var pushFlagsMap = map[string]complete.Predictor{
	"all":                 predict.Nothing,
	"atomic":              predict.Nothing,
	"delete":              predict.Nothing,
	"d":                   predict.Nothing,
	"dry-run":             predict.Nothing,
	"n":                   predict.Nothing,
	"follow-tags":         predict.Nothing,
	"force":               predict.Nothing,
	"f":                   predict.Nothing,
	"force-with-lease":    predict.Nothing,
	"ipv4":                predict.Nothing,
	"4":                   predict.Nothing,
	"ipv6":                predict.Nothing,
	"6":                   predict.Nothing,
	"mirror":              predict.Nothing,
	"no-force-with-lease": predict.Nothing,
	"no-signed":           predict.Nothing,
	"no-thin":             predict.Nothing,
	"no-verify":           predict.Nothing,
	"porcelain":           predict.Nothing,
	"progress":            predict.Nothing,
	"prune":               predict.Nothing,
	"push-option":         predict.Nothing,
	"o":                   predict.Nothing,
	"quiet":               predict.Nothing,
	"q":                   predict.Nothing,
	"receive-pack":        predict.Nothing,
	"--exec":              predict.Nothing,
	"recurse-submodules":  predict.Nothing,
	"repo":                predict.Nothing,
	"set-upstream":        predict.Nothing,
	"u":                   predict.Nothing,
	"signed":              predict.Nothing,
	"sign":                predict.Nothing,
	"tags":                predict.Nothing,
	"thin":                predict.Nothing,
	"verbose":             predict.Nothing,
	"v":                   predict.Nothing,
}
