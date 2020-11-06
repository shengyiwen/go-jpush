package jpush

import (
	"encoding/json"
	"net/http"
)

type PushModel interface {
	ToJsonElement() interface{}
}

// 所有平台
const All string = "all"

// 推送地址
const PushUrl string = "https://api.jpush.cn/v3/push"

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

type PushError struct {
	Message string `json:"message"`

	Code int `json:"code"`
}

func FromResponse(statusCode int, body string) (*PushResult, error) {
	result := &PushResult{}
	result.StatusCode = statusCode
	err := json.Unmarshal([]byte(body), result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (res *PushResult) IsSuccess() bool {
	return res.StatusCode == http.StatusOK && res.Error == nil
}

func (res *PushResult) String() string {
	marshal, _ := json.Marshal(res)
	return string(marshal)
}
