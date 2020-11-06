### go-jpush介绍

是由Go语言对极光推送客户端的实现，目前仅实现了推送功能

### go-jpush的使用

使用前进行下载到本地：

    go get github.com/shengyiwen/go-jpush

使用案例：

    import "testing"
    
    func TestClient_SendNotification(t *testing.T) {
        // todo 更换appKey 和 secretKey
        client := NewClient("appKey", "secretKey")
    
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

