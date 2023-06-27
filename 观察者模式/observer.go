package 观察者模式

import "fmt"

type Subject interface {
	register(observer Observer)
	remove(observer Observer)
	notify(msg string)
}

type Sub1 struct {
	observers []Observer
}

func (sub *Sub1) register(observer Observer) {
	sub.observers = append(sub.observers, observer)
}

func (sub *Sub1) remove(observer Observer) {
	for i, ob := range sub.observers {
		if ob == observer {
			sub.observers = append(sub.observers[:i], sub.observers[i+1:]...)
		}
	}
}

func (sub *Sub1) notify(msg string) {
	for _, o := range sub.observers {
		o.update(msg)
	}
}

type Observer interface {
	update(string)
}

type Ob1 struct {
}

func (o *Ob1) update(name string) {
	fmt.Println("我是ob1,我收到了商品名改成了：", name)
}

type Ob2 struct {
}

func (o *Ob2) update(name string) {
	fmt.Println("我是ob2,我收到了商品名改成了：", name)
}
