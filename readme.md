### go-jpush介绍

是由Go语言对极光推送客户端的实现，目前仅实现了推送功能

### go-jpush的使用

使用前进行下载到本地：

    go get github.com/shengyiwen/go-jpush

使用案例：

    import "testing"
    
    func TestClient_SendNotification(t *testing.T) {
        // todo 更换appKey 和 secretKey
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

