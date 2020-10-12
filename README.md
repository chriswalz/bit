[![Tweet](https://img.shields.io/twitter/url/http/shields.io.svg?style=social)](https://twitter.com/intent/tweet?text=Bit%20-%20a%20modern%20git%20cli%20&url=https://github.com/chriswalz/bit&hashtags=bit,git,cli,developers)
<p align="center">
<img
    src="https://user-images.githubusercontent.com/6971318/95147815-cd1d6a80-074f-11eb-8265-56466ac628f8.gif"
    width="600px" border="0" alt="bit">
<br>
<a href="https://goreportcard.com/report/github.com/chriswalz/bit"><img src="https://goreportcard.com/badge/github.com/chriswalz/bit" alt="Version"></a>
<a href="#"><img src="https://img.shields.io/github/go-mod/go-version/chriswalz/bit" alt="Version"></a>
<a href="#"><img src="https://img.shields.io/github/stars/chriswalz/bit?style=social" alt="Version"></a>
<a href="https://github.com/chriswalz/bit/tags"><img src="https://img.shields.io/github/v/tag/chriswalz/bit?sort=semver" alt="Version"></a>
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

## Install

### using `cURL` (Simplest way to install) 

```shell script
curl -sf https://gobinaries.com/chriswalz/bit | sh;
curl -sf https://gobinaries.com/chriswalz/bit/bitcomplete | sh && echo y | COMP_INSTALL=1 bitcomplete
```

### using `go` (Harder way to install)
*Caveats: GOPATH and GOBIN need to be set. Verify with `go env`. If they are not set, add this to your .bashrc or .bash_profile etc. AND open new terminal*
```shell script
export GOPATH=$HOME/go
export GOBIN=$(go env GOPATH)/bin
```

```shell script
GO111MODULE=on go get github.com/chriswalz/bit@latest;
GO111MODULE=on go get github.com/chriswalz/bit/bitcomplete@latest;
COMP_INSTALL=1 bitcomplete;
```

#### using `go` (For Windows Users) 
```shell script
go env -w GO111MODULE=on

# if latest is not working, replace it with the latest tag found here https://github.com/chriswalz/bit/releases
go get github.com/chriswalz/bit@latest; 
bit
```

*Note*: On Windows only the interactive prompt completion works not classic tab completion

Verify installation with:

`bit`

Dependencies: Git

Platform Support:
- iTerm2 (macOS)
- Terminal.app (macOS)
- Command Prompt (Windows)
- WSL/Windows Subsystem for Linux (Windows)
- gnome-terminal (Ubuntu)

## Bit specific command Usage 

Create a new commit

`bit save [commit message]`

Save your changes to the current branch [amends current commit when ahead of origin]

`bit save` 

Synchronize your changes to origin branch (Beta)

`bit sync`

You have access to ALL git commands as well. 90% of the time the above commands will have you covered. 

`bit commit -m "I can still use git commands"`, `bit pull -r origin master`

## Example Workflow
`bit switch example-branch`
Branch does not exist. Do you want to create it? Y/n

yes

Switched to a new branch 'example-branch'

[Makes some changes]

`bit save "add important feature"`

[fix an error for important feature]

`bit save`

[push changes to origin]

`bit sync`

[two days later confirm your branch is in sync with origin]

`bit sync`




## Features

- Automatic fetching & fast forwarding to keep your branches up to date and prevent merge conflicts
- Simplify your entire rebase workflow with a single command `bit sync` 
- Automatic suggestions at your fingertips 
- `bit` is **fully compatible** with `git`. All features of git are available if need be.  

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


## Principles 

1. Think in the age of the cloud
1. Embed the spirit of modern day workflows
1. Favor simplicity over complexity 
1. Bit should have happy defaults
1. Bit must be fully compatible with Git

## Inspiration

Thanks to [Gitless](https://gitless.com/), [git-extras](https://github.com/tj/git-extras), researchers in the field and of course the developers of `git` itself! Also, thanks to [go-prompt](https://github.com/c-bata/go-prompt) for the interactive prompt library

## News 
- https://www.tldrnewsletter.com/archives/20201007
- https://www.reddit.com/r/golang/comments/j5wggn/bit_an_experimental_git_cli_with_a_powerful/
- https://b.hatena.ne.jp/entry/s/github.com/chriswalz/bit
- https://news.hada.io/topic?id=2990

## Changelog 
v0.5

- [X] `bit switch`, `bit co`, `bit checkout` will show prompt 
- [X] fix bit tab completion (bitcomplete)
- [X] fix edge case where there is an invalid ref
- [X] various minor fixes
- [X] Deployment: automatic versioning 

v0.4

- [X] multiline support with Go Survey Library
- [X] color mitigation to have roughly similar colors across OSs 
- [X] fix README go get installation instructions
- [X] QOL improvements for `bit sync`
 
## How to uninstall
*go binaries are self-contained so uninstalling simply requires deleting the binary(ies)*

```
rm `which bit`
rm `which bitcomplete` 
```
