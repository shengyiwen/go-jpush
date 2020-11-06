package jpush

type Request struct {
	Cid string `json:"cid,omitempty"`

	Platform interface{} `json:"platform,omitempty"`

	Audience interface{} `json:"audience,omitempty"`

	Notification *Notification `json:"notification,omitempty"`

	Message *Message `json:"message,omitempty"`

	Option *Option `json:"options,omitempty"`

	Callback *Callback `json:"callback,omitempty"`
}

func NewRequest() *Request {
	return &Request{}
}

func (builder *Request) SetCallback(callback *Callback) {
	builder.Callback = callback
}

func (builder *Request) SetOption(option *Option) {
	builder.Option = option
}

func (builder *Request) SetMessage(message *Message) {
	builder.Message = message
}

func (builder *Request) SetNotification(notification *Notification) {
	builder.Notification = notification
}

func (builder *Request) SetAudienceAll(audience *Audience) {
	builder.Audience = audience.ToJsonElement()
}

func (builder *Request) SetCid(cid string) {
	builder.Cid = cid
}

func (builder *Request) SetPlatform(platform *Platform) {
	builder.Platform = platform.ToJsonElement()
}
