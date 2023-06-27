package 观察者模式

import "testing"

func TestObserver(t *testing.T) {
	sub := &Sub1{}
	sub.register(&Ob1{})
	sub.register(&Ob2{})
	sub.notify("A")
	sub.notify("B")
}
