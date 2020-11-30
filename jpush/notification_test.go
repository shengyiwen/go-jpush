package jpush

import (
	"encoding/json"
	"testing"
)

func TestNotificationToJsonElement(t *testing.T) {
	notification := NewNotification().
		Alert("alert").
		Android(&AndroidNotification{Title: "hello"}).
		Ios(&IosNotification{Alert: "good"})
	bytes, _ := json.Marshal(notification.ToJsonElement())
	t.Log(string(bytes))
}
