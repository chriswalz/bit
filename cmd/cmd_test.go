package cmd

import (
	"github.com/c-bata/go-prompt"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"testing"
)

func toString(list []*cobra.Command) []string {
	newList := make([]string, len(list))
	for i, v := range list {
		newList[i] = v.Use
	}
	return newList
}

func suggestionToString(list []prompt.Suggest) []string {
	newList := make([]string, len(list))
	for i, v := range list {
		newList[i] = v.Text
	}
	return newList
}

func branchToString(list []Branch) []string {
	newList := make([]string, len(list))
	for i, v := range list {
		newList[i] = v.Name
	}
	return newList
}

func TestCommonCommandsList(t *testing.T) {
	expects := []string{"pull --rebase", "commit -a --amend --no-edit", "add -u"}
	reality := toString(CommonCommandsList())
	for _, e := range expects {
		assert.Contains(t, reality, e)
	}
}

func TestBranchList(t *testing.T) {
	expects := []string{"master"}
	notexpects := []string{"origin/master", "origin/HEAD"}
	reality := branchToString(BranchList())
	for _, e := range expects {
		assert.Contains(t, reality, e)
	}
	for _, ne := range notexpects {
		assert.NotContains(t, reality, ne)
	}
}

// Tests AllBitAndGitSubCommands has common commands, git sub commands, git aliases, git-extras and bit commands
func TestAllBitAndGitSubCommands(t *testing.T) {
	expects := []string{"pull --rebase", "commit -a --amend --no-edit", "add", "push", "fetch", "pull", "co", "lg", "release", "info", "save", "sync"}
	reality := toString(AllBitAndGitSubCommands(ShellCmd))
	for _, e := range expects {
		assert.Contains(t, reality, e)
	}
}

func TestFlagSuggestionsForCommand(t *testing.T) {
	// fixme add support for all git sub commands
	expects :=
		[]struct {
			cmd             string
			expectedOptions []string
			expectedFlags   []string
		}{
			{
				"push",
				[]string{"-f"},
				[]string{"--force", "--dry-run", "--porcelain", "--delete", "--tags"},
			},
			{
				"pull",
				[]string{"-q"},
				[]string{"--ff-only", "--no-ff", "--no-edit"},
			},
			{
				"commit",
				[]string{"-a", "-F <file>"},
				[]string{"--all", "--squash=<commit>", "--reset-author", "--branch", "--allow-empty"},
			},
		}
	for _, e := range expects {
		realityFlags := suggestionToString(FlagSuggestionsForCommand(e.cmd, "--"))
		for _, ee := range e.expectedFlags {
			assert.Contains(t, realityFlags, ee)
		}
		realityOptions := suggestionToString(FlagSuggestionsForCommand(e.cmd, "-"))
		for _, ee := range e.expectedOptions {
			assert.Contains(t, realityOptions, ee)
		}
	}
}

func BenchmarkAllBitAndGitSubCommands(b *testing.B) {
	for n := 0; n < b.N; n++ {
		AllBitAndGitSubCommands(ShellCmd)
	}
}
