package jpush

// 可选参数
type Option struct {
	SendNo          int    `json:"sendno,omitempty"`            // 推送序号
	TimeToLive      int    `json:"time_to_live,omitempty"`      // 离线消息保留时长(秒)
	OverrideMsgId   int64  `json:"override_msg_id,omitempty"`   // 要覆盖的消息 ID
	ApnsProduction  bool   `json:"apns_production,omitempty"`   // 仅对IOS有效，是否是生产环境
	ApnsCollapseId  string `json:"apns_collapse_id,omitempty"`  // 更新IOS通知的标识符
	BigPushDuration int    `json:"big_push_duration,omitempty"` // 定速推送时长(分钟)

	// 推送请求下发通道，仅对厂商VIP用户有效
	// ThirdPartyChannel map[string]interface{} `json:"third_party_channel"`

}

func (o *Option) ToJsonElement() interface{} {
	return o
}
