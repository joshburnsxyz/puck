package notifysend

import (
	"testing"
)

func TestNew(t *testing.T) {
	want := &NotificationCmd{
		Expire: 5000,
		TitleText: "Hello",
		BodyText: "World"
	}

	cmd,err := New("Hello", "World")
	if want.TitleText != cmd.TitleText {
		t.Fatalf("New() Got %v wanted %v", cmd.TitleText, want.TitleText)
	}
}