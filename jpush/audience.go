package jpush

// 推送目标，推送设备对象，表示一条推送可以推送到哪些设备列表
// 如果是推送给所有设备，则直接填写all
type Audience struct {

	// 如果all为true的话，则剩下标签无效
	all bool

	// 标签，多组标签的关系是Or的关系，即取并集
	Tag []string `json:"tag,omitempty"`

	// 标签，多组标签的关系是And关系，取交集
	TagAnd []string `json:"tag_and,omitempty"`

	// 标签，多组标签的关系是先取并集，再取补集
	TagNot []string `json:"tag_not,omitempty"`

	// 别名，多个别名之间是Or的关系，取交集
	Alias []string `json:"alias,omitempty"`

	// 注册ID，多个注册Id之间是Or的关系，即取并集
	RegistrationIds []string `json:"registration_id,omitempty"`

	// 用户分群id，分群id
	Segment []string `json:"segment,omitempty"`

	// AB测Id，在页面创建AB测
	AbTest []string `json:"abtest,omitempty"`
}

func NewAudience() *Audience {
	return &Audience{}
}

func (a *Audience) All() *Audience {
	a.all = true
	return a
}

func (a *Audience) SetTag(tag []string) *Audience {
	a.Tag = tag
	return a
}

func (a *Audience) AddTag(tag string) *Audience {
	a.Tag = append(a.Tag, tag)
	return a
}

func (a *Audience) SetTagAnd(tagAnd []string) *Audience {
	a.TagAnd = tagAnd
	return a
}

func (a *Audience) AddTagAnd(tagAnd string) *Audience {
	a.TagAnd = append(a.TagAnd, tagAnd)
	return a
}

func (a *Audience) SetTagNot(tagNot []string) *Audience {
	a.TagNot = tagNot
	return a
}

func (a *Audience) AddTagNot(tagNot string) *Audience {
	a.TagNot = append(a.TagNot, tagNot)
	return a
}

func (a *Audience) SetAlias(alias []string) *Audience {
	a.Alias = alias
	return a
}

func (a *Audience) AddAlias(alias string) *Audience {
	a.Alias = append(a.Alias, alias)
	return a
}

func (a *Audience) SetRegistrationIds(registrationIds []string) *Audience {
	a.RegistrationIds = registrationIds
	return a
}

func (a *Audience) AddRegistrationIds(registrationId string) *Audience {
	a.RegistrationIds = append(a.RegistrationIds, registrationId)
	return a
}

func (a *Audience) SetSegment(segment []string) *Audience {
	a.Segment = segment
	return a
}

func (a *Audience) AddSegment(segment string) *Audience {
	a.Segment = append(a.Segment, segment)
	return a
}

func (a *Audience) SetAbTest(abTest []string) *Audience {
	a.AbTest = abTest
	return a
}

func (a *Audience) AddAbTest(abTest string) *Audience {
	a.AbTest = append(a.AbTest, abTest)
	return a
}

func (a *Audience) ToJsonElement() interface{} {
	if a.all {
		return All
	}
	return a
}
