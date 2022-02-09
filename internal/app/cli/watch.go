package cli

import (
	"os/exec"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/joshburnsxyz/puck/pkg/notifysend"
)

// WatchCmd() implements the watch command.
func WatchCmd() *cobra.Command {
	return &cobra.Command{
		Use:     "watch",
		Aliases: []string{"w"},
		Args: cobra.ArbitraryArgs,
		Short:   "run a given command and notify when its done",
		Run: func(cmd *cobra.Command, args []string) {
			cmdBinPath := args[0]
			_, args = args[0], args[1:]
			cmdN := exec.Command(args...)
			err := cmdN.Run()
			if err != nil {
				n := notifysend.New("ERROR", err)
				n.Send()
				panic(err)
			} else {
				n := notifysend.New("SUCCESS", fmt.Sprintf("%v has finished", args[0]))
				n.Send()
			}
		},
	}
}
