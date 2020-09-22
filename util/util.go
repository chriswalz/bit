package util

import (
	"bufio"
	"fmt"
	"github.com/c-bata/go-prompt"
	"github.com/spf13/cobra"
	"log"
	"os"
	"os/exec"
	"strings"
)

func Runwithcolor(args []string) error {
	_, w, err := os.Pipe()
	if err != nil {
		panic(err)
	}

	cmd := exec.Command("git", args...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.ExtraFiles = []*os.File{w}

	err = cmd.Run()
	//fmt.Println(err)
	return err
}


func CloudBranchExists() bool {
	msg, err := exec.Command("git", "pull").CombinedOutput()
	if err != nil {
		fmt.Println(err)
	}
	//log.Println("msg:", string(msg))
	//log.Println("err:", err)
	return !strings.Contains(string(msg), "There is no tracking information for the current branch")
}

func CurrentBranch() string {
	msg, err := exec.Command("git", "branch", "--show-current").CombinedOutput()
	if err != nil {
		fmt.Println(err)
	}
	return strings.TrimSpace(string(msg))
}

func IsYes(resp string) bool {
	return resp == "YES" || resp == "Y" || resp == "yes" || resp == "y"
}

func IsAheadOfCurrent() bool {
	msg, err := exec.Command("git", "status", "-sb").CombinedOutput()
	if err != nil {
		//fmt.Println(err)
	}
	return strings.Contains(string(msg), "ahead")
}

func IsBehindCurrent() bool {
	msg, err := exec.Command("git", "status", "-sb").CombinedOutput()
	if err != nil {
		//fmt.Println(err)
	}
	return strings.Contains(string(msg), "behind")
}

func NothingToCommit() bool {
	msg, err := exec.Command("git", "status").CombinedOutput()
	if err != nil {
		//fmt.Println(err)
	}
	return strings.Contains(string(msg), "nothing to commit")
}

func IsDiverged() bool {
	msg, err := exec.Command("git", "status").CombinedOutput()
	if err != nil {
		//fmt.Println(err)
	}
	return strings.Contains(string(msg), "have diverged")
}

// fixme when writing input the user cant backspace in the normal fashion
func PromptUser(prompt string) string {
	fmt.Println(prompt)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	resp := scanner.Text()
	return resp
}

func StashableChanges() bool {
	msg, err := exec.Command("git", "status").CombinedOutput()
	if err != nil {
		//fmt.Println(err)
	}
	return strings.Contains(string(msg), "Changes to be committed:") || strings.Contains(string(msg), "Changes not staged for commit:")
}

func StashList() []string {
	msg, err := exec.Command("git", "stash", "list").CombinedOutput()
	if err != nil {
		//fmt.Println(err)
	}
	return strings.Split(string(msg), "\n")
}

func BranchList() []Branch {
	msg, err := exec.Command("git", "for-each-ref", "--sort=-committerdate", "refs/heads/", "refs/remotes", "--format='%(authordate:short); %(authorname); %(color:red)%(objectname:short); %(color:yellow)%(refname:short)%(color:reset); (%(color:green)%(committerdate:relative)%(color:reset))'").CombinedOutput()
	if err != nil {
		//fmt.Println(err)
	}
	list := strings.Split(string(msg), "\n")
	var branches []Branch
	for i := 0; i < len(list)-1; i++ {
		cols := strings.Split(list[i], "; ")
		b := Branch{
			Author:       cols[1],
			Name:         cols[3],
			RelativeDate: list[i][strings.Index(list[i], "("):],
		}
		branches = append(branches, b)
	}
	return branches
}

type Branch struct {
	Author string
	Name string
	RelativeDate string
}

func SuggestionPrompt(prefix string, completer func(d prompt.Document) []prompt.Suggest) string {
	result := prompt.Input(prefix, completer,
		prompt.OptionTitle(""),
		prompt.OptionHistory([]string{""}),
		prompt.OptionPrefixTextColor(prompt.Yellow),
		prompt.OptionPreviewSuggestionTextColor(prompt.Yellow),
		prompt.OptionSelectedSuggestionBGColor(prompt.Yellow),
		prompt.OptionSuggestionBGColor(prompt.Yellow),
		prompt.OptionSelectedSuggestionTextColor(prompt.Purple),
		prompt.OptionShowCompletionAtStart(),
		prompt.OptionCompletionOnDown(),
		prompt.OptionSwitchKeyBindMode(prompt.CommonKeyBind),
		prompt.OptionAddKeyBind(prompt.KeyBind{
			Key: prompt.ControlC,
			Fn: func(b *prompt.Buffer) {
				err := os.Stdin.Close()
				log.Println(err)
				err = os.Stdin.Close()
				log.Println(err)
				os.Exit(1)
			},
		}),
	)
	branchName := strings.TrimSpace(result)
	if strings.HasPrefix(result, "origin/") {
		branchName = branchName[7:]
	}
	return branchName
}

func AllGitSubCommands() (cc []*cobra.Command) {
	msg, err := exec.Command("git", "help", "-a").CombinedOutput()
	if err != nil {
		fmt.Println(err)
	}
	commands := strings.Split(strings.Split(strings.Split(string(msg), "Main Porcelain Commands")[1], "Ancillary Commands")[0], "\n")
	for _, command := range commands {
		if command == "" {
			continue
		}
		split := strings.Split(strings.TrimSpace(command), "   ")
		c := cobra.Command{
			Use: split[0],
			Short: split[len(split) - 1],
		}
		cc = append(cc, &c)
	}

	return cc
}