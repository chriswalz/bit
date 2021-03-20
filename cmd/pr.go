package cmd

import (
	"github.com/chriswalz/complete/v3"
	"strconv"
	"strings"

	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

// prCmd represents the pr command
var prCmd = &cobra.Command{
	Use:   "pr",
	Short: "Check out a pull request from Github (requires GH CLI)",
	Long: `bit pr
bit pr`,
	Run: func(cmd *cobra.Command, args []string) {
		suggestionTree := &complete.CompTree{
			Sub: map[string]*complete.CompTree{
				"pr": {
					Desc:    "Check out a pull request from Github (requires GH CLI)",
					Dynamic: lazyLoad(GitHubPRSuggestions("")),
				},
			},
		}
		runPr(suggestionTree)
	},
	Args: cobra.NoArgs,
}

func init() {
	BitCmd.AddCommand(prCmd)
}

func runPr(suggestionMap *complete.CompTree) {
	branchName := SuggestionPrompt("> bit pr ", specificCommandCompleter("pr", suggestionMap))

	split := strings.Split(branchName, "#")
	prNumber, err := strconv.Atoi(split[len(split)-1])
	if err != nil {
		log.Debug().Err(err).Send()
		return
	}
	checkoutPullRequest(prNumber)
	return
}
