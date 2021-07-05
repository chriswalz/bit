![Twitter Follow](https://img.shields.io/twitter/follow/chriswalz___?style=social)

[![Tweet](https://img.shields.io/twitter/url/http/shields.io.svg?style=social)](https://twitter.com/intent/tweet?text=Bit%20-%20a%20modern%20git%20cli%20&url=https://github.com/chriswalz/bit&hashtags=bit,git,cli,developers)
<p align="center">
<img
    src="https://user-images.githubusercontent.com/6971318/97790948-af52f200-1ba3-11eb-8b9e-a614e44da12c.gif"
    width="600px" border="0" alt="bit">
<br>
<img alt="GitHub release (latest SemVer)" src="https://img.shields.io/github/v/release/chriswalz/bit?color=gree">
<a href="https://goreportcard.com/report/github.com/chriswalz/bit"><img src="https://goreportcard.com/badge/github.com/chriswalz/bit" alt="Version"></a>
<a href="#"><img src="https://img.shields.io/github/go-mod/go-version/chriswalz/bit" alt="Version"></a>
<a href="#"><img src="https://img.shields.io/github/stars/chriswalz/bit?style=social" alt="Version"></a>
</p>


`bit` is an experimental modernized git CLI built on top of git that provides happy defaults and other niceties:

- command and **flag suggestions** to help you navigate the plethora of options git provides you
- autocompletion for files and branch names when using 
`bit add` or `bit checkout`
- automatic fetch and **branch fast-forwarding** reducing the likelihood of merge conflicts 
- suggestions **work with git aliases**
- new commands like `bit sync` that vastly simplify your workflow 
- commands from **git-extras** such as `bit release` & `bit info`
- **fully compatible with git** allowing you to fallback to git if need be.  
- get insight into how bit works using `bit --debug`.

--- **New** ---

- [X] `bit pr` view and checkout pull requests from Github (Requires GitHub CLI)
- [X] easily update bit using `bit update`
- [X] single binary
- [X] much more suggestions available! (Roughly 10x more)
- [X] Install with homebrew & macports
- [X] Interactive prompt with env variable: BIT_INTERACTIVE=true

--- **Coming Soon** ---
- bit anticipates when you'll need to type git status and will display it proactively
- `bit fix` for all the times you did something you really wish you didn't
- improved `bit sync`
- QOL improvements when switching branches or deleting tags

## Installation

Click [here](#how-to-install) for installation instructions

## Update

run `bit update` 

Customization: 
- `BIT_THEME=inverted`
- `BIT_THEME=monochrome`

###### Common commands at your finger tips 

<p align="center">
<img
    src="https://user-images.githubusercontent.com/6971318/95157964-6eb0b600-0768-11eb-8f8a-075e2987fde8.gif"
    width="600px" border="0" alt="bit">
<br>
</p>

###### Commit, bump a tag and push with a single command 

<p align="center">
<img
    src="https://user-images.githubusercontent.com/6971318/95157973-753f2d80-0768-11eb-8ef6-31239c76d305.gif"
    width="600px" border="0" alt="bit">
<br>
</p>

###### Instant git statistics and config information 

<p align="center">
<img
    src="https://user-images.githubusercontent.com/6971318/95158539-d7e4f900-0769-11eb-8530-9ffc4610a71a.gif"
    width="600px" border="0" alt="bit">
<br>
</p>

## Bit specific command Usage

Create a new commit *(roughly equivalent to `git commit -am "commit message"`)*

`bit save [commit message]`

Save your changes to the current branch [amends current commit when ahead of origin]
*(roughly equivalent to `git commit -a` or conditionally `git commit -a --amend --no-edit`)*

`bit save`

Synchronize your changes to origin branch (Beta)
*(roughly equivalent to `git pull -r; git push`)*

`bit sync`

*(roughly equivalent to `git pull -r; git push; git pull -r origin master; git push`)*
`bit sync origin master`

You have access to ALL git commands as well.

`bit commit -m "I can still use git commands"`, `bit pull -r origin master`

## Example Workflow
`bit switch example-branch`
Branch does not exist. Do you want to create it? Y/n

yes

Switched to a new branch 'example-branch'

[Makes some changes]

`bit save "add important feature"`

*for multiline commits simply don't put the final quote until you're done typing*

[fix an error for important feature]

`bit save`

[push changes to origin]

`bit sync`

[two days later confirm your branch is in sync with origin branch]

`bit sync`

[rebase your changes from origin master]

`bit sync origin master`

## Donate 

If you would like to support the development of bit, consider [sponsoring](https://github.com/sponsors/chriswalz) me.


## Principles 

1. Think in the age of the cloud
1. Embed the spirit of modern day workflows
1. Favor simplicity over complexity 
1. Bit should have happy defaults
1. Bit must be fully compatible with Git

## Inspiration

Thanks to [Gitless](https://gitless.com/), [git-extras](https://github.com/tj/git-extras), researchers in the field and of course the developers of `git` itself! Also, thanks to [go-prompt](https://github.com/c-bata/go-prompt) for the interactive prompt library

## News 
- https://news.ycombinator.com/item?id=24751212
- https://www.tldrnewsletter.com/archives/20201007
- https://www.reddit.com/r/golang/comments/j5wggn/bit_an_experimental_git_cli_with_a_powerful/
- https://b.hatena.ne.jp/entry/s/github.com/chriswalz/bit
- https://news.hada.io/topic?id=2990
- https://twitter.com/newsycombinator/status/1315517850954727424
- https://forum.devtalk.com/t/bit-a-modernized-git-cli-written-in-go/3065
- https://gocn.vip/topics/11063
- https://golangweekly.com/issues/333
- https://archive.faabli.com/archive/2020-10-09 
- https://www.wykop.pl/wpis/52945683/unknownews-wolanie-nie-dziala-zapisz-sie-lepiej-na/
- https://blog.csdn.net/a419240016/article/details/109178001

## Changelog
v1.1.2
- [X] enhancement: add `bit sw` as alias for `bit switch`
- [X] fix: bit save will amend commits only when the commit doesn't exist in any other branch
v1.1
- [X] enhancement: enable interactive prompt (keep bit running) with env variable: BIT_INTERACTIVE=true

v1.0
- [X] enhancement: significantly more autocompletions
- [X] enhancement: use fuzzy search for branch suggestions
- [X] refactor: completions now represented in tree data structure
- [X] fix: bit save no longer shows debug error outside debug mode
- [X] fix: use --is-inside-work-tree to determine whether inside a git repo
- [X] fix: gracefully handle bad release tags for `bit release bump`
- [X] fix: bit pr not listing PR in some cases
- [X] security: fix vuln on Windows where an exe in a malicious repository could run arbitrary code. Special thanks to RyotaK - https://github.com/Ry0taK for identifying this issue

v0.9
- [X] enhancement: improve bit sync behavior including `bit sync <upstream> <branch>`
- [X] enhancement: bit sync provides a rebase option for diverged branches`
- [X] fix: enable compatibility with non-english languages 
- [X] enhancement: `bit merge` suggestions

v0.8
- [X] feature: checkout Pull Requests from github (requires `gh pr list` to work)
- [X] enhancement: install with homebrew: `brew install bit-git`
- [X] enhancement: Color themes `BIT_THEME=inverted` or `BIT_THEME=monochrome` light terminal backgrounds
- [X] fix: bit clone outside a git repo
- [X] enhancement: bit is now a single binary

v0.7
- [X] feature: update your cli with `bit update`

v0.6
- [X] fix: improved git compatibility for older versions of git 
- [X] feature: emacs key binds 
- [X] feature: relative and absolute branch times
- [X] feature: completions for rebase & log
- [X] enhancement: smarter suggestions
- [X] fix: show proper descriptions for some flags
- [X] fix: prevent panic on classical tab completion for some users

v0.5

- [X] `bit switch`, `bit co`, `bit checkout` will show prompt 
- [X] fix bit tab completion (bitcomplete)
- [X] fix edge case where there is an invalid ref
- [X] various minor fixes
- [X] more completions
- [X] better suggestion filtering
- [X] absolute and relative times for branch suggestions

v0.4

- [X] multiline support with Go Survey Library
- [X] color mitigation to have roughly similar colors across OSs 
- [X] fix README go get installation instructions
- [X] QOL improvements for `bit sync`

## How to uninstall
*go binaries are self-contained so uninstalling simply requires deleting the binary(ies)*

```
rm `which bit`
```

If you ran `bit complete` optionally remove a line from your `bash_profile`, `.zshrc` etc.

`complete -o nospace -C /Users/{_USER_}/go/bin/bit bit`

## How to install

### using `cURL` (Simplest way to install)

Like bit? [Sponsor](https://github.com/sponsors/chriswalz) bit for $5

```shell script
curl -sf https://gobinaries.com/chriswalz/bit | sh;
bit complete;
echo "Type bit then press <ENTER> to show interactive prompt"
bit;
```

To overwrite installation location

`export PREFIX=/opt/bit/git && mkdir -p ${PREFIX}  ## optional: override default install location /usr/local/bin`


`bit`, `bit checkout` & `bit switch` will show interactive prompts after you press ENTER

### using `go` (Harder way to install)
*Caveats: GOPATH and GOBIN need to be set. Verify with `go env`. If they are not set, add this to your .bashrc or .bash_profile etc. AND open new terminal*
```shell script
export GOPATH=$HOME/go
export GOBIN=$(go env GOPATH)/bin
```

```shell script
GO111MODULE=on go get github.com/chriswalz/bit@latest;
bit complete
```

### using `Homebrew` (For MacOS users)

```shell script
brew install bit-git
bit complete
bit
```

Not working? Try `brew doctor`

### using `MacPorts` (For MacOS users)

```shell script
sudo port selfupdate
sudo port install bit
```


#### using `go` (For Windows Users)
```shell script
go env -w GO111MODULE=on

# if latest is not working, replace it with the latest tag found here https://github.com/chriswalz/bit/releases
go get github.com/chriswalz/bit@latest; 
bit
```

#### using `Chocolatey` (For Windows Users)
```shell script
choco install bit-git
```

#### using `zinit`
```shell script
zinit ice lucit wait"0" as"program" from"gh-r" pick"bit"
zinit light "chriswalz/bit"
```

*Note*: On Windows only the interactive prompt completion works not classic tab completion

#### using AUR (For Arch Linux Users)
For building a stable version from source, use the [`bit` package](https://aur.archlinux.org/packages/bit)

For building the latest git version from source, use the [`bit-git` package](https://aur.archlinux.org/packages/bit-git)

*Note*: These Packages are community-driven and not offically published my the bit maintainer.

Verify installation with:

`bit`

Dependencies: Git

Platform Support:
- iTerm2 (macOS)
- Terminal.app (macOS)
- Command Prompt (Windows)
- WSL/Windows Subsystem for Linux (Windows)
- gnome-terminal (Ubuntu)
