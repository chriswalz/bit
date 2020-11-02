package cmd

import (
	"github.com/rs/zerolog/log"
	"github.com/thoas/go-funk"
	"os/exec"
	"strconv"
	"strings"
)

type PullRequest struct {
	Title        string
	Number       int
	AuthorBranch string
	State        string
}

func ListGHPullRequests() []*PullRequest {
	if !GHCliExistsAndLoggedIn() {
		return []*PullRequest{
			{AuthorBranch: "Either no PRs or GH CLI not installed"},
		}
	}
	msg, err := execCommand("gh", "pr", "list", "--limit=200", "--state=all").CombinedOutput()
	if err != nil {
		log.Debug().Err(err)
	}
	prsraw := strings.Split(strings.TrimSpace(string(msg)), "\n")
	return funk.Map(prsraw, func(raw string) *PullRequest {
		fields := strings.Split(raw, "\t")
		number, _ := strconv.Atoi(fields[0])
		return &PullRequest{
			Number:       number,
			Title:        fields[1],
			AuthorBranch: fields[2],
			State:        fields[3],
		}
	}).([]*PullRequest)
}

func checkoutPullRequest(pr int) {
	if !GHCliExistsAndLoggedIn() {
		return
	}
	_, err := execCommand("gh", "pr", "checkout", strconv.Itoa(pr)).CombinedOutput()
	if err != nil {
		log.Debug().Err(err)
	}
}

func GHCliExistsAndLoggedIn() bool {
	_, err := exec.LookPath("gh")
	if err != nil {
		log.Debug().Msg("GitHub CLI doesn't exist,")
		return false
	}
	//msg , err := execCommand("gh", "auth", "status").CombinedOutput()
	//if err != nil {
	//	log.Debug().Msg(string(msg) + err.Error())
	//	return false
	//}
	//if strings.Contains(string(msg), "You are not logged into any GitHub") {
	//	log.Debug().Msg("not logged into to Github")
	//	return false
	//}
	return true
}
