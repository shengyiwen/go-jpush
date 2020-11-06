package jpush

// 推送平台，如果推送目标为IOS平台，推送Notification时需要在options中
// 通过ans_production字段决定环境，True表示生产环境，False表示推送开发环境

type DeviceType string

const (
	// 安卓平台
	Android DeviceType = "android"

	// IOS平台
	Ios DeviceType = "ios"

	// WinPhone平台
	WinPhone DeviceType = "winphone"
)

type Platform struct {
	all bool

	deviceType []DeviceType
}

func NewPlatform() *Platform {
	return &Platform{}
}

func (p *Platform) All() *Platform {
	p.all = true
	return p
}

func (p *Platform) Android() *Platform {
	p.deviceType = append(p.deviceType, Android)
	return p
}

func (p *Platform) Ios() *Platform {
	p.deviceType = append(p.deviceType, Ios)
	return p
}

func (p *Platform) WinPhone() *Platform {
	p.deviceType = append(p.deviceType, WinPhone)
	return p
}

func (p *Platform) AddDeviceType(deviceType DeviceType) *Platform {
	p.deviceType = append(p.deviceType, deviceType)
	return p
}

func (p *Platform) ToJsonElement() interface{} {
	if p.all {
		return All
	}
	return &p.deviceType
}
