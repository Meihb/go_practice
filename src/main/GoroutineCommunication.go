package main

import (
	"fmt"
	"runtime"
	"sync"
)

/*
A 是传统模式的通信模型,python也是如此
 */
var counter int = 0
func Count(lock *sync.Mutex) {
	lock.Lock()
	counter++
	fmt.Println("in gocoroutine",counter)
	lock.Unlock()
}

/*
B 竞争状态简述
 */
var (
	countB int32
	wg    sync.WaitGroup
)
func incCount() {
	defer wg.Done()
	for i := 0; i < 2; i++ {
		value := countB
		runtime.Gosched()
		value++
		countB = value
	}
}
func main() {
	/*
	A
	 */
	/*
	lock := &sync.Mutex{}
	for i := 0; i < 10; i++ {
		go Count(lock)
	}
	for {
		lock.Lock()
		c := counter
		lock.Unlock()
		runtime.Gosched()
		if c >= 10 {
			break
		}
	}
*/

	/*
	B
	 */

	wg.Add(2)
	go incCount()
	go incCount()
	wg.Wait()
	fmt.Println(countB)//结果多变,2、3、4都有可能,因为没有竞争保护
}