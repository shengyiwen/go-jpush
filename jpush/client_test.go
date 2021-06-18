package jpush

import "testing"

func TestClient_SendNotification(t *testing.T) {

	client := NewClient("806a24650079d40b9f0cea61", "71e9a6e9fb046e92cca3289d", true)

	platform := NewPlatform().Android().Ios()

	audience := NewAudience().AddRegistrationId("1507bfd3f7afd7426fc")

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
