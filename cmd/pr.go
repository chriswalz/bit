package cmd

import (
	"strconv"
	"strings"

	"github.com/c-bata/go-prompt"
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
		suggestionMap := map[string]func() []prompt.Suggest{
			"pr": lazyLoad(GitHubPRSuggestions),
		}
		runPr(suggestionMap)
	},
	Args: cobra.NoArgs,
}

func init() {
	ShellCmd.AddCommand(prCmd)
}

func runPr(suggestionMap map[string]func() []prompt.Suggest) {
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
