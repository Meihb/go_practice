/*
Recover 是一个Go语言的内建函数，可以让进入宕机流程中的 goroutine 恢复过来，recover 仅在延迟函数 defer 中有效，
在正常的执行过程中，调用 recover 会返回 nil 并且没有其他任何效果，如果当前的 goroutine 陷入恐慌，调用 recover 可以捕获到
panic 的输入值，并且恢复正常的执行。

panic 和 recover 的组合有如下特性：
有 panic 没 recover，程序宕机。
有 panic 也有 recover，程序不会宕机，执行完对应的 defer 后，从宕机点退出当前函数后继续执行。
提示
虽然 panic/recover 能模拟其他语言的异常机制，但并不建议在编写普通函数时也经常性使用这种特性。

在 panic 触发的 defer 函数内，可以继续调用 panic，进一步将错误外抛，直到程序整体崩溃。

如果想在捕获错误时设置当前函数的返回值，可以对返回值使用命名返回值方式直接进行设置。

从panic退出到recover往下执行
*/
package main

import (
	"fmt"
	"runtime"
)

//崩溃时需要传递的上下文信息
type panicContext struct {
	function string
}

//保护方式允许一个函数
func ProtectRun(entry func()) {
	//延迟处理匿名函数并执行
	defer func() {
		//发生宕机时,获取panic传递的上下文并打印
		err := recover()

		switch err.(type) {
		case runtime.Error: //运行时错误
			fmt.Println("runtime error:", err)
		default:
			fmt.Println("error:", err)
		}
	}()
	entry()

}

func main() {
	fmt.Println("运行前")
	//允许一段手动触发的错误
	ProtectRun(func() {
		fmt.Println("手动宕机前")
		//使用panic传递上下文
		panic(&panicContext{"手动触发Panic"})
		fmt.Println("手动宕机后")
	})

	ProtectRun(func() {
		fmt.Println("赋值宕机前")
		var a *int
		*a = 1
		fmt.Println("赋值宕机后")
	})
}
