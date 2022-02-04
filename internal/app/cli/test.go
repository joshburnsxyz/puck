package cli

import (
	"github.com/spf13/cobra"
	"github.com/joshburnsxyz/puck/pkg/notifysend"
)

// TestCmd() implements the test command.
func TestCmd() *cobra.Command {
	return &cobra.Command{
		Use:     "test",
		Aliases: []string{"t"},
		Args: cobra.NoArgs,
		Short:   "display a test notification",
		Run: func(cmd *cobra.Command, args []string) {
				n := notifysend.New("TEST", "this is a test notification")
				n.Send()
		},
	}
}