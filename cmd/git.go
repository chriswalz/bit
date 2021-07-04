package cmd

import (
	"fmt"
	exec "golang.org/x/sys/execabs"
	"os"
	"strings"

	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

func CloudBranchExists() bool {
	msg, err := execCommand("git", "diff", "@{u}").CombinedOutput()
	if err != nil {
		log.Debug().Err(err).Send()
	}
	// log.Println("msg:", string(msg))
	// log.Println("err:", err)
	return !strings.Contains(string(msg), "fatal: no upstream configured for branch")
}

// CommitOnlyInCurrentBranch Useful for verifying that a commit is only in one branch so that it's not risky to amend it
func CommitOnlyInCurrentBranch(branch, commitId string) bool {
	msg, err := execCommand("git", "branch", "--contains", commitId).CombinedOutput()
	if err != nil {
		log.Debug().Err(err).Send()
	}
	branches := strings.Split(strings.TrimSpace(string(msg)), "\n")
	if len(branches) != 1 {
		return false
	}
	return branch == branches[0][2:]
}

func GetLastCommitId() string {
	// git rev-parse HEAD
	msg, err := execCommand("git", "rev-parse", "HEAD").CombinedOutput()
	if err != nil {
		log.Debug().Err(err).Send()
	}
	commitId := strings.TrimSpace(string(msg))
	return commitId
}

func CurrentBranch() string {
	msg, err := execCommand("git", "branch", "--show-current").CombinedOutput()
	if err != nil {
		log.Debug().Err(err).Send()
	}
	return strings.TrimSpace(string(msg))
}

func IsAheadOfCurrent() bool {
	msg, err := execCommand("git", "status", "-sb").CombinedOutput()
	if err != nil {
		log.Debug().Msg(err.Error())
	}
	return strings.Contains(string(msg), "ahead")
}

func IsGitRepo() bool {
	out, _ := execCommand("git", "rev-parse", "--is-inside-work-tree").CombinedOutput()
	return strings.TrimSpace(string(out)) == "true"
}

func IsBehindCurrent() bool {
	msg, err := execCommand("git", "status", "-sb").CombinedOutput()
	if err != nil {
		log.Debug().Err(err).Send()
	}
	return strings.Contains(string(msg), "behind")
}

func NothingToCommit() bool {
	// git diff-index HEAD --
	msg, err := execCommand("git", "diff-index", "HEAD", "--").CombinedOutput()
	if err != nil {
		log.Debug().Err(err).Send()
	}
	changedFiles := strings.Split(strings.TrimSpace(string(msg)), "\n")
	if len(changedFiles) == 1 && changedFiles[0] == "" {
		return true
	}
	return false
}

func IsDiverged() bool {
	msg, err := execCommand("git", "status").CombinedOutput()
	if err != nil {
		log.Debug().Err(err).Send()
	}
	return strings.Contains(string(msg), "have diverged")
}

func StashableChanges() bool {
	msg, err := execCommand("git", "status").CombinedOutput()
	if err != nil {
		log.Debug().Err(err).Send()
	}
	return strings.Contains(string(msg), "Changes to be committed:") || strings.Contains(string(msg), "Changes not staged for commit:")
}

func MostRecentCommonAncestorCommit(branchA, branchB string) string {
	msg, err := execCommand("git", "merge-base", branchA, branchB).CombinedOutput()
	if err != nil {
		log.Debug().Err(err).Send()
	}
	return string(msg)
}

func StashList() []string {
	msg, err := execCommand("git", "stash", "list").CombinedOutput()
	if err != nil {
		log.Debug().Err(err).Send()
	}
	return strings.Split(string(msg), "\n")
}

func refreshBranch() error {
	msg, err := execCommand("git", "pull", "--ff-only").CombinedOutput()
	if err != nil {
		return err
	}
	if strings.Contains(strings.TrimSpace(string(msg)), "Fast-forward") {
		fmt.Println("Branch was fast-forwarded by bit")
	}
	return nil
}

func refreshOnBranch(branchName string) error {
	_, err := execCommand("git", "pull", "--ff-only", branchName).CombinedOutput()
	if err != nil {
		return err
	}
	log.Debug().Msg("Branch was fast-forwarded by bit.")
	return nil
}

func branchListRaw() (string, error) {
	msg, err := execCommand("git", "for-each-ref", "--sort=-committerdate", "refs/heads/", "refs/remotes", "--format='%(authordate); %(authorname); %(color:red)%(objectname:short); %(color:yellow)%(refname:short)%(color:reset); (%(color:green)%(committerdate:relative)%(color:reset))'").CombinedOutput()
	return string(msg), err
}

func FileChangesList() []FileChange {
	msg, err := execCommand("git", "status", "--porcelain=v2").CombinedOutput()
	if err != nil {
		log.Debug().Err(err).Send()
	}
	var changes []FileChange
	// if user has an older version of git porcelain=v2 is not supported. don't show CL suggestions for now 2.7
	if strings.Contains(string(msg), "option `porcelain' takes no value") {
		return changes
	}
	list := strings.Split(string(msg), "\n")
	statusMap := map[string]string{
		"M.": "Added",
		"MM": "Partially Added",
		"A.": "New File",
		".M": "Not Staged",
		"?":  "Untracked",
	}
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

func AllGitAliases() (cc []*cobra.Command) {
	// fixme exec.Command used instead of execCommand since alias not showing with LANG=C
	msg, err := exec.Command("git", "config", "--get-regexp", "^alias").CombinedOutput()
	if err != nil {
		log.Debug().Err(err).Send()
		return cc
	}
	aliases := strings.Split(strings.TrimSpace(string(msg)), "\n")
	for _, alias := range aliases {
		if alias == "" {
			continue
		}

		split := strings.Fields(strings.TrimSpace(alias)[6:])
		if len(split) < 2 {
			continue
		}
		c := cobra.Command{
			Use:   split[0],
			Short: strings.Join(split[1:], " "),
		}
		cc = append(cc, &c)
	}

	return cc
}

func PrintGitVersion() {
	RunInTerminalWithColor("git", []string{"--version"})
}

func checkoutBranch(branch string) bool {
	msg, err := execCommand("git", "checkout", branch).CombinedOutput()
	if err != nil {
		log.Debug().Err(err).Send()
	}
	if strings.Contains(string(msg), "did not match any file") {
		return false
	}
	fmt.Println(string(msg))
	return true
}

func tagCurrentBranch(version string) error {
	msg, err := execCommand("git", "tag", version).CombinedOutput()
	if err != nil {
		return fmt.Errorf("%v: %w", string(msg), err)
	}
	return err
}

func execCommand(name string, arg ...string) *exec.Cmd {
	log.Debug().Msg(name + " " + strings.Join(arg, " "))
	c := exec.Command(name, arg...)
	c.Env = os.Environ()

	if name == "git" {
		// exec commands are parsed by bit without getting printed.
		// parsing git assumes english
		c.Env = append(c.Env, "LANG=C")
	}
	return c
}
