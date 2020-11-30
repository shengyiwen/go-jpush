package jpush

// 推送目标，推送设备对象，表示一条推送可以推送到哪些设备列表
// 如果是推送给所有设备，则直接填写all
type Audience struct {

	// 如果all为true的话，则剩下标签无效
	all bool

	// 标签，多组标签的关系是Or的关系，即取并集
	tag []string

	// 标签，多组标签的关系是And关系，取交集
	tagAnd []string

	// 标签，多组标签的关系是先取并集，再取补集
	tagNot []string

	// 别名，多个别名之间是Or的关系，取交集
	alias []string

	// 注册ID，多个注册Id之间是Or的关系，即取并集
	registrationIds []string

	// 用户分群id，分群id
	segment []string

	// AB测Id，在页面创建AB测
	abTest []string
}

func (a *Audience) ToJsonElement() interface{} {
	if a.all {
		return All
	}

	var jsonElement = make(map[string]interface{})

	AssertNotNull(a.tag, func() {
		jsonElement["tag"] = a.tag
	})

	AssertNotNull(a.tagAnd, func() {
		jsonElement["tag_and"] = a.tagAnd
	})

	AssertNotNull(a.alias, func() {
		jsonElement["alias"] = a.alias
	})

	AssertNotNull(a.registrationIds, func() {
		jsonElement["registration_id"] = a.registrationIds
	})

	AssertNotNull(a.segment, func() {
		jsonElement["segment"] = a.segment
	})

	AssertNotNull(a.abTest, func() {
		jsonElement["abtest"] = a.abTest
	})

	return jsonElement
}

func NewAudience() *Audience {
	return &Audience{}
}

func (a *Audience) All() *Audience {
	a.all = true
	return a
}

func (a *Audience) Tag(tag []string) *Audience {
	a.tag = tag
	return a
}

func (a *Audience) AddTag(tag string) *Audience {
	a.tag = append(a.tag, tag)
	return a
}

func (a *Audience) TagAnd(tagAnd []string) *Audience {
	a.tagAnd = tagAnd
	return a
}

func (a *Audience) AddTagAnd(tagAnd string) *Audience {
	a.tagAnd = append(a.tagAnd, tagAnd)
	return a
}

func (a *Audience) TagNot(tagNot []string) *Audience {
	a.tagNot = tagNot
	return a
}

func (a *Audience) AddTagNot(tagNot string) *Audience {
	a.tagNot = append(a.tagNot, tagNot)
	return a
}

func (a *Audience) Alias(alias []string) *Audience {
	a.alias = alias
	return a
}

func (a *Audience) AddAlias(alias string) *Audience {
	a.alias = append(a.alias, alias)
	return a
}

func (a *Audience) RegistrationIds(registrationIds []string) *Audience {
	a.registrationIds = registrationIds
	return a
}

func (a *Audience) AddRegistrationId(registrationId string) *Audience {
	a.registrationIds = append(a.registrationIds, registrationId)
	return a
}

func (a *Audience) Segment(segment []string) *Audience {
	a.segment = segment
	return a
}

func (a *Audience) AddSegment(segment string) *Audience {
	a.segment = append(a.segment, segment)
	return a
}

func (a *Audience) AbTest(abTest []string) *Audience {
	a.abTest = abTest
	return a
}

func (a *Audience) AddAbTest(abTest string) *Audience {
	a.abTest = append(a.abTest, abTest)
	return a
}
