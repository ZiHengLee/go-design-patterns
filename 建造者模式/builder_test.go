package 建造者模式

import (
	"fmt"
	"testing"
)

func TestBuilder(t *testing.T) {
	c := new(Controller)
	// 创建建造者
	builder := new(MyBuilder)
	c.SetBuilder(builder)
	c.builder.NewMe()
	c.builder.BuildName()
	c.builder.BuildAge()
	c.builder.BuildSex()

	me := c.builder.FinalMe()
	fmt.Println(me.MeDescribe())
}
