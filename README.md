
<p align="center">
<img
    src="https://user-images.githubusercontent.com/6971318/94226426-96309480-fec5-11ea-929f-34551c5356dd.png"
    width="800px" border="0" alt="croc">
<br>
<a href="https://github.com/chriswalz/bit/tags"><img src="https://img.shields.io/badge/version-v0.3.4-brightgreen.svg?style=flat-square" alt="Version"></a>
</p>

`bit` is a modernized CLI built on top of git that provides command and --flag suggestions automatic fetch and fast-forwarding along with other niceties:

- command and **flag suggestions** to help you navigate the plethora of options git provides you
- autocompletion for files and branch names when using `bit add` or `bit checkout`
- automatic fetch and **branch fast-forwarding** reducing the likelihood of merge conflicts 
- suggestions **work with git aliases**
- new commands like `bit sync` that vastly simplify your workflow 
- commands from git-extras such as *git-release* & *git-info*
- fully compatible with git allowing you to fallback to git if need be.  

## Installation


`curl -sf https://gobinaries.com/chriswalz/bit | sh`

Verify installation with:

`bit`

Dependencies: Git

## Usage 

Create a new commit

`bit save [commit message]`

Save your changes to the current branch [amends current commit]

`bit save` 

Synchronize yours changes to origin branch 

`bit sync`

Switch branches

`bit switch [branch-name]`

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
- Every branch is a completely independent line. Changes in your working directory are saved and in the branch that you switch saved changes are retrieved.
- Simplify your entire rebase workflow with a single command `bit sync [branch-name]` 
- Automatic suggestions at your fingertips 
- `bit` is **fully compatible** with `git`. All features of git are available if need be.  


## Principles 

1. Think in the age of the cloud
1. Embed the spirit of modern day workflows
1. Favor simplicity over complexity 
1. Bit should have happy defaults
1. Bit must be fully compatible with Git

## Inspiration

Thanks to [Gitless](https://gitless.com/), [git-extras](https://github.com/tj/git-extras), researchers in the field and of course the developers of git itself! 
