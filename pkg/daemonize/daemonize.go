package daemonize

import (
	"log"
	"os"
	"os/exec"
	"errors"
)

func Cmd(inp exec.Cmd) errors.Error {
	inp.Stdout = os.Stdout // note to future self: redirect output to /dev/null or possibly a logger
	err := inp.Start()
	if err != nil {
		return err
	}
	log.Printf("Background Process %d started", inp.Process.Pid)
	return nil
}
