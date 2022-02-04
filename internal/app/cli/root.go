package cli

import (
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

// ExecuteRootCmd prepares all CLI commands
func ExecuteRootCmd() {
	c := cobra.Command{
		Use: "puckd",
		Short: "Notifications for commands",
	}

	//c.AddCommand(NewCmd())

	if err := c.Execute(); err != nil {
		log.Fatal().Err(err)
	}
}
