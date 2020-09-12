package cmd

import (
	"bufio"
	"fmt"
	"github.com/chriswalz/bit/util"
	"github.com/spf13/cobra"
	"os"
	"os/exec"
	"strings"
)

// saveCmd represents the save command
var saveCmd = &cobra.Command{
	Use:   "save",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
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
	// if nothing to commit
	// do nothing asd

	util.Runwithcolor([]string{"add", "."})
	if msg == "" {
		// if ahead of master
		if isAheadOfCurrent() || !cloudBranchExists(){
			util.Runwithcolor([]string{"commit", "--amend", "--no-edit"}) // amend if already ahead
		} else {
			resp := promptUser("Please provide a description of your commit (what you're saving)")
			util.Runwithcolor([]string{"commit", "-m " + resp})
		}
	} else {
		util.Runwithcolor([]string{"commit", "-m " + msg})
	}
}

func promptUser(prompt string) string {
	fmt.Println(prompt)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	resp := scanner.Text()
	return resp
}

func isAheadOfCurrent() bool {
	msg, err := exec.Command("git", "status", "-sb").CombinedOutput()
	if err != nil {
		//fmt.Println(err)
	}
	return strings.Contains(string(msg), "ahead")
}