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
	Short: "Save your changes to your current branch",
	Long: `E.g. bit save; bit save "commit message"`,
	Run: func(cmd *cobra.Command, args []string) {
		if nothingToCommit() {
			fmt.Println("nothing to save or commit")
			return
		}
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


	util.Runwithcolor([]string{"add", "."})
	if msg == "" {
		// if ahead of master
		if isAheadOfCurrent() || !cloudBranchExists(){
			util.Runwithcolor([]string{"commit", "--amend", "--no-edit"}) // amend if already ahead
		} else {
			util.Runwithcolor([]string{"status", "-sb"})
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

func isBehindCurrent() bool {
	msg, err := exec.Command("git", "status", "-sb").CombinedOutput()
	if err != nil {
		//fmt.Println(err)
	}
	return strings.Contains(string(msg), "behind")
}

func nothingToCommit() bool {
	msg, err := exec.Command("git", "status").CombinedOutput()
	if err != nil {
		//fmt.Println(err)
	}
	return strings.Contains(string(msg), "nothing to commit")
}