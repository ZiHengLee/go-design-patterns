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

// GetSingleton goer常用方式
func GetSingleton() Singleton {
	once.Do(func() {
		instance = &singleton{}
	})

	return instance
}

// 饿汉
//var instance = &singleton{}
//func GetSingleton() *singleton {
//	return instance
//}

// 懒汉
//func GetSingleton() *singleton {
//	if instance == nil {
//		instance = &singleton{}
//	}
//	return instance
//}

// 支持并发的方式
//var mu sync.Mutex
//func GetSingleton() *singleton {
//	mu.Lock()
//	defer mu.Unlock()
//	if instance == nil {
//		instance = &singleton{}
//	}
//	return instance
//}

// 双重锁定
//var mu sync.Mutex
//func GetSingleton() *singleton {
//	if instance == nil {
//		mu.Lock()
//		defer mu.Unlock()
//		if instance == nil {
//			instance = &singleton{}
//		}
//	}
//	return instance
//}
