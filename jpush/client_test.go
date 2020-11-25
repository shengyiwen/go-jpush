package jpush

import "testing"

func TestClient_SendNotification(t *testing.T) {
	client := NewClient("appKey", "masterSecret")

	request := NewRequest()

	platform := NewPlatform().Android().Ios()
	request.SetPlatform(platform)

	audience := NewAudience().SetRegistrationIds([]string{"140fe1da9e2573338e7"})
	request.SetAudienceAll(audience)

	notification := NewNotification().SetAndroid(&AndroidNotification{
		Alert: "测试一下看看能否正常发送",
		Title: "测试",
	})
	request.SetNotification(notification)

	result, err := client.Send(request)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("是否发送成功", result.IsSuccess())
	t.Log(result)
}
