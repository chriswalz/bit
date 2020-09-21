package cmd

import (
	"fmt"
	"github.com/c-bata/go-prompt"
	"github.com/chriswalz/bit/util"
	"github.com/spf13/cobra"
	"os"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "bit",
	Short: "Bit is Git with a simple interface. Plus you can still use all the old git commands",
	Long: `v0.3.3`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		cmdMap := map[string]*cobra.Command{}
		for _, c := range cmd.Commands() {
			cmdMap[c.Name()] = c
		}
		resp := util.SuggestionPrompt("bit ", rootCommandCompleter(cmd.Commands()))
		cmd.SetArgs([]string{resp})
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

func rootCommandCompleter(cmds []*cobra.Command) func(d prompt.Document) []prompt.Suggest {
	return func(d prompt.Document) []prompt.Suggest {
		var suggestions []prompt.Suggest
		for _, branch := range cmds {
			suggestions = append(suggestions, prompt.Suggest{
				Text:        branch.Name(),
				Description: branch.Short,
			})
		}

		return prompt.FilterHasPrefix(suggestions, d.GetWordBeforeCursor(), true)
	}
}