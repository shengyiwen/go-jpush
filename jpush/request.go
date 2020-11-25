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

func (builder *Request) SetCallback(callback *Callback) *Request {
	builder.Callback = callback
	return builder
}

func (builder *Request) SetOption(option *Option) *Request {
	builder.Option = option
	return builder
}

func (builder *Request) SetMessage(message *Message) *Request {
	builder.Message = message
	return builder
}

func (builder *Request) SetNotification(notification *Notification) *Request {
	builder.Notification = notification
	return builder
}

func (builder *Request) SetAudience(audience *Audience) *Request {
	builder.Audience = audience.ToJsonElement()
	return builder
}

func (builder *Request) SetCid(cid string) *Request {
	builder.Cid = cid
	return builder
}

func (builder *Request) SetPlatform(platform *Platform) *Request {
	builder.Platform = platform.ToJsonElement()
	return builder
}
