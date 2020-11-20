package cmd

import (
	"fmt"
	"github.com/chriswalz/complete/v2"
	"testing"

	"github.com/c-bata/go-prompt"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
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

func branchToString(list []*Branch) []string {
	newList := make([]string, len(list))
	for i, v := range list {
		newList[i] = v.FullName
	}
	return newList
}

func TestCommonCommandsList(t *testing.T) {
	expects := []string{"pull --rebase origin master", "commit -a --amend --no-edit"}
	reality := toString(CommonCommandsList())
	for _, e := range expects {
		assert.Contains(t, reality, e)
	}
}

func TestBranchList(t *testing.T) {
	expects := []string{"master", "fix-sync-upstream"}
	notexpects := []string{"origin/master", "origin/HEAD", "origin/fix-sync-upstream"}
	reality := branchToString(BranchList())
	for _, e := range expects {
		assert.Contains(t, reality, e)
	}
	for _, ne := range notexpects {
		assert.NotContains(t, reality, ne)
	}
}

func TestToStructuredBranchList(t *testing.T) {
	expects :=
		[]struct {
			raw                     string
			expectedFirstBranchName string
			expectedAuthor          string
			expectedRelativeDate    string
			expectedAbsoluteDate    string
		}{
			{
				`'Fri Sep 11 01:19:12 2020 -0400; John Doe; bf84c09; origin/other-branch; (3 days ago)'
'Fri Sep 11 01:19:12 2020 -0400; John Doe; bf84c09; origin/master; (3 days ago)'`,
				"origin/other-branch",
				"John Doe",
				"3 days ago",
				"Fri Sep 11 01:19:12 2020 -0400",
			},
			{
				`warning: ignoring broken ref refs/remotes/origin/HEAD
'Fri Sep 11 01:19:12 2020 -0400; John Doe; e5cffc5; origin/release-v2.11.0; (7 days ago)'
'Fri Sep 11 01:19:12 2020 -0400; John Doe; 2f41d5e; origin/feature_FD-5860; (8 days ago)'`,
				"origin/release-v2.11.0",
				"John Doe",
				"7 days ago",
				"Fri Sep 11 01:19:12 2020 -0400",
			},
		}
	for _, e := range expects {
		fmt.Println(e.expectedFirstBranchName)
		list := toStructuredBranchList(e.raw)
		assert.Greaterf(t, len(list), 0, e.expectedFirstBranchName)
		reality := list[0]
		assert.Equal(t, reality.FullName, e.expectedFirstBranchName)
		assert.Equal(t, reality.Author, e.expectedAuthor)
		assert.Contains(t, reality.RelativeDate, e.expectedRelativeDate)
		assert.Equal(t, reality.AbsoluteDate, e.expectedAbsoluteDate)
	}
}

// Tests AllBitAndGitSubCommands has common commands, git sub commands, git aliases, git-extras and bit commands
func TestAllBitAndGitSubCommands(t *testing.T) {
	expects := []string{"pull --rebase origin master", "commit -a --amend --no-edit", "co", "lg"}
	reality := toString(AllBitAndGitSubCommands(BitCmd))
	for _, e := range expects {
		assert.Contains(t, reality, e)
	}
}

func TestParseManPage(t *testing.T) {
	reality := parseManPage("rebase")
	assert.NotContains(t, reality, "GIT-REBASE(1)")
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
				"rebase",
				[]string{"-i"},
				[]string{"--continue", "--abort", "--merge"},
			},
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

func TestListGHPullRequests(t *testing.T) {
	expect := PullRequest{
		Title:  "Update Go install",
		Number: 32,
		State:  "open",
	}
	prs := ListGHPullRequests()
	for _, pr := range prs {
		if pr.Number == expect.Number {
			assert.Contains(t, pr.Title, expect.Title)
			return
		}
	}
	assert.Fail(t, "PR missing")
}

func TestCompletion(t *testing.T) {
	suggestionsTree, _ := CreateSuggestionMap(BitCmd)
	expects :=
		[]struct {
			line        string
			predictions []string
		}{
			{
				"bit rebase ",
				[]string{"--continue", "--abort", "--merge"},
			},
			{
				"bit push ",
				[]string{"--force", "--dry-run", "--porcelain", "--delete", "--tags"},
			},
			{
				"bit pull ",
				[]string{"--ff-only", "--no-ff", "--no-edit"},
			},
		}
	for _, e := range expects {
		reality, err := complete.CompleteLine(e.line, suggestionsTree)
		assert.Equal(t, err, nil)

		for _, p := range e.predictions {
			assert.Contains(t, reality, p)
		}
	}
}

func BenchmarkAllBitAndGitSubCommands(b *testing.B) {
	for n := 0; n < b.N; n++ {
		AllBitAndGitSubCommands(BitCmd)
	}
}
