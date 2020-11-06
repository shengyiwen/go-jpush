package jpush

// 回调参数，仅对VIP用户提供
// 可以指定 callback 参数，方便用户临时变更回调 URL 或者回调带上其自定义参数，满足其日常业务需求
type Callback struct {
	Url string `json:"url,omitempty"`

	Params map[string]interface{} `json:"params,omitempty"`

	Type string `json:"type,omitempty"`
}
