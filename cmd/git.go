package cmd

import (
	"fmt"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"os/exec"
	"strings"
)

func CloudBranchExists() bool {
	msg, err := execCommand("git", "pull").CombinedOutput()
	if err != nil {
		log.Debug().Err(err)
	}
	//log.Println("msg:", string(msg))
	//log.Println("err:", err)
	return !strings.Contains(string(msg), "There is no tracking information for the current branch")
}

func CurrentBranch() string {
	msg, err := execCommand("git", "branch", "--show-current").CombinedOutput()
	if err != nil {
		log.Debug().Err(err)
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
	_, err := execCommand("git", "status").CombinedOutput()
	if err != nil {
		return false
	}
	return true
}

func IsBehindCurrent() bool {
	msg, err := execCommand("git", "status", "-sb").CombinedOutput()
	if err != nil {
		log.Debug().Err(err)
	}
	return strings.Contains(string(msg), "behind")
}

func NothingToCommit() bool {
	msg, err := execCommand("git", "status").CombinedOutput()
	if err != nil {
		log.Debug().Err(err)
	}
	return strings.Contains(string(msg), "nothing to commit")
}

func IsDiverged() bool {
	msg, err := execCommand("git", "status").CombinedOutput()
	if err != nil {
		log.Debug().Err(err)
	}
	return strings.Contains(string(msg), "have diverged")
}

func StashableChanges() bool {
	msg, err := execCommand("git", "status").CombinedOutput()
	if err != nil {
		log.Debug().Err(err)
	}
	return strings.Contains(string(msg), "Changes to be committed:") || strings.Contains(string(msg), "Changes not staged for commit:")
}

func MostRecentCommonAncestorCommit(branchA, branchB string) string {
	msg, err := execCommand("git", "merge-base", branchA, branchB).CombinedOutput()
	if err != nil {
		log.Debug().Err(err)
	}
	return string(msg)
}

func StashList() []string {
	msg, err := execCommand("git", "stash", "list").CombinedOutput()
	if err != nil {
		log.Debug().Err(err)
	}
	return strings.Split(string(msg), "\n")
}

func refreshBranch() error {
	msg, err := execCommand("git", "pull", "--ff-only").CombinedOutput()
	if err != nil {
		return err
	}
	if strings.TrimSpace(string(msg)) == "Already up to date." {
		return nil
	}
	log.Debug().Msg("Branch was fast-forwarded")
	return nil
}

func refreshOnBranch(branchName string) error {
	_, err := execCommand("git", "pull", "--ff-only", branchName).CombinedOutput()
	if err != nil {
		return err
	}
	log.Debug().Msg("Branch was fast-forwarded")
	return nil
}

func branchListRaw() (string, error) {
	msg, err := execCommand("git", "for-each-ref", "--sort=-committerdate", "refs/heads/", "refs/remotes", "--format='%(authordate); %(authorname); %(color:red)%(objectname:short); %(color:yellow)%(refname:short)%(color:reset); (%(color:green)%(committerdate:relative)%(color:reset))'").CombinedOutput()
	return string(msg), err
}

func FileChangesList() []FileChange {
	msg, err := execCommand("git", "status", "--porcelain=v2").CombinedOutput()
	if err != nil {
		log.Debug().Err(err)
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
	msg, err := execCommand("git", "config", "--get-regexp", "^alias").CombinedOutput()
	if err != nil {
		log.Debug().Err(err)
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
	msg, err := execCommand("git", "--version").CombinedOutput()
	if err != nil {
		log.Debug().Err(err)
	}
	log.Debug().Msg(string(msg))
}

func checkoutBranch(branch string) bool {
	msg, err := execCommand("git", "checkout", branch).CombinedOutput()
	if err != nil {
		log.Debug().Err(err)
	}
	return !strings.Contains(string(msg), "did not match any file")
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
	return exec.Command(name, arg...)
}