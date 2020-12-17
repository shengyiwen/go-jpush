package jpush

type Request struct {
	cid          string
	platform     *Platform
	audience     *Audience
	notification *Notification
	message      *Message
	option       *Option
	callback     *Callback
}

func (builder *Request) ToJsonElement() interface{} {
	var jsonElement = make(map[string]interface{})

	AssertNotNull(builder.cid, func() {
		jsonElement["cid"] = builder.cid
	})

	AssertNotNull(builder.platform, func() {
		jsonElement["platform"] = builder.platform.ToJsonElement()
	})

	AssertNotNull(builder.audience, func() {
		jsonElement["audience"] = builder.audience.ToJsonElement()
	})

	AssertNotNull(builder.notification, func() {
		jsonElement["notification"] = builder.notification.ToJsonElement()
	})

	AssertNotNull(builder.message, func() {
		jsonElement["message"] = builder.message.ToJsonElement()
	})

	AssertNotNull(builder.option, func() {
		jsonElement["option"] = builder.option.ToJsonElement()
	})

	AssertNotNull(builder.callback, func() {
		jsonElement["callback"] = builder.callback.ToJsonElement()
	})

	return jsonElement
}

func NewRequest() *Request {
	return &Request{}
}

func (builder *Request) Callback(callback *Callback) *Request {
	builder.callback = callback
	return builder
}

func (builder *Request) Option(option *Option) *Request {
	builder.option = option
	return builder
}

func (builder *Request) Message(message *Message) *Request {
	builder.message = message
	return builder
}

func (builder *Request) Notification(notification *Notification) *Request {
	builder.notification = notification
	return builder
}

func (builder *Request) Audience(audience *Audience) *Request {
	builder.audience = audience
	return builder
}

func (builder *Request) Cid(cid string) *Request {
	builder.cid = cid
	return builder
}

func (builder *Request) Platform(platform *Platform) *Request {
	builder.platform = platform
	return builder
}
