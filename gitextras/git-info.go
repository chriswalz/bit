package gitextras

// GitInfo is the git-info bash script from git-extras
const GitInfo = `#!/usr/bin/env bash
GREEN="$(tput setaf 2)"
NORMAL="$(tput sgr0)"
if [ "$1" = "--color" ] || [ "$2" = "--color" ] || \
     [ "$1" = "-c" ] || [ "$2" = "-c" ] ; then
  COLOR_TITLE="$GREEN"
else
  COLOR_TITLE="$NORMAL"
fi

HIDE_CONFIG=
if [ "$1" != "--no-config" ] && [ "$2" != "--no-config" ]; then
  HIDE_CONFIG=1
fi

get_config() {
  cmd_get_config="$(git config --get-all git-extras.info.config-grep)"
  if [ -z "$cmd_get_config" ]; then
    git config --list
  else
    eval "$cmd_get_config"
  fi
}

most_recent_commit() {
  cmd_get_log="$(git config --get-all git-extras.info.log)"
  if [ -z "$cmd_get_log" ]; then
    git log --max-count=1 --pretty=short
  else
    eval "$cmd_get_log"
  fi
}

submodules() {
  # short sha1
  git submodule status | sed 's/\([^a-f\d]\{0,2\}\)\([a-f\d]\{7\}\)\([a-f\d]\{33\}\)\(.*\)/\1\2\4/'
}

local_branches() {
  git branch
}

remote_branches() {
  git branch -r
}

remote_urls() {
  git remote -v
}

echon() {
  echo "$@"
  echo
}

echo
echon "${COLOR_TITLE}## Remote URLs:${NORMAL}"
echon "$(remote_urls)"

echon "${COLOR_TITLE}## Remote Branches:${NORMAL}"
echon "$(remote_branches)"

echon "${COLOR_TITLE}## Local Branches:${NORMAL}"
echon "$(local_branches)"

SUBMODULES_LOG=$(submodules)
if [ ! -z "$SUBMODULES_LOG" ]; then
  echon "${COLOR_TITLE}## Submodule(s):${NORMAL}"
  echon "$SUBMODULES_LOG"
fi

echon "${COLOR_TITLE}## Most Recent Commit:${NORMAL}"
echon "$(most_recent_commit)"

if [ ! -z "$HIDE_CONFIG" ]; then
  echon "${COLOR_TITLE}## Configuration (.git/config):${NORMAL}"
  echon "$(get_config)"
fi
`
