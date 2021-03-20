package cmd

import (
	"errors"
	"fmt"
	"github.com/chriswalz/complete/v3"
	exec "golang.org/x/sys/execabs"
	"os"
	"regexp"
	"runtime"
	"runtime/debug"
	"strconv"
	"strings"
	"sync"

	"github.com/AlecAivazis/survey/v2"
	"github.com/c-bata/go-prompt"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"github.com/thoas/go-funk"
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
	log.Debug().Err(err).Send()
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

func BranchList() []*Branch {
	rawBranchData, err := branchListRaw()
	if err != nil {
		log.Debug().Err(err).Send()
	}
	return toStructuredBranchList(rawBranchData)
}

func toStructuredBranchList(rawBranchData string) []*Branch {
	list := strings.Split(strings.TrimSpace(rawBranchData), "\n")

	m := map[string]*Branch{}
	var branches []*Branch
	for _, line := range list {
		// first character of each should start with ' which all commits have based on expected raw formatting
		if !strings.HasPrefix(line, `'`) {
			continue
		}

		cols := strings.Split(line[1:], "; ")
		b := &Branch{
			Author:       cols[1],
			FullName:     cols[3],
			RelativeDate: cols[4],
			AbsoluteDate: cols[0],
		}
		if b.FullName == "origin/master" || b.FullName == "origin/HEAD" {
			continue
		}
		if !strings.HasPrefix(b.FullName, "origin/") {
			m[b.FullName] = b
		}

		branches = append(branches, b)
	}
	return funk.Filter(branches, func(b *Branch) bool {
		if strings.HasPrefix(b.FullName, "origin/") && m[b.FullName[7:]] != nil {
			return false
		}
		return true
	}).([]*Branch)
}

func GenBumpedSemVersion(out string) (string, error) {
	out = strings.TrimSpace(out)

	// 0.0.1
	if len(out) <= 0 {
		return "", errors.New("no tags exists. Consider running `bit release 0.0.1`")
	}
	i := strings.LastIndex(out, ".")
	minor, err := strconv.Atoi(out[i+1:])
	if err != nil {
		return "", err
	}
	minor++
	out = out[:i] + "." + strconv.Itoa(minor)
	return out, nil
}

func AddCommandToShellHistory(cmd string, args []string) {
	// not possible??
	msg, err := exec.Command("/bin/bash", "-c", "history").CombinedOutput()
	if err != nil {
		log.Debug().Err(err).Send()
	}
	log.Debug().Msg(string(msg))
}

func BranchListSuggestions() []complete.Suggestion {
	branches := BranchList()
	var suggestions []complete.Suggestion
	for _, branch := range branches {
		suggestions = append(suggestions, complete.Suggestion{
			Name: branch.FullName,
			Desc: fmt.Sprintf("%s  %s  %s", branch.Author, branch.RelativeDate, branch.AbsoluteDate),
		})
	}
	return suggestions
}

func GitAddSuggestions() []complete.Suggestion {
	fileChanges := FileChangesList()
	var suggestions []complete.Suggestion
	suggestions = append(suggestions, complete.Suggestion{
		Name: "-u",
		Desc: "Add modified and deleted files and exclude untracked files.",
	})
	for _, fc := range fileChanges {
		suggestions = append(suggestions, complete.Suggestion{
			Name: fc.Name,
			Desc: fc.Status + " ~~~",
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

func GitHubPRSuggestions(prefix string) func(prefix string) []complete.Suggestion {
	return func(prefix string) []complete.Suggestion {
		log.Debug().Msg("Github suggestions retrieving")
		prs := ListGHPullRequests()
		suggestions := funk.Map(prs, func(pr *PullRequest) complete.Suggestion {
			return complete.Suggestion{
				Name: fmt.Sprintf("%s:%s-#%d", pr.State, pr.AuthorBranch, pr.Number),
				Desc: pr.Title,
			}
		})
		return suggestions.([]complete.Suggestion)
	}
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

func CobraCommandToName(cmds []*cobra.Command) []string {
	var ss []string
	for _, cmd := range cmds {
		ss = append(ss, cmd.Use)
	}
	return ss
}

func CobraCommandToDesc(cmds []*cobra.Command) []string {
	var ss []string
	for _, cmd := range cmds {
		ss = append(ss, cmd.Short)
	}
	return ss
}

type Branch struct {
	Author       string
	FullName     string
	RelativeDate string
	AbsoluteDate string
}

type FileChange struct {
	Name   string
	Status string
}

type PromptTheme struct {
	PrefixTextColor             prompt.Color
	SelectedSuggestionBGColor   prompt.Color
	SuggestionBGColor           prompt.Color
	SuggestionTextColor         prompt.Color
	SelectedSuggestionTextColor prompt.Color
	DescriptionBGColor          prompt.Color
	DescriptionTextColor        prompt.Color
}

var DefaultTheme = PromptTheme{
	PrefixTextColor:             prompt.Yellow,
	SelectedSuggestionBGColor:   prompt.Yellow,
	SuggestionBGColor:           prompt.Yellow,
	SuggestionTextColor:         prompt.DarkGray,
	SelectedSuggestionTextColor: prompt.Blue,
	DescriptionBGColor:          prompt.Black,
	DescriptionTextColor:        prompt.White,
}

var InvertedTheme = PromptTheme{
	PrefixTextColor:             prompt.Blue,
	SelectedSuggestionBGColor:   prompt.LightGray,
	SelectedSuggestionTextColor: prompt.White,
	SuggestionBGColor:           prompt.Blue,
	SuggestionTextColor:         prompt.White,
	DescriptionBGColor:          prompt.LightGray,
	DescriptionTextColor:        prompt.Black,
}

var MonochromeTheme = PromptTheme{}

func SuggestionPrompt(prefix string, completer func(d prompt.Document) []prompt.Suggest) string {
	theme := DefaultTheme
	themeName := os.Getenv("BIT_THEME")
	if strings.EqualFold(themeName, "inverted") {
		theme = InvertedTheme
	}
	if strings.EqualFold(themeName, "monochrome") {
		theme = MonochromeTheme
	}
	result := prompt.Input(prefix, completer,
		prompt.OptionTitle(""),
		prompt.OptionHistory([]string{""}),
		prompt.OptionPrefixTextColor(theme.PrefixTextColor), // fine
		prompt.OptionSelectedSuggestionBGColor(theme.SelectedSuggestionBGColor),
		prompt.OptionSuggestionBGColor(theme.SuggestionBGColor),
		prompt.OptionSuggestionTextColor(theme.SuggestionTextColor),
		prompt.OptionSelectedSuggestionTextColor(theme.SelectedSuggestionTextColor),
		prompt.OptionDescriptionBGColor(theme.DescriptionBGColor),
		prompt.OptionDescriptionTextColor(theme.DescriptionTextColor),
		// prompt.OptionPreviewSuggestionBGColor(prompt.Yellow),
		// prompt.OptionPreviewSuggestionTextColor(prompt.Yellow),
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
	//gitCmds := AllGitSubCommands()
	//bitCmds, _ := AllBitSubCommands(rootCmd)
	commonCommands := CommonCommandsList()
	return concatCopyPreAllocate([][]*cobra.Command{commonCommands, gitAliases})
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
		// str = parseManPage(gitSubCmd)
	}

	list := strings.Split(str, ".\n\n")

	// list := strings.Split(strings.Split(op[1], "CONFIGURATION")[0], "\n\n")
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
		log.Debug().Err(err).Send()
	}
}

func parseManPage(subCmd string) string {
	msg, err := exec.Command("/bin/bash", "-c", "git help "+subCmd+" | col -b").CombinedOutput()
	if err != nil {
		log.Debug().Err(err).Send()
	}
	splitA := strings.Split(string(msg), "\n\nOPTIONS")
	splitB := regexp.MustCompile(`\.\n\n[A-Z]+`).Split(splitA[1], 2)
	rawFlagSection := splitB[0]
	// removeTabs := strings.ReplaceAll(rawFlagSection, "\t", "%%%")
	// removeTabs := strings.ReplaceAll(rawFlagSection, "\n\t", "")
	// removeTabs = strings.ReplaceAll(removeTabs, "%%%", "\n\n\t")
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
	return command == "checkout" || command == "switch" || command == "co" || command == "pr" || command == "merge" || command == "rebase"
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
	return func() []prompt.Suggest {
		return suggestions
	}
}

func lazyLoad(predictFunc func(prefix string) []complete.Suggestion) func(prefix string) []complete.Suggestion {
	var result []complete.Suggestion
	return func(prefix string) []complete.Suggestion {
		if result == nil {
			result = predictFunc(prefix)
		}
		return result
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
	return func() []prompt.Suggest {
		wg.Wait()
		return suggestions
	}
}
