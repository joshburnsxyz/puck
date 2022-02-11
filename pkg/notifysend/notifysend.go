package notifysend

import (
	"os/exec"
	"fmt"
)

type NotificationCmd struct {
	Expire    int
	Urgency   string
	TitleText string
	BodyText  string
}

func New(title string, body string) *NotificationCmd {
	return &NotificationCmd{
		Expire:    5000,
		Urgency:   "normal",
		TitleText: title,
		BodyText:  body,
	}
}

func (n *NotificationCmd) Send() {
	cmd := exec.Command("notify-send", "-t", fmt.Sprintf("%v",n.Expire), "-u", fmt.Sprintf("%s",n.Urgency), fmt.Sprintf("%s",n.TitleText), fmt.Sprintf("%s",n.BodyText))
	err := cmd.Run()
	if err != nil {
		fmt.Println(err)
	}
}
