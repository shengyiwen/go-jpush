package jpush

import (
	"encoding/json"
	"net/http"
)

// 所有平台
const All string = "all"

// 推送地址
const PushUrl string = "https://api.jpush.cn/v3/push"

// 推送Model对象，即所有推送内容的配置都能转为Json元素
type PushModel interface {

	// 转为Json元素
	ToJsonElement() interface{}
}

// 消息推送结果
type PushResult struct {

	// 消息Id
	MsgId string `json:"msg_id"`

	// 推送编号
	SendNo string `json:"sendno"`

	// 状态码
	StatusCode int `json:"statusCode"`

	// 错误消息
	Error *PushError `json:"error"`
}

// 是否调用成功
func (res *PushResult) IsSuccess() bool {
	return res.StatusCode == http.StatusOK && res.Error == nil
}

// String处理
func (res *PushResult) String() string {
	marshal, _ := json.Marshal(res)
	return string(marshal)
}

// 推送错误结果
type PushError struct {
	Message string `json:"message"`

	Code int `json:"code"`
}

// 从推送返回内容解析推送结果
func FromResponse(statusCode int, body string) (*PushResult, error) {
	result := &PushResult{}
	result.StatusCode = statusCode
	err := json.Unmarshal([]byte(body), result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
