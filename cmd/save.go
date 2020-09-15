package cmd

import (
	"fmt"
	"github.com/chriswalz/bit/util"
	"github.com/spf13/cobra"
	"strings"
)

// saveCmd represents the save command
var saveCmd = &cobra.Command{
	Use:   "save",
	Short: "Save your changes to your current branch",
	Long: `E.g. bit save; bit save "commit message"`,
	Run: func(cmd *cobra.Command, args []string) {
		msg := ""
		if len(args) > 0 {
			msg = strings.Join(args, " ")
		}
		save(msg)
	},
	//Args: cobra.MaximumNArgs(1),
}
// add comment

func init() {
	rootCmd.AddCommand(saveCmd)
	// saveCmd.PersistentFlags().String("foo", "", "A help for foo")
	// saveCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func save(msg string) {
	if util.NothingToCommit() {
		fmt.Println("nothing to save or commit")
		return
	}

	if msg == "" {
		//if ahead of master
		if util.IsAheadOfCurrent() || !util.CloudBranchExists(){
			util.Runwithcolor([]string{"commit", "-a", "--amend", "--no-edit"}) // amend if already ahead
		} else {
			util.Runwithcolor([]string{"status", "-sb"})
			resp := util.PromptUser("Please provide a description of your commit (what you're saving)")
			util.Runwithcolor([]string{"commit", "-a", "-m " + resp})
		}
	} else {
		util.Runwithcolor([]string{"commit", "-a", "-m " + msg})
	}
}