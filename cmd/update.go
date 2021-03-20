package cmd

import (
	"fmt"
	exec "golang.org/x/sys/execabs"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"github.com/tj/go-update"
	"github.com/tj/go-update/stores/github"
)

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Updates bit to the latest or specified version",
	Long: `bit update
bit update v0.7.4 (note: v is required)`,
	Run: func(cmd *cobra.Command, args []string) {
		targetVersion := ""
		if len(args) == 1 {
			targetVersion = args[0]
		}
		if !strings.HasPrefix(targetVersion, "v") {
			targetVersion = "v" + targetVersion
		}
		currentVersion := GetVersion()
		if !strings.HasPrefix(currentVersion, "v") {
			currentVersion = "v" + currentVersion
		}

		log.Debug().Msg(currentVersion + " -> " + targetVersion)

		// open-source edition
		p := &update.Manager{
			Command: "bit",
			Store: &github.Store{
				Owner:   "chriswalz",
				Repo:    "bit",
				Version: currentVersion[1:],
			},
		}

		// fetch latest or specified release
		release, err := getLatestOrSpecified(p, targetVersion[1:])
		if err != nil {
			fmt.Println(errors.Wrap(err, "fetching latest or specified release").Error())
			return
		}

		// no updates
		if release == nil || currentVersion == release.Version {
			fmt.Println("No updates available, you're up to date!")
			return
		}

		// find the tarball for this system
		a := release.FindTarball(runtime.GOOS, runtime.GOARCH)
		if a == nil {
			fmt.Println(fmt.Errorf("failed to find a binary for %s %s", runtime.GOOS, runtime.GOARCH))
			return
		}

		// download tarball to a tmp dir
		tarball, err := a.Download()
		if err != nil {
			fmt.Println(errors.Wrap(err, "downloading tarball"))
			return
		}

		// determine path
		path, err := exec.LookPath("bit")
		if err != nil {
			fmt.Println(errors.Wrap(err, "looking up executable path"))
			return
		}

		// if path is a symlink - get resolved path
		fi, err := os.Lstat(path)
		if err == nil && fi.Mode()&os.ModeSymlink == os.ModeSymlink {
			// Bit path is a symlink
			fmt.Println("bit is symlinked. If you used homebrew try:\nbrew upgrade bit-git")
			// path = resolvedSymlink
			return
		}
		log.Debug().Msg("bit is not symlinked")

		dst := filepath.Dir(path)

		// install it
		if err := p.InstallTo(tarball, dst); err != nil {
			fmt.Println(errors.Wrap(err, "installing"))
			return
		}

		fmt.Println("Bit is supported through donations. Consider donating here: ❤ https://github.com/sponsors/chriswalz ")
		fmt.Printf("Updated bit %s to %s in %s\n", currentVersion, release.Version, dst)
	},
	Args: cobra.MaximumNArgs(1),
}

func init() {
	BitCmd.AddCommand(updateCmd)
}

// getLatestOrSpecified returns the latest or specified release.
func getLatestOrSpecified(s update.Store, version string) (*update.Release, error) {
	if version == "" {
		return getLatest(s)
	}

	return s.GetRelease(version)
}

// getLatest returns the latest release, error, or nil when there is none.
func getLatest(s update.Store) (*update.Release, error) {
	releases, err := s.LatestReleases()
	if err != nil {
		return nil, errors.Wrap(err, "fetching releases")
	}

	if len(releases) == 0 {
		return nil, nil
	}

	return releases[0], nil
}
