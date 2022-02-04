package notifysend

import (
	"os/exec"
	"github.com/rs/zerolog/log"
	"bytes"
)

type NotificationCmd struct{
	Expire: int
	TitleText: string
	BodyText: string
}

func New(title: string, body: string) *NotificationCmd {
	return &NotificationCmd{
		Expire: 5000,
		TitleText: title,
		BodyText: body
	}
}

func (n *NotificationCmd) Send() {
	cmd := exec.Command("notify-send", "-t", n.Expire, n.TitleText, n.BodyText)
	err := cmd.Run()
	if err != nil {
		panic(err)
	}
}