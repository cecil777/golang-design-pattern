package singleton

import "sync"

// Singleton 是单例模式机接口，导出的
// 通过该接口可以避免 GetInstance 返回一个包私有类型的指针
type Singleton interface {
	foo()
}

type singleton struct {
}

func (s singleton) foo() {
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
