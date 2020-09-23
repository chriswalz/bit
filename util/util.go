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

func AddCommandToShellHistory(cmd string, args []string)  {
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
		".M": "Not Staged",
		"M.": "Added",
		"MM": "Partially Added",
		"?": "Untracked",
	}
	var changes []FileChange
	for i := 0; i < len(list)-1; i++ {
		cols := strings.Fields(strings.TrimSpace(list[i]))
		b := FileChange{
			Name:   cols[len(cols)-1],
			Status: statusMap[cols[1]],
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
			Description: "Status: " + fc.Status,
		})
	}
	return suggestions
}

func CobraCommandToSuggestions(cmds []*cobra.Command) []prompt.Suggest {
	var suggestions []prompt.Suggest
	for _, branch := range cmds {
		suggestions = append(suggestions, prompt.Suggest{
			Text:        branch.Name(),
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
			Use:   split[0],
			Short: strings.TrimSpace(split[len(split)-1]),
		}
		cc = append(cc, &c)
	}

	return cc
}

func FlagSuggestions(gitSubCmd string, flagtype string) []prompt.Suggest {
	//msg, err := exec.Command("git", "help", gitSubCmd).CombinedOutput()
	//if err != nil {
	//	//fmt.Println(err)
	//}
	//op := strings.Split(string(msg), "OPTIONS")
	if gitSubCmd != "commit" && gitSubCmd != "push" && gitSubCmd != "status" {
		return []prompt.Suggest{}
	}
	str := commitFlagsStr
	if gitSubCmd == "push" {
		str = pushFlagsStr
	} else if gitSubCmd == "status" {
		str = statusFlagsStr
	} else if gitSubCmd == "commit" {
		str = commitFlagsStr
	} else {
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

var commitFlagsStr = `-a
--all
Tell the command to automatically stage files that have been modified and deleted, but new files you have not told Git about are not affected.

-p
--patch
Use the interactive patch selection interface to chose which changes to commit. See git-add[1] for details.

-C <commit>
--reuse-message=<commit>
Take an existing commit object, and reuse the log message and the authorship information (including the timestamp) when creating the commit.

-c <commit>
--reedit-message=<commit>
Like -C, but with -c the editor is invoked, so that the user can further edit the commit message.

--fixup=<commit>
Construct a commit message for use with rebase --autosquash. The commit message will be the subject line from the specified commit with a prefix of "fixup! ". See git-rebase[1] for details.

--squash=<commit>
Construct a commit message for use with rebase --autosquash. The commit message subject line is taken from the specified commit with a prefix of "squash! ". Can be used with additional commit message options (-m/-c/-C/-F). See git-rebase[1] for details.

--reset-author
When used with -C/-c/--amend options, or when committing after a conflicting cherry-pick, declare that the authorship of the resulting commit now belongs to the committer. This also renews the author timestamp.

--short
When doing a dry-run, give the output in the short-format. See git-status[1] for details. Implies --dry-run.

--branch
Show the branch and tracking info even in short-format.

--porcelain
When doing a dry-run, give the output in a porcelain-ready format. See git-status[1] for details. Implies --dry-run.

--long
When doing a dry-run, give the output in the long-format. Implies --dry-run.

-z
--null
When showing short or porcelain status output, print the filename verbatim and terminate the entries with NUL, instead of LF. If no format is given, implies the --porcelain output format. Without the -z option, filenames with "unusual" characters are quoted as explained for the configuration variable core.quotePath (see git-config[1]).

-F <file>
--file=<file>
Take the commit message from the given file. Use - to read the message from the standard input.

--author=<author>
Override the commit author. Specify an explicit author using the standard A U Thor <author@example.com> format. Otherwise <author> is assumed to be a pattern and is used to search for an existing commit by that author (i.e. rev-list --all -i --author=<author>); the commit author is then copied from the first such commit found.

--date=<date>
Override the author date used in the commit.

-m <msg>
--message=<msg>
Use the given <msg> as the commit message. If multiple -m options are given, their values are concatenated as separate paragraphs.

The -m option is mutually exclusive with -c, -C, and -F.

-t <file>
--template=<file>
When editing the commit message, start the editor with the contents in the given file. The commit.template configuration variable is often used to give this option implicitly to the command. This mechanism can be used by projects that want to guide participants with some hints on what to write in the message in what order. If the user exits the editor without editing the message, the commit is aborted. This has no effect when a message is given by other means, e.g. with the -m or -F options.

-s
--signoff
Add Signed-off-by line by the committer at the end of the commit log message. The meaning of a signoff depends on the project, but it typically certifies that committer has the rights to submit this work under the same license and agrees to a Developer Certificate of Origin (see http://developercertificate.org/ for more information).

-n
--no-verify
This option bypasses the pre-commit and commit-msg hooks. See also githooks[5].

--allow-empty
Usually recording a commit that has the exact same tree as its sole parent commit is a mistake, and the command prevents you from making such a commit. This option bypasses the safety, and is primarily for use by foreign SCM interface scripts.

--allow-empty-message
Like --allow-empty this command is primarily for use by foreign SCM interface scripts. It allows you to create a commit with an empty commit message without using plumbing commands like git-commit-tree[1].

--cleanup=<mode>
This option determines how the supplied commit message should be cleaned up before committing. The <mode> can be strip, whitespace, verbatim, scissors or default.

strip
Strip leading and trailing empty lines, trailing whitespace, commentary and collapse consecutive empty lines.

whitespace
Same as strip except #commentary is not removed.

verbatim
Do not change the message at all.

scissors
Same as whitespace except that everything from (and including) the line found below is truncated, if the message is to be edited. "#" can be customized with core.commentChar.

# ------------------------ >8 ------------------------
default
Same as strip if the message is to be edited. Otherwise whitespace.

The default can be changed by the commit.cleanup configuration variable (see git-config[1]).

-e
--edit
The message taken from file with -F, command line with -m, and from commit object with -C are usually used as the commit log message unmodified. This option lets you further edit the message taken from these sources.

--no-edit
Use the selected commit message without launching an editor. For example, git commit --amend --no-edit amends a commit without changing its commit message.

--amend
Replace the tip of the current branch by creating a new commit. The recorded tree is prepared as usual (including the effect of the -i and -o options and explicit pathspec), and the message from the original commit is used as the starting point, instead of an empty message, when no other message is specified from the command line via options such as -m, -F, -c, etc. The new commit has the same parents and author as the current one (the --reset-author option can countermand this).

It is a rough equivalent for:

	$ git reset --soft HEAD^
	$ ... do something else to come up with the right tree ...
	$ git commit -c ORIG_HEAD
but can be used to amend a merge commit.

You should understand the implications of rewriting history if you amend a commit that has already been published. (See the "RECOVERING FROM UPSTREAM REBASE" section in git-rebase[1].)

--no-post-rewrite
Bypass the post-rewrite hook.

-i
--include
Before making a commit out of staged contents so far, stage the contents of paths given on the command line as well. This is usually not what you want unless you are concluding a conflicted merge.

-o
--only
Make a commit by taking the updated working tree contents of the paths specified on the command line, disregarding any contents that have been staged for other paths. This is the default mode of operation of git commit if any paths are given on the command line, in which case this option can be omitted. If this option is specified together with --amend, then no paths need to be specified, which can be used to amend the last commit without committing changes that have already been staged. If used together with --allow-empty paths are also not required, and an empty commit will be created.

--pathspec-from-file=<file>
Pathspec is passed in <file> instead of commandline args. If <file> is exactly - then standard input is used. Pathspec elements are separated by LF or CR/LF. Pathspec elements can be quoted as explained for the configuration variable core.quotePath (see git-config[1]). See also --pathspec-file-nul and global --literal-pathspecs.

--pathspec-file-nul
Only meaningful with --pathspec-from-file. Pathspec elements are separated with NUL character and all other characters are taken literally (including newlines and quotes).

-u[<mode>]
--untracked-files[=<mode>]
Show untracked files.

The mode parameter is optional (defaults to all), and is used to specify the handling of untracked files; when -u is not used, the default is normal, i.e. show untracked files and directories.

The possible options are:

no - Show no untracked files

normal - Shows untracked files and directories

all - Also shows individual files in untracked directories.

The default can be changed using the status.showUntrackedFiles configuration variable documented in git-config[1].

-v
--verbose
Show unified diff between the HEAD commit and what would be committed at the bottom of the commit message template to help the user describe the commit by reminding what changes the commit has. Note that this diff output doesn’t have its lines prefixed with #. This diff will not be a part of the commit message. See the commit.verbose configuration variable in git-config[1].

If specified twice, show in addition the unified diff between what would be committed and the worktree files, i.e. the unstaged changes to tracked files.

-q
--quiet
Suppress commit summary message.

--dry-run
Do not create a commit, but show a list of paths that are to be committed, paths with local changes that will be left uncommitted and paths that are untracked.

--status
Include the output of git-status[1] in the commit message template when using an editor to prepare the commit message. Defaults to on, but can be used to override configuration variable commit.status.

--no-status
Do not include the output of git-status[1] in the commit message template when using an editor to prepare the default commit message.

-S[<keyid>]
--gpg-sign[=<keyid>]
--no-gpg-sign
GPG-sign commits. The keyid argument is optional and defaults to the committer identity; if specified, it must be stuck to the option without a space. --no-gpg-sign is useful to countermand both commit.gpgSign configuration variable, and earlier --gpg-sign.

--
Do not interpret any more arguments as options.

<pathspec>…​
When pathspec is given on the command line, commit the contents of the files that match the pathspec without recording the changes already added to the index. The contents of these files are also staged for the next commit on top of what have been staged before.
`

var pushFlagsStr = `
--all
Push all branches (i.e. refs under refs/heads/); cannot be used with other <refspec>.

--prune
Remove remote branches that don’t have a local counterpart. For example a remote branch tmp will be removed if a local branch with the same name doesn’t exist any more. This also respects refspecs, e.g. git push --prune remote refs/heads/*:refs/tmp/* would make sure that remote refs/tmp/foo will be removed if refs/heads/foo doesn’t exist.

--mirror
Instead of naming each ref to push, specifies that all refs under refs/ (which includes but is not limited to refs/heads/, refs/remotes/, and refs/tags/) be mirrored to the remote repository. Newly created local refs will be pushed to the remote end, locally updated refs will be force updated on the remote end, and deleted refs will be removed from the remote end. This is the default if the configuration option remote.<remote>.mirror is set.

-n
--dry-run
Do everything except actually send the updates.

--porcelain
Produce machine-readable output. The output status line for each ref will be tab-separated and sent to stdout instead of stderr. The full symbolic names of the refs will be given.

-d
--delete
All listed refs are deleted from the remote repository. This is the same as prefixing all refs with a colon.

--tags
All refs under refs/tags are pushed, in addition to refspecs explicitly listed on the command line.

--follow-tags
Push all the refs that would be pushed without this option, and also push annotated tags in refs/tags that are missing from the remote but are pointing at commit-ish that are reachable from the refs being pushed. This can also be specified with configuration variable push.followTags. For more information, see push.followTags in git-config[1].

--[no-]signed
--signed=(true|false|if-asked)
GPG-sign the push request to update refs on the receiving side, to allow it to be checked by the hooks and/or be logged. If false or --no-signed, no signing will be attempted. If true or --signed, the push will fail if the server does not support signed pushes. If set to if-asked, sign if and only if the server supports signed pushes. The push will also fail if the actual call to gpg --sign fails. See git-receive-pack[1] for the details on the receiving end.

--[no-]atomic
Use an atomic transaction on the remote side if available. Either all refs are updated, or on error, no refs are updated. If the server does not support atomic pushes the push will fail.

-o <option>
--push-option=<option>
Transmit the given string to the server, which passes them to the pre-receive as well as the post-receive hook. The given string must not contain a NUL or LF character. When multiple --push-option=<option> are given, they are all sent to the other side in the order listed on the command line. When no --push-option=<option> is given from the command line, the values of configuration variable push.pushOption are used instead.

--receive-pack=<git-receive-pack>
--exec=<git-receive-pack>
Path to the git-receive-pack program on the remote end. Sometimes useful when pushing to a remote repository over ssh, and you do not have the program in a directory on the default $PATH.

--[no-]force-with-lease
--force-with-lease=<refname>
--force-with-lease=<refname>:<expect>
Usually, "git push" refuses to update a remote ref that is not an ancestor of the local ref used to overwrite it.

This option overrides this restriction if the current value of the remote ref is the expected value. "git push" fails otherwise.

Imagine that you have to rebase what you have already published. You will have to bypass the "must fast-forward" rule in order to replace the history you originally published with the rebased history. If somebody else built on top of your original history while you are rebasing, the tip of the branch at the remote may advance with her commit, and blindly pushing with --force will lose her work.

This option allows you to say that you expect the history you are updating is what you rebased and want to replace. If the remote ref still points at the commit you specified, you can be sure that no other people did anything to the ref. It is like taking a "lease" on the ref without explicitly locking it, and the remote ref is updated only if the "lease" is still valid.

--force-with-lease alone, without specifying the details, will protect all remote refs that are going to be updated by requiring their current value to be the same as the remote-tracking branch we have for them.

--force-with-lease=<refname>, without specifying the expected value, will protect the named ref (alone), if it is going to be updated, by requiring its current value to be the same as the remote-tracking branch we have for it.

--force-with-lease=<refname>:<expect> will protect the named ref (alone), if it is going to be updated, by requiring its current value to be the same as the specified value <expect> (which is allowed to be different from the remote-tracking branch we have for the refname, or we do not even have to have such a remote-tracking branch when this form is used). If <expect> is the empty string, then the named ref must not already exist.

Note that all forms other than --force-with-lease=<refname>:<expect> that specifies the expected current value of the ref explicitly are still experimental and their semantics may change as we gain experience with this feature.

"--no-force-with-lease" will cancel all the previous --force-with-lease on the command line.

A general note on safety: supplying this option without an expected value, i.e. as --force-with-lease or --force-with-lease=<refname> interacts very badly with anything that implicitly runs git fetch on the remote to be pushed to in the background, e.g. git fetch origin on your repository in a cronjob.

The protection it offers over --force is ensuring that subsequent changes your work wasn’t based on aren’t clobbered, but this is trivially defeated if some background process is updating refs in the background. We don’t have anything except the remote tracking info to go by as a heuristic for refs you’re expected to have seen & are willing to clobber.

If your editor or some other system is running git fetch in the background for you a way to mitigate this is to simply set up another remote:

git remote add origin-push $(git config remote.origin.url)
git fetch origin-push
Now when the background process runs git fetch origin the references on origin-push won’t be updated, and thus commands like:

git push --force-with-lease origin-push
Will fail unless you manually run git fetch origin-push. This method is of course entirely defeated by something that runs git fetch --all, in that case you’d need to either disable it or do something more tedious like:

git fetch              # update 'master' from remote
git tag base master    # mark our base point
git rebase -i master   # rewrite some commits
git push --force-with-lease=master:base master:master
I.e. create a base tag for versions of the upstream code that you’ve seen and are willing to overwrite, then rewrite history, and finally force push changes to master if the remote version is still at base, regardless of what your local remotes/origin/master has been updated to in the background.

-f
--force
Usually, the command refuses to update a remote ref that is not an ancestor of the local ref used to overwrite it. Also, when --force-with-lease option is used, the command refuses to update a remote ref whose current value does not match what is expected.

This flag disables these checks, and can cause the remote repository to lose commits; use it with care.

Note that --force applies to all the refs that are pushed, hence using it with push.default set to matching or with multiple push destinations configured with remote.*.push may overwrite refs other than the current branch (including local refs that are strictly behind their remote counterpart). To force a push to only one branch, use a + in front of the refspec to push (e.g git push origin +master to force a push to the master branch). See the <refspec>... section above for details.

--repo=<repository>
This option is equivalent to the <repository> argument. If both are specified, the command-line argument takes precedence.

-u
--set-upstream
For every branch that is up to date or successfully pushed, add upstream (tracking) reference, used by argument-less git-pull[1] and other commands. For more information, see branch.<name>.merge in git-config[1].

--[no-]thin
These options are passed to git-send-pack[1]. A thin transfer significantly reduces the amount of sent data when the sender and receiver share many of the same objects in common. The default is --thin.

-q
--quiet
Suppress all output, including the listing of updated refs, unless an error occurs. Progress is not reported to the standard error stream.

-v
--verbose
Run verbosely.

--progress
Progress status is reported on the standard error stream by default when it is attached to a terminal, unless -q is specified. This flag forces progress status even if the standard error stream is not directed to a terminal.

--no-recurse-submodules
--recurse-submodules=check|on-demand|only|no
May be used to make sure all submodule commits used by the revisions to be pushed are available on a remote-tracking branch. If check is used Git will verify that all submodule commits that changed in the revisions to be pushed are available on at least one remote of the submodule. If any commits are missing the push will be aborted and exit with non-zero status. If on-demand is used all submodules that changed in the revisions to be pushed will be pushed. If on-demand was not able to push all necessary revisions it will also be aborted and exit with non-zero status. If only is used all submodules will be recursively pushed while the superproject is left unpushed. A value of no or using --no-recurse-submodules can be used to override the push.recurseSubmodules configuration variable when no submodule recursion is required.

--no-verify
Toggle the pre-push hook (see githooks[5]). The default is --verify, giving the hook a chance to prevent the push. With --no-verify, the hook is bypassed completely.

-4
--ipv4
Use IPv4 addresses only, ignoring IPv6 addresses.

-6
--ipv6
Use IPv6 addresses only, ignoring IPv4 addresses.
`

var statusFlagsStr = `-s
--short
Give the output in the short-format.

-b
--branch
Show the branch and tracking info even in short-format.

--show-stash
Show the number of entries currently stashed away.

--porcelain[=<version>]
Give the output in an easy-to-parse format for scripts. This is similar to the short output, but will remain stable across Git versions and regardless of user configuration. See below for details.

The version parameter is used to specify the format version. This is optional and defaults to the original version v1 format.

--long
Give the output in the long-format. This is the default.

-v
--verbose
In addition to the names of files that have been changed, also show the textual changes that are staged to be committed (i.e., like the output of git diff --cached). If -v is specified twice, then also show the changes in the working tree that have not yet been staged (i.e., like the output of git diff).

-u[<mode>]
--untracked-files[=<mode>]
Show untracked files.

The mode parameter is used to specify the handling of untracked files. It is optional: it defaults to all, and if specified, it must be stuck to the option (e.g. -uno, but not -u no).

The possible options are:

no - Show no untracked files.

normal - Shows untracked files and directories.

all - Also shows individual files in untracked directories.

When -u option is not used, untracked files and directories are shown (i.e. the same as specifying normal), to help you avoid forgetting to add newly created files. Because it takes extra work to find untracked files in the filesystem, this mode may take some time in a large working tree. Consider enabling untracked cache and split index if supported (see git update-index --untracked-cache and git update-index --split-index), Otherwise you can use no to have git status return more quickly without showing untracked files.

The default can be changed using the status.showUntrackedFiles configuration variable documented in git-config[1].

--ignore-submodules[=<when>]
Ignore changes to submodules when looking for changes. <when> can be either "none", "untracked", "dirty" or "all", which is the default. Using "none" will consider the submodule modified when it either contains untracked or modified files or its HEAD differs from the commit recorded in the superproject and can be used to override any settings of the ignore option in git-config[1] or gitmodules[5]. When "untracked" is used submodules are not considered dirty when they only contain untracked content (but they are still scanned for modified content). Using "dirty" ignores all changes to the work tree of submodules, only changes to the commits stored in the superproject are shown (this was the behavior before 1.7.0). Using "all" hides all changes to submodules (and suppresses the output of submodule summaries when the config option status.submoduleSummary is set).

--ignored[=<mode>]
Show ignored files as well.

The mode parameter is used to specify the handling of ignored files. It is optional: it defaults to traditional.

The possible options are:

traditional - Shows ignored files and directories, unless --untracked-files=all is specified, in which case individual files in ignored directories are displayed.

no - Show no ignored files.

matching - Shows ignored files and directories matching an ignore pattern.

When matching mode is specified, paths that explicitly match an ignored pattern are shown. If a directory matches an ignore pattern, then it is shown, but not paths contained in the ignored directory. If a directory does not match an ignore pattern, but all contents are ignored, then the directory is not shown, but all contents are shown.

-z
Terminate entries with NUL, instead of LF. This implies the --porcelain=v1 output format if no other format is given.

--column[=<options>]
--no-column
Display untracked files in columns. See configuration variable column.status for option syntax.--column and --no-column without options are equivalent to always and never respectively.

--ahead-behind
--no-ahead-behind
Display or do not display detailed ahead/behind counts for the branch relative to its upstream branch. Defaults to true.

--renames
--no-renames
Turn on/off rename detection regardless of user configuration. See also git-diff[1] --no-renames.

--find-renames[=<n>]
Turn on rename detection, optionally setting the similarity threshold. See also git-diff[1] --find-renames.`
