package jpush

// 通知，是一条推送的实体对象之一，另一个是消息
// 会作为"通知"推送到客户端
type Notification struct {

	// 通知到内容在各个平台上，都可能只有这一个基础的属性Alert
	// 这个Alert是个快捷定义，在各个平台通知里也有Alert定义，且平台通知里的优先级要高
	Alert string `json:"alert,omitempty"`

	// Ios的通知内容
	Ios *IosNotification `json:"ios,omitempty"`

	// Android的通知内容
	Android *AndroidNotification `json:"android,omitempty"`

	Voip *map[string]interface{} `json:"voip,omitempty"`
}

func NewNotification() *Notification {
	return &Notification{}
}

func (n *Notification) SetAlert(alert string) *Notification {
	n.Alert = alert
	return n
}

func (n *Notification) SetIos(ios *IosNotification) *Notification {
	n.Ios = ios
	return n
}

func (n *Notification) SetAndroid(android *AndroidNotification) *Notification {
	n.Android = android
	return n
}

func (n *Notification) SetVopi(vopi *map[string]interface{}) *Notification {
	n.Voip = vopi
	return n
}

func (n *Notification) ToJsonElement() interface{} {
	return n
}

// Android平台的通知
type AndroidNotification struct {

	// 通知内容，如果制定了，覆盖Notification中的Alert
	// 内容可以为空，为空表示不展示到通知栏
	Alert string `json:"alert,omitempty"`

	// 通知标题，如果指定了，展示App名称的地方，会替换成这个字段
	Title string `json:"title,omitempty"`

	// 通知样式Id，AndroidSdk可以设置通知栏样式，并通过这里的Id来指定使用哪套
	BuilderId string `json:"builder_id,omitempty"`

	// android通知channelId
	ChannelId string `json:"channel_id,omitempty"`

	// 通知栏展示优先级，默认是0，范围[-2, 2]
	Priority int `json:"priority,omitempty"`

	// 通知栏条目过滤或排序
	Category string `json:"category,omitempty"`

	// 通知栏样式类型，默认为0，可以选择0、1、2、3
	// 用来指定选择哪种通知栏样式，其他值无效，bigText = 1, Inbox = 2, bigPicture = 3
	Style int `json:"style,omitempty"`

	// 大文本通知栏样式
	BigText string `json:"big_text,omitempty"`

	// 文本条目通知栏样式
	Inbox string `json:"inbox,omitempty"`

	// 大图片通知栏样式
	BigPicPath string `json:"big_pic_path,omitempty"`

	// 扩展字段
	Extras map[string]interface{} `json:"extras,omitempty"`

	// 通知栏大图标
	LargeIcon string `json:"large_icon,omitempty"`

	// 指定跳转页面
	Intent map[string]interface{} `json:"intent,omitempty"`

	// 指定跳转页面
	UrlActivity string `json:"uri_activity,omitempty"`

	// 指定跳转页面
	UriAction string `json:"uri_action,omitempty"`

	// 角标数字，取值范围1-99
	BadgeAddNum int `json:"badge_add_num,omitempty"`

	// 桌面图标对应的应用入口Activity类
	BadgeClass string `json:"badge_class,omitempty"`

	// 填写Android工程中/res/raw/路径下铃声文件名称，无需文件名后缀
	Sound string `json:"sound,omitempty"`

	// 定时展示开始时间（yyyy-MM-dd HH:mm:ss）
	ShowBeginTime string `json:"show_begin_time,omitempty"`

	// 定时展示结束时间（yyyy-MM-dd HH:mm:ss）
	ShowEndTime string `json:"show_end_time,omitempty"`

	// APP在前台，通知是否展示
	DisplayForeground string `json:"display_foreground,omitempty"`
}

type IosNotification struct {

	// 通知内容，如果制定了，覆盖Notification中的Alert
	// 内容可以为空，为空表示不展示到通知栏
	Alert string `json:"alert,omitempty"`

	// 通知提示声音或警告通知
	Sound interface{} `json:"sound,omitempty"`

	// 应用角标
	Badge int `json:"badge,omitempty"`

	// 推送唤醒
	ContentAvailable bool `json:"content-available,omitempty"`

	// 通知扩展
	MutableContent bool `json:"mutable-content,omitempty"`

	// IOS 8才支持
	Category string `json:"category,omitempty"`

	// 附加字段
	Extra map[string]interface{} `json:"extras,omitempty"`

	// 通知分组
	ThreadId string `json:"thread-id,omitempty"`
}
