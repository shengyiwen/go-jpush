package jpush

// 应用内消息，或称为:自定义消息、透传消息
// 此部分内容不会展示到通知栏上，JPush SDK 收到消息内容后透传给 App。需要 App 自行处理
// iOS 在推送应用内消息通道（非 APNS）获取此部分内容，即需 App 处于前台。Windows Phone 暂时不支持应用内消息
type Message struct {

	// 消息内容本身
	Content string `json:"msg_content,omitempty"`

	// 消息标题
	Title string `json:"title,omitempty"`

	// 消息内容类型
	ContentType string `json:"content_type,omitempty"`

	// JSON 格式的可选参数
	Extras map[string]interface{} `json:"extras,omitempty"`
}

func (m *Message) ToJsonElement() interface{} {
	return m
}
