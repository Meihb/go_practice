package singleInstance

import (
	"sync"
)

type Tool struct {
}

var instance *Tool

//线程排他锁
var lock sync.Mutex

//首先判断是否为空实例,若否,加锁再实例化
func getInstance() *Tool {
	if instance == nil {
		lock.Lock()
		if (instance == nil) {
			instance = new(Tool)
		}
		lock.Unlock()
	}
	return instance
}

//sync,Once
var once sync.Once

func getInstance2() *Tool {
	once.Do(func() {
		instance = new(Tool)
	})
	return instance
}
