package cmd

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

func Runwithcolor(cmdName string, args []string) error {
	_, w, err := os.Pipe()
	if err != nil {
		panic(err)
	}

	cmd := exec.Command(cmdName, args...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.ExtraFiles = []*os.File{w}

	err = cmd.Run()
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

func isGitRepo() bool {
	_, err := exec.Command("git", "status").CombinedOutput()
	if err != nil {
		return false
	}
	return true
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
		if b.Name == "origin/master" || b.Name == "origin/HEAD" {
			continue
		}
		branches = append(branches, b)
	}
	return branches
}

func GenBumpedSemVersion() string {
	msg, err := exec.Command("/bin/sh", "-c", `git describe --tags --abbrev=0 | awk -F. '{$NF+=1; OFS="."; print $0}'`).CombinedOutput()
	if err != nil {
		fmt.Println(err)
	}
	out := string(msg)
	return out
}

func AddCommandToShellHistory(cmd string, args []string) {
	// not possible??
	msg, err := exec.Command("/bin/bash", "-c", "history").CombinedOutput()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(msg)
}

func BranchListSuggestions() []prompt.Suggest {
	branches := BranchList()
	var suggestions []prompt.Suggest
	for _, branch := range branches {
		suggestions = append(suggestions, prompt.Suggest{
			Text:        branch.Name,
			Description: branch.RelativeDate + " " + branch.Author,
		})
	}
	return suggestions
}

func FileChangesList() []FileChange {
	msg, err := exec.Command("git", "status", "--porcelain=v2").CombinedOutput()
	if err != nil {
		//fmt.Println(err)
	}
	list := strings.Split(string(msg), "\n")
	statusMap := map[string]string{
		"M.": "Added",
		"MM": "Partially Added",
		"A.": "New File",
		".M": "Not Staged",
		"?":  "Untracked",
	}
	var changes []FileChange
	for i := 0; i < len(list)-1; i++ {
		cols := strings.Fields(strings.TrimSpace(list[i]))
		b := FileChange{
			Name:   cols[len(cols)-1],
			Status: statusMap[cols[1]],
		}
		if len(cols) == 2 {
			b.Status = statusMap[cols[0]]
		}
		changes = append(changes, b)
	}
	return changes
}

func GitAddSuggestions() []prompt.Suggest {
	fileChanges := FileChangesList()
	var suggestions []prompt.Suggest
	suggestions = append(suggestions, prompt.Suggest{
		Text:        "-u",
		Description: "Add modified and deleted files and exclude untracked files.",
	})
	for _, fc := range fileChanges {
		suggestions = append(suggestions, prompt.Suggest{
			Text:        fc.Name,
			Description: fc.Status + " ~~~",
		})
	}
	return suggestions
}

func GitResetSuggestions() []prompt.Suggest {
	fileChanges := FileChangesList()
	var suggestions []prompt.Suggest
	for _, fc := range fileChanges {
		if fc.Status == "Added" || fc.Status == "Partially Added" || fc.Status == "New File" {
			suggestions = append(suggestions, prompt.Suggest{
				Text:        fc.Name,
				Description: fc.Status + " ~~~",
			})
		}
	}
	return suggestions
}

func CobraCommandToSuggestions(cmds []*cobra.Command) []prompt.Suggest {
	var suggestions []prompt.Suggest
	for _, branch := range cmds {
		suggestions = append(suggestions, prompt.Suggest{
			Text:        branch.Use,
			Description: branch.Short,
		})
	}
	return suggestions
}

type Branch struct {
	Author       string
	Name         string
	RelativeDate string
}

type FileChange struct {
	Name   string
	Status string
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
		prompt.OptionDescriptionBGColor(prompt.Black),
		prompt.OptionDescriptionTextColor(prompt.White),
		prompt.OptionShowCompletionAtStart(),
		prompt.OptionCompletionOnDown(),
		prompt.OptionSwitchKeyBindMode(prompt.CommonKeyBind),
		/*prompt.OptionAddKeyBind(prompt.KeyBind{
			Key: prompt.ControlC,
			Fn: func(b *prompt.Buffer) {
				os.Stdin.WriteString("exit\n")
				//err := os.Stdin.Close()
				//os.Exit(1)
			},
		}), */
	)
	branchName := strings.TrimSpace(result)
	if strings.HasPrefix(result, "origin/") {
		branchName = branchName[7:]
	}
	return branchName
}

func AllGitAliases() (cc []*cobra.Command) {
	msg, err := exec.Command("git", "config", "--get-regexp", "alias").CombinedOutput()
	if err != nil {
		fmt.Println(err)
	}
	aliases := strings.Split(string(msg), "\n")
	for _, alias := range aliases {
		if alias == "" {
			continue
		}
		split := strings.Fields(strings.TrimSpace(alias)[6:])
		c := cobra.Command{
			Use:   split[0],
			Short: strings.Join(split[1:], " "),
		}
		cc = append(cc, &c)
	}

	return cc
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
			Use:   split[0],
			Short: strings.TrimSpace(split[len(split)-1]),
		}
		cc = append(cc, &c)
	}

	return cc
}

func AllBitSubCommands(rootCmd *cobra.Command) ([]*cobra.Command, map[string]*cobra.Command) {
	bitCmds := rootCmd.Commands()
	bitCmdMap := map[string]*cobra.Command{}
	for _, bitCmd := range bitCmds {
		bitCmdMap[bitCmd.Name()] = bitCmd
	}
	return bitCmds, bitCmdMap
}

func AllBitAndGitSubCommands(rootCmd *cobra.Command) (cc []*cobra.Command) {
	gitAliases := AllGitAliases()
	gitCmds := AllGitSubCommands()
	bitCmds, _ := AllBitSubCommands(rootCmd)
	commonCommands := CommonCommandsList()
	return concatCopyPreAllocate([][]*cobra.Command{commonCommands, gitCmds, bitCmds, gitAliases})
}

func concatCopyPreAllocate(slices [][]*cobra.Command) []*cobra.Command {
	var totalLen int
	for _, s := range slices {
		totalLen += len(s)
	}
	tmp := make([]*cobra.Command, totalLen)
	var i int
	for _, s := range slices {
		i += copy(tmp[i:], s)
	}
	return tmp
}

func FlagSuggestionsForCommand(gitSubCmd string, flagtype string) []prompt.Suggest {
	str := ""

	// git help pull | col -b > man.txt
	//if gitSubCmd != "commit" && gitSubCmd != "push" && gitSubCmd != "status" {
	//	msg, err := exec.Command("/bin/sh", "-c", "git help " + gitSubCmd + " | col -bx").CombinedOutput()
	//	if err != nil {
	//		//fmt.Println(err)
	//	}
	//	out := string(msg)
	//	//out = stripCtlAndExtFromUTF8(out)
	//	split := strings.Split(out, "OPTIONS")
	//	fmt.Println(out, split)
	//	out=split[1]
	//	//return []prompt.Suggest{}
	//}
	flagMap := map[string]string{
		"add":      addFlagsStr,
		"diff":     diffFlagsStr,
		"status":   statusFlagsStr,
		"commit":   commitFlagsStr,
		"branch":   branchFlagsStr,
		"tags":     tagFlagsStr,
		"checkout": checkoutFlagsStr,
		"merge":    mergeFlagsStr,
		"pull":     pullFlagsStr,
		"push":     pushFlagsStr,
		"log":      logFlagsStr,
	}
	str = flagMap[gitSubCmd]
	if str == "" {
		return []prompt.Suggest{}
	}

	list := strings.Split(str, ".\n\n")

	//list := strings.Split(strings.Split(op[1], "CONFIGURATION")[0], "\n\n")
	var suggestions []prompt.Suggest
	for i := 0; i < len(list)-1; i++ {
		line := list[i]
		if !strings.HasPrefix(line, "-") && !strings.HasPrefix(line, "--") {
			continue
		}
		split := strings.Split(line, "\n")
		flags := split[:len(split)-1]
		desc := split[len(split)-1]
		for _, flag := range flags {
			if strings.HasPrefix(flag, "--") && flagtype == "--" {
				suggestions = append(suggestions, prompt.Suggest{
					Text:        flag,
					Description: desc,
				})
			}
			if !strings.HasPrefix(flag, "--") && flagtype == "-" {
				suggestions = append(suggestions, prompt.Suggest{
					Text:        flag,
					Description: desc,
				})
			}
		}

		//log.Println(list[i])
	}
	return suggestions
}

func CommonCommandsList() []*cobra.Command {
	return []*cobra.Command{
		{
			Use:   "pull --rebase",
			Short: "Rebase on origin branch",
		},
		{
			Use:   "pull --rebase origin master",
			Short: "Rebase on origin master branch",
		},
		{
			Use:   "push --force-with-lease",
			Short: "force push with a safety net",
		},
		{
			Use:   "stash pop",
			Short: "Use most recently stashed changes",
		},
		{
			Use:   "add -u",
			Short: "Add all modified or deleted files to Staging Area",
		},
		{
			Use:   "commit -am \"",
			Short: "Commit all tracked files",
		},
		{
			Use:   "commit -a --amend --no-edit",
			Short: "Amend most recent commit with new changes",
		},
		{
			Use:   "commit --amend --no-edit",
			Short: "Amend most recent commit with added changes",
		},
	}
}

func RunScriptWithString(path string, script string, args ...string) {
	f, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}
	f.WriteString(script)
	defer f.Close()
	err = Runwithcolor(path, args)
	if err != nil {
		fmt.Println(err)
	}
}
