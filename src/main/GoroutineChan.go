package main

import (
	"fmt"
	"runtime"
)

func main() {
	/*
		通道是引用类型，需要使用 make 进行创建，格式如下：
		通道实例 := make(chan 数据类型)

		数据类型：通道内传输的元素类型。
		通道实例：通过make创建的通道句柄。
	*/
	//ch1 := make(chan int)                 // 创建一个整型类型的通道
	//ch2 := make(chan interface{})         // 创建一个空接口类型的通道, 可以存放任意格式
	//type Equip struct{ /* 一些字段 */ }
	//ch2 := make(chan *Equip)             // 创建Equip指针类型的通道, 可以存放*Equip

	// 创建一个空接口通道
	//ch := make(chan interface{})
	// 将0放入通道中
	/*
		1) 通道发送数据的格式
		通道的发送使用特殊的操作符<-，将数据通过通道发送的格式为：
		通道变量 <- 值  ch是符号左值则是发送数据,右值则是接受数据

		通道变量：通过make创建好的通道实例。
		值：可以是变量、常量、表达式或者函数返回值等。值的类型必须与ch通道的元素类型一致。
	*/
	//ch <- 0
	// 将hello字符串放入通道中
	//把数据往通道中发送时，如果接收方一直都没有接收，那么发送操作将持续阻塞。注意力,发送发将会一直阻塞
	// Go 程序运行时能智能地发现一些永远无法发送成功的语句并做出提示
	//ch <- "hello" //fatal error: all goroutines are asleep - deadlock!

	/*
		通道接收同样使用<-操作符，通道接收有如下特性：
		① 通道的收发操作在不同的两个 goroutine 间进行。

		由于通道的数据在没有接收方处理时，数据发送方会持续阻塞，因此通道的接收必定在另外一个 goroutine 中进行。

		② 接收将持续阻塞直到发送方发送数据。

		如果接收方接收时，通道中没有发送方发送数据，接收方也会发生阻塞，直到发送方发送数据为止。

		③ 每次接收一个元素。
		通道一次只能接收一个数据元素。

		data:=<-ch 阻塞式接受
		data, ok := <-ch 非阻塞式接受 和type assertion的con,error类似
		<-ch 忽略返回值
	*/

	ch := make(chan int)

	/*
		go func() {
			fmt.Println("start goroutine")

			ch <- 0
			fmt.Println("exit goroutine")
		}()

		fmt.Println("wait goroutine")
		<-ch
		fmt.Println("all done")
	*/

	/*
			循环


		go func() {
			for i := 0; i <= 3; i++ {
				ch <- i
				time.Sleep(time.Second * 1)
			}
		}()

		for data:= range ch {//for range chan 可以一直接受数据,不过好像就不能作为非阻塞了
			// 打印通道数据
			fmt.Println(data)
			// 当遇到数据0时, 退出接收循环
			if data == 3 {
				break
			}
		}
	*/

	//测试非阻塞
	go func() {
		fmt.Println("start goroutine")

		runtime.Gosched()
		ch <- 0
		fmt.Println("exit goroutine")
	}()

	data, err := <-ch
	if (!err) {
		fmt.Println("Got nothing")
	} else {
		fmt.Println("Got it ",data)
	}



}
