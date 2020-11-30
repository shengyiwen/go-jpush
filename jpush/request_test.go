package jpush

import (
	"encoding/json"
	"testing"
)

func TestRequestToJsonElement(t *testing.T) {
	request := NewRequest().
		Audience(NewAudience().AddRegistrationId("123456")).
		Platform(NewPlatform().All()).
		Notification(NewNotification().Android(&AndroidNotification{Alert: "test", Title: "test"}))

	bytes, _ := json.Marshal(request.ToJsonElement())
	t.Log(string(bytes))

}
