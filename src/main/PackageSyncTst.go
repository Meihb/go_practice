package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var lock sync.Mutex

//排它锁
func tstMutex() {
	ch := make(chan struct{}, 2)
	var l sync.Mutex
	go func() {
		l.Lock()
		defer l.Unlock()
		fmt.Println("goroutine1: 我会锁定大概 2s")
		time.Sleep(time.Second * 2)
		fmt.Println("goroutine1: 我解锁了，你们去抢吧")
		ch <- struct{}{}
	}()
	go func() {
		fmt.Println("goroutine2: 等待解锁")
		l.Lock()
		defer l.Unlock()
		fmt.Println("goroutine2: 欧耶，我也解锁了")
		ch <- struct{}{}
	}()
	// 等待 goroutine 执行结束
	for i := 0; i < 2; i++ {
		<-ch
	}
}

//读写锁 读锁等于没锁呀,写锁不就是排它锁嘛,区别是读锁和写锁之间互斥
/*
当有一个 goroutine 获得写锁定，其它无论是读锁定还是写锁定都将阻塞直到写解锁；
当有一个 goroutine 获得读锁定，其它读锁定仍然可以继续；
当有一个或任意多个读锁定，写锁定将等待所有读锁定解锁之后才能够进行写锁定。

所以说这里的读锁定（RLock）目的其实是告诉写锁定，有很多协程或者进程正在读取数据，写操作需要等它们读（读解锁）完
才能进行写（写锁定）。

我们可以将其总结为如下三条：
同时只能有一个 goroutine 能够获得写锁定；
同时可以有任意多个 gorouinte 获得读锁定；
同时只能存在写锁定或读锁定（读和写互斥）
*/
var rw sync.RWMutex
var count int

func tstRW() {
	ch := make(chan struct{}, 10)
	for i := 0; i < 5; i++ {
		go read(i, ch)
	}
	for i := 0; i < 5; i++ {
		go write(i, ch)
	}
	for i := 0; i < 10; i++ {
		<-ch
	}
}
func read(n int, ch chan struct{}) {
	rw.RLock()
	fmt.Printf("goroutine %d 进入读操作...\n", n)
	v := count
	fmt.Printf("goroutine %d 读取结束，值为：%d\n", n, v)
	rw.RUnlock()
	ch <- struct{}{}
}
func write(n int, ch chan struct{}) {
	rw.Lock()
	fmt.Printf("goroutine %d 进入写操作...\n", n)
	v := rand.Intn(1000)
	count = v
	fmt.Printf("goroutine %d 写入结束，新值为：%d\n", n, v)
	rw.Unlock()
	ch <- struct{}{}
}

func main() {
	var a = 0
	for i := 0; i < 100; i++ {
		//  func(idx int) { //匿名函数直接执行
		//	a += 1
		//	fmt.Println(a)
		//}(i)
		go func(idx int) {
			lock.Lock()
			defer lock.Unlock()
			a += 1
			fmt.Printf("goroutine %d,a=%d\n", idx, a)
		}(i)
	}
	//time.Sleep(time.Second)

	//tstMutex()

	for i := 0; i < 100; i++ {

	}
}
