package 建造者模式

type Controller struct {
	builder Builder
}

func (c *Controller) SetBuilder(builder Builder) {
	c.builder = builder
}

type Builder interface {
	NewMe()
	BuildName()
	BuildSex()
	BuildAge()
	FinalMe() *Me
}

type MyBuilder struct {
	Me *Me
}

type Me struct {
	Name string
	Sex  string
	Age  string
}

func (me *Me) MeDescribe() string {
	return me.Name + me.Age + me.Age
}

func (mb *MyBuilder) NewMe() {
	mb.Me = new(Me)
}

func (mb *MyBuilder) BuildName() {
	mb.Me.Name = "我叫头脑风暴"
}

func (mb *MyBuilder) BuildSex() {
	mb.Me.Sex = "我是男性"
}

func (mb *MyBuilder) BuildAge() {
	mb.Me.Age = "我今年25岁"
}

func (mb *MyBuilder) FinalMe() *Me {
	return mb.Me
}
