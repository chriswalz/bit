# bit

Bit is a simple version control system built on top of git. Bit is super easy to learn and will vastly simplify your development workflow. 

At times you may still want/need to use git. When the times arrive 

## Installation


curl -sf https://gobinaries.com/chriswalz/bit | sh

note: git must be installed 

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


## Principles 

Think in the age of the cloud
Development workflows have changed in the 15 years since Gits creation 
Favor simplicity over complexity 
Bit should have happy defaults
Bit must be fully compatible with Git