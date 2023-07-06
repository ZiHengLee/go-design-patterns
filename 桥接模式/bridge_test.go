package 桥接模式

import (
	"testing"
)

func TestBridge(t *testing.T) {
	//这样就可以随意组合了
	a := NewSwitchA(&A{})
	a.Op()
	b := NewSwitchA(&B{})
	b.Op()
	c := NewSwitchB(&A{})
	c.Op()
	d := NewSwitchB(&B{})
	d.Op()
}
