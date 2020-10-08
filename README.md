[![Tweet](https://img.shields.io/twitter/url/http/shields.io.svg?style=social)](https://twitter.com/intent/tweet?text=Bit%20-%20a%20modern%20git%20cli%20&url=https://github.com/chriswalz/bit&hashtags=bit,git,cli,developers)
<p align="center">
<img
    src="https://user-images.githubusercontent.com/6971318/95147815-cd1d6a80-074f-11eb-8265-56466ac628f8.gif"
    width="600px" border="0" alt="bit">
<br>
<a href="https://github.com/chriswalz/bit/tags"><img src="https://img.shields.io/badge/version-v0.4.10-brightgreen.svg?style=flat-square" alt="Version"></a>
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

## Installation

### Curl

```shell script
curl -sf https://gobinaries.com/chriswalz/bit | sh;
curl -sf https://gobinaries.com/chriswalz/bit/bitcomplete | sh && echo y | COMP_INSTALL=1 bitcomplete
```


### Go Get 
Caveats: GOPATH and GOBIN need to be set. Verify with `go env`. If they're not set add this to your .bashrc or .bash_profile etc. AND reset terminal
```
export GOPATH=$HOME/go
export GOBIN=$(go env GOPATH)/bin
```


```shell script
GO111MODULE=on go get github.com/chriswalz/bit@latest;
GO111MODULE=on go get github.com/chriswalz/bit/bitcomplete@latest;

COMP_INSTALL=1 bitcomplete;
bit
```


*Note*: Tab completion only works on Mac, Linux, Ubuntu etc.

Verify installation with:

`bit`

Dependencies: Git

Platform Support:
- iTerm2 (macOS)
- Terminal.app (macOS)
- Command Prompt (Windows)
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
