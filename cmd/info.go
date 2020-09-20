package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"os/exec"
)

// infoCmd represents the info command
var infoCmd = &cobra.Command{
	Use:   "info",
	Short: "Get general information about the status of your repository",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("--- INFO ---")
		out, err := exec.Command("/bin/sh", "git-extras/git-info.sh").Output()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(out))

		fmt.Println("--- SUMMARY ---")
		out, err = exec.Command("/bin/sh", "git-extras/git-summary.sh").Output()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(out))

		fmt.Println("--- EFFORT ---")
		out, err = exec.Command("/bin/sh", "git-extras/git-effort.sh").Output()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(out))
	},
	//Args: cobra.MaximumNArgs(1),
}

func init() {
	rootCmd.AddCommand(infoCmd)
}
