#!/usr/bin/env bash

GREEN="$(tput setaf 2)"
NORMAL="$(tput sgr0)"
COLOR_TITLE="$NORMAL"

echon() {
  echo -e "$@\n"  
}

cechon() {
    echo -e "${COLOR_TITLE}## $@${NORMAL}\n"
}

# me, the absolute path of the script
me=$(realpath --strip ${BASH_SOURCE:-0})
# here, the abolute path of the directory where the script lives.
here=$(dirname ${me})
# verb, the action. git-update.sh -> update. This script will expect to call update().
verb=$(basename ${me##*-} .sh)

function update() {
  set -euo pipefail
  IFS=$'\n\t'

  # I assume there's some way to find the version of bit before downloading it.
  # I just invented one.
  remote_version=$(2>/dev/null curl -sSf https://gobinaries.com/chriswalz/bit.version.txt) || true
  installed_version=$(${here}/bit --version|head -1|awk '{print $3;}')

  # No need to download if you have the latest version.
  if [[ "${remote_version}" == "${installed_version}" ]] ; then
      >&2 cechon "${here}/bit already at version '${remote_version}'"
      exit 0
  fi

  # Otherwise, create a temporary place to install bit* and check the version again.
  export PREFIX=/tmp/bit-${verb}/$$
  mkdir -p ${PREFIX}
  # Remove the download when done.
  trap "rm -rf /tmp/bit-${verb}" EXIT
  
  # Get the latest.
  curl -sf https://gobinaries.com/chriswalz/bit | sh
  curl -sf https://gobinaries.com/chriswalz/bit/bitcomplete | sh
  new_version=$(${PREFIX}/bit --version|head -1|awk '{print $3;}')

  # Check the version one more time. Punt if you have the latest.
  if [[ "${new_version}" == "${installed_version}" ]] ; then
      >&2 cechon "${here}/bit already at version '${installed_version}'"
  else
      # New bits! Archive the installed ones with their version id...
      mv ${here}/bit{,${installed_version}}
      mv ${here}/bitcomplete{,${installed_version}}
      # ... and copy the newer versions.
      mv ${PREFIX}/bit* ${here}
      echo y | COMP_INSTALL=1 ${here}/bitcomplete
      >&2 cechon "Installed ${here}/bit* version '${new_version}' in '${here}'"
  fi
}

${verb} $*
