package cmd

import (
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

var undo_commitCmd = &cobra.Command{
	Use:   "undo-commit",
	Short: "Undo your previous commit if it is not yet pushed to repository",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		if IsAheadOfCurrent() {
			err := execCommand("git", "reset", "--soft", "HEAD~1").Run()
			if err != nil {
				log.Debug().Err(err).Send()
			}
		}
	},
}

func init() {
	BitCmd.AddCommand(undo_commitCmd)
}
