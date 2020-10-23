package cmd

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/tj/go-update"
	"github.com/tj/go-update/stores/github"
	"os/exec"
	"path/filepath"
	"runtime"
)

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Updates bit to the latest or specified version",
	Long:  `bit update
bit update v0.7.4 (note: v is required)`,
	Run: func(cmd *cobra.Command, args []string) {
		targetVersion := ""
		if len(args) == 1 {
			targetVersion = args[0][1:]
		}

		currentVersion := GetVersion()

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
		release, err := getLatestOrSpecified(p, targetVersion)
		if err != nil {
			fmt.Println(errors.Wrap(err, "fetching latest or specified release").Error())
			return
		}

		// no updates
		if release == nil || currentVersion == release.Version {
			fmt.Println("No updates available, update to date!")
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
		dst := filepath.Dir(path)

		// install it
		if err := p.InstallTo(tarball, dst); err != nil {
			fmt.Println(errors.Wrap(err, "installing"))
			return
		}

		fmt.Printf("Updated bit %s to %s in %s", currentVersion, release.Version, dst)

	},
	Args: cobra.MaximumNArgs(1),
}

func init() {
	ShellCmd.AddCommand(updateCmd)
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
