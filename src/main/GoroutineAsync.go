package main
import (
	"fmt"
	"sync"
)
var wgAsync sync.WaitGroup
func main() {
	fmt.Println("aaa")

	for i := 0; i < 5; i++ {
		wgAsync.Add(1)
		go AsyncFunc(i)
	}
	wgAsync.Wait()
}

type Vector []float64
// 分配给每个CPU的计算任务
func (v Vector) DoSome(i, n int, u Vector, c chan int) {
	for ; i < n; i++ {
		v[i] += u.Op(v[i])
	}
	c <- 1 // 发信号告诉任务管理者我已经计算完成了
}
const NCPU = 16 // 假设总共有16核
func (v Vector) DoAll(u Vector) {
	c := make(chan int, NCPU) // 用于接收每个CPU的任务完成信号
	for i := 0; i < NCPU; i++ {
		go v.DoSome(i*len(v)/NCPU, (i+1)*len(v)/NCPU, u, c)
	}
	// 等待所有CPU的任务完成
	for i := 0; i < NCPU; i++ {
		<-c // 获取到一个数据，表示一个CPU计算完成了
	}
	// 到这里表示所有计算已经结束
}


func AsyncFunc(index int) {
	defer wgAsync.Done()
	sum := 0
	for i := 0; i < 10000; i++ {
		sum += 1
	}
	fmt.Printf("线程%d, sum为:%d\n", index, sum)
}