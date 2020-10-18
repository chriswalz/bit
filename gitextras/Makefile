# Makefile to generate git-<action>.go files from corresponding git-<action>.sh files.
# Makes it easier to modify and test scripts and then incorporate them on a build.

# Every git-*.sh file will generate a git-*.go file with the correct conventions.
EXTRAS := $(patsubst %.sh, %.go, $(wildcard git-*.sh))
git-%.go : git-%.sh
	echo -n "// generated from $<; DO NOT EDIT\npackage gitextras; const Git$* = \`" > $@
	cat $< >> $@
	echo "\`" >> $@

.PHONY: all
# Wrap git-<action>.sh in git-<action>.go as needed.
all : $(EXTRAS) .gitignore

# Tell git to ignore the generated files.
.gitignore : $(shell grep -l -e '^// generated' git-*.go)
	echo $< | tr ' ' '\n' | sort > $@
