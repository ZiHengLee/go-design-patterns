package 单例模式

import (
	"fmt"
	"sync"
)

type Singleton interface {
	hello()
}

// singleton 是单例模式类，包私有的
type singleton struct{}

func (s singleton) hello() {
	fmt.Println("world")
}

var (
	instance *singleton
	once     sync.Once
)

// GetInstance 用于获取单例模式对象
func GetInstance() Singleton {
	once.Do(func() {
		instance = &singleton{}
	})

	return instance
}
