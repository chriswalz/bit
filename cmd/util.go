package cmd

import (
	"fmt"
	"github.com/AlecAivazis/survey/v2"
	"github.com/c-bata/go-prompt"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"github.com/thoas/go-funk"
	"os"
	"os/exec"
	"regexp"
	"runtime"
	"runtime/debug"
	"strings"
	"sync"
)

func RunInTerminalWithColor(cmdName string, args []string) error {
	dir, err := os.Getwd()
	if err != nil {
		return err
	}
	return RunInTerminalWithColorInDir(cmdName, dir, args)
}

func RunInTerminalWithColorInDir(cmdName string, dir string, args []string) error {
	log.Debug().Msg(cmdName + " " + strings.Join(args, " "))

	_, w, err := os.Pipe()
	if err != nil {
		panic(err)
	}

	cmd := exec.Command(cmdName, args...)
	cmd.Dir = dir
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if runtime.GOOS != "windows" {
		cmd.ExtraFiles = []*os.File{w}
	}

	err = cmd.Run()
	log.Debug().Err(err)
	return err
}

func AskConfirm(q string) bool {
	ans := false
	survey.AskOne(&survey.Confirm{
		Message: q,
	}, &ans)
	return ans
}

func AskMultiLine(q string) string {
	text := ""
	prompt := &survey.Multiline{
		Message: q,
	}
	survey.AskOne(prompt, &text)
	return text
}

func BranchList() []Branch {
	rawBranchData, err := branchListRaw()
	if err != nil {
		log.Debug().Err(err)
	}
	return toStructuredBranchList(rawBranchData)
}

func toStructuredBranchList(rawBranchData string) []Branch {

	list := strings.Split(strings.TrimSpace(rawBranchData), "\n")

	var branches []Branch
	for _, line := range list {
		// first character of each should start with ' which all commits have based on expected raw formatting
		if !strings.HasPrefix(line, `'`) {
			continue
		}

		cols := strings.Split(line[1:], "; ")
		b := Branch{
			Author:       cols[1],
			Name:         cols[3],
			RelativeDate: cols[4],
			AbsoluteDate: cols[0],
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
		log.Debug().Err(err)
	}
	out := string(msg)
	return strings.TrimSpace(out)
}

func AddCommandToShellHistory(cmd string, args []string) {
	// not possible??
	msg, err := exec.Command("/bin/bash", "-c", "history").CombinedOutput()
	if err != nil {
		log.Debug().Err(err)
	}
	log.Debug().Msg(string(msg))
}

func BranchListSuggestions() []prompt.Suggest {
	branches := BranchList()
	var suggestions []prompt.Suggest
	for _, branch := range branches {
		suggestions = append(suggestions, prompt.Suggest{
			Text:        branch.Name,
			Description: fmt.Sprintf("%s  %s  %s", branch.Author, branch.RelativeDate, branch.AbsoluteDate),
		})
	}
	return suggestions
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

func GitHubPRSuggestions() []prompt.Suggest {
	log.Debug().Msg("Github suggestions retrieving")
	prs := ListGHPullRequests()
	suggestions := funk.Map(prs, func(pr *PullRequest) prompt.Suggest {
		return prompt.Suggest{
			Text:        fmt.Sprintf("%s:%s-#%d", pr.State, pr.AuthorBranch, pr.Number),
			Description: fmt.Sprintf("%s", pr.Title),
		}
	})
	return suggestions.([]prompt.Suggest)
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
	AbsoluteDate string
}

type FileChange struct {
	Name   string
	Status string
}

func SuggestionPrompt(prefix string, completer func(d prompt.Document) []prompt.Suggest) string {
	result := prompt.Input(prefix, completer,
		prompt.OptionTitle(""),
		prompt.OptionHistory([]string{""}),
		prompt.OptionPrefixTextColor(prompt.Yellow), // fine
		//prompt.OptionPreviewSuggestionBGColor(prompt.Yellow),
		//prompt.OptionPreviewSuggestionTextColor(prompt.Yellow),
		prompt.OptionSelectedSuggestionBGColor(prompt.Yellow),
		prompt.OptionSuggestionBGColor(prompt.Yellow),
		prompt.OptionSuggestionTextColor(prompt.DarkGray),
		prompt.OptionSelectedSuggestionTextColor(prompt.Blue),
		prompt.OptionDescriptionBGColor(prompt.Black),
		prompt.OptionDescriptionTextColor(prompt.White),
		prompt.OptionShowCompletionAtStart(),
		prompt.OptionCompletionOnDown(),
		prompt.OptionSwitchKeyBindMode(prompt.EmacsKeyBind),
		prompt.OptionAddKeyBind(prompt.KeyBind{
			Key: prompt.ControlC,
			Fn:  exit,
		}),
		prompt.OptionAddASCIICodeBind(prompt.ASCIICodeBind{
			ASCIICode: []byte{0x1b, 0x62},
			Fn:        prompt.GoLeftWord,
		}),
		prompt.OptionAddASCIICodeBind(prompt.ASCIICodeBind{
			ASCIICode: []byte{0x1b, 0x66},
			Fn:        prompt.GoRightWord,
		}),
	)
	branchName := strings.TrimSpace(result)
	if strings.HasPrefix(result, "origin/") {
		branchName = branchName[7:]
	}
	return branchName
}

type Exit int

func exit(_ *prompt.Buffer) {
	panic(Exit(0))
}

func HandleExit() {
	switch v := recover().(type) {
	case nil:
		return
	case Exit:
		os.Exit(int(v))
	default:
		fmt.Println(v)
		fmt.Println(string(debug.Stack()))
		fmt.Println("OS:", runtime.GOOS, runtime.GOARCH)
		fmt.Println("bit version " + GetVersion())
		PrintGitVersion()

	}
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

	flagMap := map[string]string{
		"add":      addFlagsStr,
		"diff":     diffFlagsStr,
		"status":   statusFlagsStr,
		"commit":   commitFlagsStr,
		"branch":   branchFlagsStr,
		"tag":      tagFlagsStr,
		"checkout": checkoutFlagsStr,
		"merge":    mergeFlagsStr,
		"pull":     pullFlagsStr,
		"push":     pushFlagsStr,
		"log":      logFlagsStr,
		"rebase":   rebaseFlagsStr,
		"reset":    resetFlagsStr,
	}
	str = flagMap[gitSubCmd]
	if str == "" {
		return []prompt.Suggest{}
		//str = parseManPage(gitSubCmd)
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
	}
	return suggestions
}

func CommonCommandsList() []*cobra.Command {
	return []*cobra.Command{
		{
			Use:   "status",
			Short: "Show the working tree status",
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
		{
			Use:   "merge --squash",
			Short: "Squash and merge changes from a specified branch",
		},
		{
			Use:   "release bump",
			Short: "Commit unstaged changes, bump minor tag, push",
		},
		{
			Use:   "log --oneline",
			Short: "Display one commit per line",
		},
		{
			Use:   "diff --cached",
			Short: "Shows all staged changes",
		},
	}
}

func RunScriptWithString(path string, script string, args ...string) {
	var err error
	err = RunInTerminalWithColor("bin/sh", args)
	if err != nil {
		log.Debug().Err(err)
	}
}

func parseManPage(subCmd string) string {
	msg, err := exec.Command("/bin/bash", "-c", "git help "+subCmd+" | col -b").CombinedOutput()
	if err != nil {
		log.Debug().Err(err)
	}
	splitA := strings.Split(string(msg), "\n\nOPTIONS")
	splitB := regexp.MustCompile(`\.\n\n[A-Z]+`).Split(splitA[1], 2)
	rawFlagSection := splitB[0]
	//removeTabs := strings.ReplaceAll(rawFlagSection, "\t", "%%%")
	//removeTabs := strings.ReplaceAll(rawFlagSection, "\n\t", "")
	//removeTabs = strings.ReplaceAll(removeTabs, "%%%", "\n\n\t")
	return rawFlagSection
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func isBranchCompletionCommand(command string) bool {
	return command == "checkout" || command == "switch" || command == "co" || command == "pr" || command == "merge"
}

func isBranchChangeCommand(command string) bool {
	return command == "checkout" || command == "switch" || command == "co" || command == "pr"
}

func Find(slice []string, val string) int {
	for i, item := range slice {
		if item == val {
			return i
		}
	}
	return -1
}

func parseCommandLine(command string) ([]string, error) {
	var args []string
	state := "start"
	current := ""
	quote := "\""
	escapeNext := true
	for i := 0; i < len(command); i++ {
		c := command[i]

		if state == "quotes" {
			if string(c) != quote {
				current += string(c)
			} else {
				args = append(args, current)
				current = ""
				state = "start"
			}
			continue
		}

		if escapeNext {
			current += string(c)
			escapeNext = false
			continue
		}

		if c == '\\' {
			escapeNext = true
			continue
		}

		if c == '"' || c == '\'' {
			state = "quotes"
			quote = string(c)
			continue
		}

		if state == "arg" {
			if c == ' ' || c == '\t' {
				args = append(args, current)
				current = ""
				state = "start"
			} else {
				current += string(c)
			}
			continue
		}

		if c != ' ' && c != '\t' {
			state = "arg"
			current += string(c)
		}
	}

	if state == "quotes" {
		return []string{}, fmt.Errorf("Unclosed quote in command line: %s", command)
	}

	if current != "" {
		args = append(args, current)
	}

	return args, nil
}

func memoize(suggestions []prompt.Suggest) func() []prompt.Suggest {
	return func () []prompt.Suggest {
		return suggestions
	}
}

func lazyLoad(suggestionFunc func() []prompt.Suggest) func() []prompt.Suggest {
	var suggestions []prompt.Suggest
	return func () []prompt.Suggest {
		if suggestions == nil {
			suggestions = suggestionFunc()
		}
		return suggestions
	}
}

func asyncLoad(suggestionFunc func() []prompt.Suggest) func() []prompt.Suggest {
	var suggestions []prompt.Suggest
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		suggestions = suggestionFunc()
	}()
	return func () []prompt.Suggest {
		wg.Wait()
		return suggestions
	}
}