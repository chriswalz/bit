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

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "bit version [branch-name]",
	Long: `For existing branches simply run bit version [branch-name]. 

For creating a new branch it's the same command! You'll simply be prompted for the name for the new branch
`,
	Run: func(cmd *cobra.Command, args []string) {
		util.Runwithcolor([]string{"fetch"})
		if !localOrRemoteBranchExists(args[0]) {
			//fmt.Println(strings.Contains(err.Error(), "did not"))
			fmt.Println("Branch does not exist. Do you want to create it? Y/n")
			var resp string
			fmt.Scanln(&resp)
			if resp == "YES" || resp == "Y" || resp == "yes" || resp == "y" {
				util.Runwithcolor([]string{"checkout", "-b", args[0]})
			}
		}
	},
	Args: cobra.ExactArgs(1),
}

func init() {
	rootCmd.AddCommand(versionCmd)
	// versionCmd.PersistentFlags().String("foo", "", "A help for foo")
	// versionCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func localOrRemoteBranchExists(branch string) bool {
	msg, err := exec.Command("git", "checkout", branch).CombinedOutput()
	if err != nil {
		//fmt.Println(err)
	}
	return !strings.Contains(string(msg), "did not match any file")
}