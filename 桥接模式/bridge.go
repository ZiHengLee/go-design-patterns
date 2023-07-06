package 桥接模式

import "fmt"

//就用上面的例子，开关灯

// LightI 灯接口
type LightI interface {
	Run()
	Name() string
}

type A struct{}

func (a *A) Name() string {
	return "A"
}

func (a *A) Run() {
	fmt.Println(a.Name() + ":I'm shining")
}

type B struct{}

func (b *B) Name() string {
	return "B"
}

func (b *B) Run() {
	fmt.Println(b.Name() + ":I'm shining")
}

// SwitchI 开关接口
type SwitchI interface {
	Op()
}

// SwitchA A开关
type SwitchA struct {
	l LightI
}

func NewSwitchA(a LightI) *SwitchA {
	return &SwitchA{l: a}
}

func (sa *SwitchA) Op() {
	fmt.Println("im SwitchA open " + sa.l.Name())
	sa.l.Run()
}

// SwitchB B开关
type SwitchB struct {
	l LightI
}

func NewSwitchB(b LightI) *SwitchB {
	return &SwitchB{l: b}
}

func (sb *SwitchB) Op() {
	fmt.Println("im SwitchB open " + sb.l.Name())
	sb.l.Run()
}
