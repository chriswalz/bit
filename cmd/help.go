package cmd

import (
	"github.com/chriswalz/bit/util"
	"github.com/spf13/cobra"
)

// helpCmd represents the help command
var helpCmd = &cobra.Command{
	Use:   "help",
	Short: "bit help",
	Long: `For existing branches simply run bit help [branch-name]. 

For creating a new branch it's the same command! You'll simply be prompted to confirm that you want to create a new branch
`,
	Run: func(cmd *cobra.Command, args []string) {
		util.Runwithcolor([]string{"fetch"})
	},
	Args: cobra.ExactArgs(1),
}

func init() {
	rootCmd.AddCommand(helpCmd)
	// helpCmd.PersistentFlags().String("foo", "", "A help for foo")
	// helpCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
