package cli

import (
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

// ExecuteRootCmd prepares all CLI commands
func ExecuteRootCmd() {
	c := cobra.Command{
		Use:   "puck",
		Short: "Notifications for commands",
		Version: "1.0",
	}

	c.AddCommand(WatchCmd())
	c.AddCommand(TestCmd())

	if err := c.Execute(); err != nil {
		log.Fatal().Err(err)
	}
}
