package jpush

import "testing"

func TestClient_SendNotification(t *testing.T) {

	client := NewClient("x", "x", true)

	platform := NewPlatform().Android().Ios()

	audience := NewAudience().AddRegistrationId("140fe1da9e2573338e7")

	notification := NewNotification().
		Android(&AndroidNotification{Alert: "测试一下看看能否正常发送", Title: "测试"}).
		Ios(&IosNotification{Alert: "测试一下看看能够正常发送"})

	request := NewRequest().
		Platform(platform).
		Audience(audience).
		Notification(notification)

	result, err := client.Send(request)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("是否发送成功", result.IsSuccess())
	t.Log(result)
}
