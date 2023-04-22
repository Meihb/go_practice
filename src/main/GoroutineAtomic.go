package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)
var (
	shutdown int64
	wgAtomic       sync.WaitGroup
)
func main() {
	/*
	runtime.GOMAXPROCS(逻辑CPU数量)

	这里的逻辑CPU数量可以有如下几种数值：
	<1：不修改任何数值。
	=1：单核心执行。
	>1：多核并发执行。

	*/

	fmt.Println(runtime.GOMAXPROCS(runtime.NumCPU()))
	wgAtomic.Add(2)
	go doWork("A")
	go doWork("B")
	time.Sleep(1 * time.Second)
	fmt.Println("Shutdown Now")
	atomic.StoreInt64(&shutdown, 1)
	wgAtomic.Wait()
}
func doWork(name string) {
	defer wgAtomic.Done()
	for {
		fmt.Printf("Doing %s Work\n", name)
		time.Sleep(250 * time.Millisecond)
		if atomic.LoadInt64(&shutdown) == 1 {
			fmt.Printf("Shutting %s Down\n", name)
			break
		}
	}
}
