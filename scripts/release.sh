set -e

bit release bump;
goreleaser --rm-dist;