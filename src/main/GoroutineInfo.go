package main

import "fmt"

/*
goroutine 是一种非常轻量级的实现，可在单个进程里执行成千上万的并发任务，它是Go语言并发设计的核心。
说到底 goroutine 其实就是线程，但是它比线程更小，十几个 goroutine 可能体现在底层就是五六个线程，
而且Go语言内部也实现了 goroutine 之间的内存共享。

使用 go 关键字就可以创建 goroutine，将 go 声明放到一个需调用的函数之前，在相同地址空间调用运行这个函数，
这样该函数执行时便会作为一个独立的并发线程，这种线程在Go语言中则被称为 goroutine。
*/

func GetThingDone(a, b int) { //忽略返回值
	fmt.Println((a + b))
}
func main() {
	//go 关键字放在方法调用前新建一个 goroutine 并执行方法体
	go GetThingDone(1, 2);
	//新建一个匿名方法并执行

	go func(param1, param2 int) {
	}(1, 2)
	//直接新建一个 goroutine 并在 goroutine 中执行代码块
	//go { 这个方式不能用了
	//	fmt.Printf("aaa")
	//}


	/*
	channel 是进程内的通信方式，因此通过 channel 传递对象的过程和调用函数时的参数传递行为比较一致，
	比如也可以传递指针等。如果需要跨进程通信，我们建议用分布式系统的方法来解决，比如使用 Socket 或者 HTTP 等通信协议。
	Go语言对于网络方面也有非常完善的支持。
	 */
	//ci := make(chan int)
	//cs := make(chan string)
	//cf := make(chan interface{})

	/*
	在编写 Socket 网络程序时，需要提前准备一个线程池为每一个 Socket 的收发包分配一个线程。开发人员需要在线程数量和 CPU
	数量间建立一个对应关系，以保证每个任务能及时地被分配到 CPU 上进行处理，同时避免多个任务频繁地在线程间切换执行而损失
	效率。

	虽然，线程池为逻辑编写者提供了线程分配的抽象机制。但是，如果面对随时随地可能发生的并发和线程处理需求，线程池就不是
	非常直观和方便了。能否有一种机制：使用者分配足够多的任务，系统能自动帮助使用者把任务分配到 CPU 上，让这些任务尽量并
	发运作。这种机制在 Go语言中被称为 goroutine。

	goroutine 是 Go语言中的轻量级线程实现，由 Go 运行时（runtime）管理。Go 程序会智能地将 goroutine 中的任务合理地分
	配给每个 CPU。

	所有 goroutine 在 main() 函数结束时会一同结束。！！！
	 */


	/*
	goroutine 是轻量级线程,跟协程coroutine有不少的区别
	goroutines 意味着并行（或者可以以并行的方式部署），coroutines 一般来说不是这样的，goroutines 通过通道来通信；
	coroutines 通过让出和恢复操作来通信，goroutines 比 coroutines 更强大，也很容易从 coroutines 的逻辑复用到
	goroutines。

	狭义地说，goroutine 可能发生在多线程环境下，goroutine 无法控制自己获取高优先度支持；coroutine 始终发生在单线程，
	coroutine 程序需要主动交出控制权，宿主才能获得控制权并将控制权交给其他 coroutine。
	goroutine 间使用 channel 通信，coroutine 使用 yield 和 resume 操作。

	goroutine 和 coroutine 的概念和运行机制都是脱胎于早期的操作系统。
	coroutine 的运行机制属于协作式任务处理，早期的操作系统要求每一个应用必须遵守操作系统的任务处理规则，应用程序在不需要
	使用 CPU 时，会主动交出 CPU 使用权。如果开发者无意间或者故意让应用程序长时间占用 CPU，操作系统也无能为力，表现出来的
	效果就是计算机很容易失去响应或者死机。

	goroutine 属于抢占式任务处理，已经和现有的多线程和多进程任务处理非常类似。应用程序对 CPU 的控制最终还需要由操作
	系统来管理，操作系统如果发现一个应用程序长时间大量地占用 CPU，那么用户有权终止这个任务。
	 */
}
