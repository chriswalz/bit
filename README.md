# bit

Bit is an experimental simple version control system built on top of git. Bit is super easy to learn and will vastly simplify your development workflow. 

At times, you may still want/need to use a git command. You can simply run the git command or run it via bit (e.g. bit rebase -i SHA-HASH) 

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
- `bit` is **fully compatible** with `git`. All features of git are available if need be.  


## Principles 

1. Think in the age of the cloud
1. Embed the spirit of modern day workflows
1. Favor simplicity over complexity 
1. Bit should have happy defaults
1. Bit must be fully compatible with Git