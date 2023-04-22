package main

import (
	"fmt"
	"sync"
	"time"
)

/*
sync.NewCond(&mutex)：生成一个cond，需要传入一个mutex，因为阻塞等待通知的操作以及通知解除阻塞的操作就是基于sync.Mutex来实现的。
sync.Wait()：用于等待通知
sync.Signal()：用于发送单个通知
sync.Broadcat()：用于广播
 */
func main() {

	mutex := sync.Mutex{}
	var cond = sync.NewCond(&mutex)
	mail := 1
	go func() {
		for count := 0; count <= 15; count++ {
			time.Sleep(time.Second)
			mail = count
			cond.Broadcast()
		}
	}()
	// worker1
	go func() {
		for mail != 5 { // 触发的条件，如果不等于5，就会进入cond.Wait()等待，此时cond.Broadcast()通知进来的时候，wait阻塞解除，进入下一个循环，此时发现mail != 5，跳出循环，开始工作。
			cond.L.Lock()
			cond.Wait()
			cond.L.Unlock()
		}
		fmt.Println("worker1 started to work")
		time.Sleep(3 * time.Second)
		fmt.Println("worker1 work end")
	}()
	// worker2
	go func() {
		for mail != 10 {
			cond.L.Lock()
			cond.Wait()
			cond.L.Unlock()
		}
		fmt.Println("worker2 started to work")
		time.Sleep(3 * time.Second)
		fmt.Println("worker2 work end")
	}()
	// worker3
	go func() {
		for mail != 10 {
			cond.L.Lock()
			cond.Wait()
			cond.L.Unlock()
		}
		fmt.Println("worker3 started to work")
		time.Sleep(3 * time.Second)
		fmt.Println("worker3 work end")
	}()
	select {}
}
