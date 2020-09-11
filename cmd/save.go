/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"github.com/chriswalz/bit/util"
	"github.com/spf13/cobra"
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
	util.Runwithcolor([]string{"add", "."})
	if msg == "" {
		// if ahead of master
		if isAheadOfCurrent() || !cloudBranchExists(){
			util.Runwithcolor([]string{"commit", "--amend", "--no-edit"}) // amend if already ahead
		} else {
			fmt.Println("Please provide a description of your commit (what you're saving)")
			var resp string
			fmt.Scanln(&resp)
			util.Runwithcolor([]string{"commit", "-m " + resp})
		}
	} else {
		util.Runwithcolor([]string{"commit", "-m " + msg})
	}
}

func isAheadOfCurrent() bool {
	msg, err := exec.Command("git", "status", "-sb").CombinedOutput()
	if err != nil {
		//fmt.Println(err)
	}
	return strings.Contains(string(msg), "ahead")
}