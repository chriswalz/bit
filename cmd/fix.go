package cmd

import (
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

var fixCmd = &cobra.Command{
	Use:   "fix sub-command",
	Short: "For all the times you did something you really wish you didn't",
	Args:  cobra.NoArgs,
}

var undo_commitCmd = &cobra.Command{
	Use:   "undo-commit",
	Short: "soft undos last commit if not pushed already",
	Run: func(cmd *cobra.Command, args []string) {
		if IsAheadOfCurrent() {
			err := execCommand("git", "reset", "--soft", "HEAD~1").Run()
			if err != nil {
				log.Debug().Err(err).Send()
			}
		}
	},
	Args: cobra.NoArgs,
}

func init() {
	BitCmd.AddCommand(fixCmd)
	fixCmd.AddCommand(undo_commitCmd)
}
