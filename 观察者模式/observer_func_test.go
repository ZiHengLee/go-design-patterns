package 观察者模式

import (
	"fmt"
	"testing"
	"time"
)

func add(a, b int) {
	s := fmt.Sprintf("%v+%v=%v", a, b, a+b)
	fmt.Println(s)
}

func sub(a, b int) {
	s := fmt.Sprintf("%v-%v=%v", a, b, a-b)
	fmt.Println(s)
}

func TestAsyncEventBus_Publish(t *testing.T) {
	bus := NewAsyncEventBus()
	bus.Subscribe("加", add)
	bus.Subscribe("减", sub)

	bus.Publish("加", 1, 2)
	bus.Publish("减", 3, 3)
	time.Sleep(1 * time.Second)
}
