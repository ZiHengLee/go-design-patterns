### 单例模式：确保无论对象被实例化多少次，全局都只有一个实例存在

#### 懒汉式：单例对象是在包加载时立即被创建
```golang
var instance = &singleton{}
func GetSingleton() *singleton {
    return instance
}
```
#### 饿汉式：实例会在第一次被使用时被创建
```golang
func GetSingleton() *singleton {
    if instance == nil {
        instance = &singleton{}
    }
    return instance
}
```
#### 并发单例模式：
```golang
var mu sync.Mutex
func GetSingleton() *singleton {
    mu.Lock()
    defer mu.Unlock()
    if instance == nil {
        instance = &singleton{}
    }
    return instance
}
```
#### 双重锁定模式：只有在实例为空的时候初始化
```golang
var mu sync.Mutex
func GetSingleton() *singleton {
    if instance == nil {
        mu.Lock()
        defer mu.Unlock()
        if instance == nil {
            instance = &singleton{}
        }
    }
    return instance
}
```
#### goer常用sync.once来写单例模式
```golang
func GetSingleton() Singleton {
    once.Do(func() {
        instance = &singleton{}
    })
    return instance
}
```

