package cmd

import (
	"github.com/chriswalz/complete/v3"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"github.com/thoas/go-funk"
	"time"
)

func toAutoCLI(suggs []complete.Suggestion) func(prefix string) []complete.Suggestion {
	return func(prefix string) []complete.Suggestion {
		return suggs
	}
}

func CreateSuggestionMap(cmd *cobra.Command) (*complete.CompTree, map[string]*cobra.Command) {
	start := time.Now()
	_, bitCmdMap := AllBitSubCommands(cmd)
	log.Debug().Msg((time.Now().Sub(start)).String())
	start = time.Now()
	allBitCmds := AllBitAndGitSubCommands(cmd)
	log.Debug().Msg((time.Now().Sub(start)).String())
	//commonCommands := CobraCommandToSuggestions(CommonCommandsList())
	start = time.Now()
	branchListSuggestions := BranchListSuggestions()
	log.Debug().Msg((time.Now().Sub(start)).String())
	start = time.Now()
	cobraCmdNames := CobraCommandToSuggestions(allBitCmds)
	log.Debug().Msg((time.Now().Sub(start)).String())
	start = time.Now()
	gitAddSuggestions := GitAddSuggestions()
	log.Debug().Msg((time.Now().Sub(start)).String())
	start = time.Now()
	//gitResetSuggestions := GitResetSuggestions()
	log.Debug().Msg((time.Now().Sub(start)).String())
	start = time.Now()
	gitmojiSuggestions := GitmojiSuggestions()
	log.Debug().Msg((time.Now().Sub(start)).String())

	st := b
	st.Sub["version"] = &complete.CompTree{Desc: "Print bit and git version"}
	if st.Flags == nil {
		st.Flags = map[string]*complete.CompTree{}
	}
	st.Flags["version"] = &complete.CompTree{Desc: "Print bit and git version"}
	// add dynamic predictions and bit specific commands
	st.Sub["add"].Dynamic = toAutoCLI(gitAddSuggestions)
	st.Sub["checkout"].Dynamic = toAutoCLI(branchListSuggestions)
	st.Sub["co"].Dynamic = toAutoCLI(branchListSuggestions)
	st.Sub["log"].Dynamic = toAutoCLI(branchListSuggestions)
	st.Sub["merge"].Dynamic = toAutoCLI(branchListSuggestions)
	st.Sub["rebase"].Dynamic = toAutoCLI(branchListSuggestions)
	st.Sub["release"] = &complete.CompTree{Desc: "Commit unstaged changes, bump minor tag, push",
		Args: map[string]*complete.CompTree{
			"bump":      {Desc: "increment minor version by 1"},
			"<version>": {Desc: "specify a specific version"},
		},
	}
	st.Sub["pr"] = &complete.CompTree{
		Desc:    "Check out a pull request from Github (requires GH CLI)",
		Dynamic: lazyLoad(GitHubPRSuggestions("")), // lazyLoad(GitHubPRSuggestions)), FIXME
	}
	st.Sub["gitmoji"] = &complete.CompTree{
		Desc:    "(Pre-alpha) Commit using gitmojis",
		Dynamic: toAutoCLI(gitmojiSuggestions),
	}
	st.Sub["save"] = &complete.CompTree{
		Desc: "Save your changes to your current branch",
		Flags: map[string]*complete.CompTree{
			"no-verify": {Desc: "bypass pre-commit and commit-msg hooks"},
		},
	}
	st.Sub["update"] = &complete.CompTree{Desc: "Updates bit to the latest or specified version"}
	st.Sub["complete"] = &complete.CompTree{Desc: "Add classical tab completion to bit"}
	st.Sub["sync"] = &complete.CompTree{Desc: "Synchronizes local changes with changes on origin or specified branch"}
	st.Sub["reset"].Args = map[string]*complete.CompTree{
		"HEAD~1": {Desc: "previous commit"},
	}
	st.Sub["status"].Flags["porcelain"] = &complete.CompTree{
		Args: map[string]*complete.CompTree{
			"v1": {Desc: "the first version of the porcelain format"},
			"v2": {Desc: "an updated version of the porcelain format"},
		},
	}
	st.Sub["submodule"] = &complete.CompTree{Desc: "Initialize, update or inspect submodules"}
	st.Sub["switch"] = &complete.CompTree{Desc: "Switch branches", Dynamic: toAutoCLI(branchListSuggestions)}

	for k, v := range descriptionMap {
		if st.Sub[k] == nil {
			continue
		}
		st.Sub[k].Desc = v
	}

	// dynamically add "Common Commands" & "Git aliases"
	for _, cmd := range cobraCmdNames {
		if st.Sub[cmd.Text] != nil {
			continue
		}
		st.Sub[cmd.Text] = &complete.CompTree{Desc: cmd.Description}
	}

	// fixme is this necessary?
	funk.ForEach(branchListSuggestions, func(s complete.Suggestion) {
		if descriptionMap[s.Name] != "" {
			return
		}
		descriptionMap[s.Name] = s.Desc
	})

	funk.ForEach(gitmojiSuggestions, func(s complete.Suggestion) {
		if descriptionMap[s.Name] != "" {
			return
		}
		descriptionMap[s.Name] = s.Desc
	})

	// command
	// flags
	// commands
	// value

	//completerSuggestionMap := map[string]func() []prompt.Suggest{
	//	"":         memoize([]prompt.Suggest{}),
	//	"shell":    memoize(combraCommandSuggestions),
	//	"checkout": memoize(branchListSuggestions),
	//	"switch":   memoize(branchListSuggestions),
	//	"co":       memoize(branchListSuggestions),
	//	"merge":    memoize(branchListSuggestions),
	//	"rebase":   memoize(branchListSuggestions),
	//	"log":      memoize(branchListSuggestions),
	//	"add":      memoize(gitAddSuggestions),
	//	"release": memoize([]prompt.Suggest{
	//		{Text: "bump", Description: "Increment SemVer from tags and release e.g. if latest is v0.1.2 it's bumped to v0.1.3 "},
	//		{Text: "<version>", Description: "Name of release version e.g. v0.1.2"},
	//	}),
	//	"reset":   memoize(gitResetSuggestions),
	//	"pr":      lazyLoad(GitHubPRSuggestions),
	//	"gitmoji": memoize(gitmoji),
	//	"save":    memoize(gitmoji),
	//	//"_any": commonCommands,
	//}
	return st, bitCmdMap
}

var descriptionMap = map[string]string{
	"add":             "Add file contents to the index",
	"am":              "Apply a series of patches from a mailbox",
	"archive":         "Create an archive of files from a named tree",
	"branch":          "List, create, or delete branches",
	"bisect":          "Use binary search to find the commit that introduced a bug",
	"bundle":          "Move objects and refs by archive",
	"commit":          "Record changes to the repository",
	"clone":           "Clone a repository into a new directory",
	"checkout":        "Switch branches or restore working tree files",
	"co":              "Switch branches or restore working tree files",
	"fetch":           "Download objects and refs from another repository",
	"diff":            "Show changes between commits, commit and working tree, etc",
	"cherry-pick":     "Apply the changes introduced by some existing commits",
	"citool":          "Graphical alternative to git-commit",
	"clean":           "Remove untracked files from the working tree",
	"describe":        "Give an object a human readable name based on an available ref",
	"format-patch":    "Prepare patches for e-mail submission",
	"gc":              "Cleanup unnecessary files and optimize the local repository",
	"gitk":            "The Git repository browser",
	"grep":            "Print lines matching a pattern",
	"gui":             "A portable graphical interface to Git",
	"init":            "Create an empty Git repository or reinitialize an existing one",
	"log":             "Show commit logs",
	"merge":           "Join two or more development histories together",
	"mv":              "Move or rename a file, a directory, or a symlink",
	"notes":           "Add or inspect object notes",
	"pull":            "Fetch from and integrate with another repository or a local branch",
	"push":            "Update remote refs along with associated objects",
	"range-diff":      "Compare two commit ranges (e.g. two versions of a branch)",
	"rebase":          "Reapply commits on top of another base tip",
	"reset":           "Reset current HEAD to the specified state",
	"restore":         "Restore working tree files",
	"revert":          "Revert some existing commits",
	"rm":              "Remove files from the working tree and from the index",
	"show":            "Show various types of objects",
	"stash":           "Stash the changes in a dirty working directory away",
	"shortlog":        "Summarize 'git log' output",
	"status":          "Show the working tree status",
	"submodule":       "Initialize, update or inspect submodules",
	"switch":          "Switch branches",
	"tag":             "Create, list, delete or verify a tag object signed with GPG",
	"worktree":        "Manage multiple working trees",
	"config":          "Get and set repository or global options",
	"fast-import":     "Backend for fast Git data importers",
	"filter-branch":   "Rewrite branches",
	"mergetool":       "Run merge conflict resolution tools to resolve merge conflicts",
	"pack-refs":       "Pack heads and tags for efficient repository access",
	"prune":           "Prune all unreachable objects from the object database",
	"reflog":          "Manage reflog information",
	"remote":          "Manage set of tracked repositories",
	"rename":          "",
	"remove":          "",
	"set-head":        "",
	"repack":          "Pack unpacked objects in a repository",
	"replace":         "Create, list, delete refs to replace objects",
	"annotate":        "Annotate file lines with commit information",
	"blame":           "Show what revision and author last modified each line of a file",
	"count-objects":   "Count unpacked number of objects and their disk consumption",
	"difftool":        "Show changes using common diff tools",
	"fsck":            "Verifies the connectivity and validity of the objects in the database",
	"gitweb":          "Git web interface (web frontend to Git repositories)",
	"help":            "Display help information about Git",
	"instaweb":        "Instantly browse your working repository in gitweb",
	"merge-tree":      "Show three-way merge without touching index",
	"rerere":          "Reuse recorded resolution of conflicted merges",
	"show-branch":     "Show branches and their commits",
	"verify-commit":   "Check the GPG signature of commits",
	"verify-tag":      "Check the GPG signature of tags",
	"whatchanged":     "Show logs with difference each commit introduces",
	"archimport":      "Import a GNU Arch repository into Git",
	"cvsexportcommit": "Export a single commit to a CVS checkout",
	"cvsimport":       "Salvage your data out of another SCM people love to hate",
	"cvsserver":       "A CVS server emulator for Git",
	"imap-send":       "Send a collection of patches from stdin to an IMAP folder",
	"p4":              "Import from and submit to Perforce repositories",
	"fast-export":     "Git data exporter",
	"release":         "Commit unstaged changes, bump minor tag, push",
	"pr":              "Check out a pull request from Github (requires GH CLI)",
	"info":            "Get general information about the status of your repository",
	"gitmoji":         "(Pre-alpha) Commit using gitmojis",
	"save":            "Save your changes to your current branch",
	"update":          "Updates bit to the latest or specified version",
	"complete":        "Add classical tab completion to bit",
	"sync":            "Synchronizes local changes with changes on origin or specified branch",
}
