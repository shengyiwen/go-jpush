package jpush

import "testing"

func TestClient_SendNotification(t *testing.T) {
	client := NewClient("x", "x")

	request := Builder()

	platform := NewPlatform().Android()

	request.SetPlatform(platform)

	audience := NewAudience()
	audience.SetRegistrationIds([]string{"13065ffa4e6342a6f5c"})

	notification := &Notification{}
	notification.Android = &AndroidNotification{
		Alert: "测试一下看看能否正常发送",
		Title: "测试",
	}

	request.SetAudienceAll(audience)
	request.SetNotification(notification)

	result, err := client.Send(request)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("是否发送成功", result.IsSuccess())
	t.Log(result)
}
