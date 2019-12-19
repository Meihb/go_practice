/*
garbage collection
Go语言自带垃圾回收机制（GC）。GC 通过独立的进程执行，它会搜索不再使用的变量，并将其释放。需要注意的是，
GC 在运行时会占用机器资源。

GC 是自动进行的，如果要手动进行 GC，可以使用 runtime.GC() 函数，显式的执行 GC。显式的进行 GC 只在某
些特殊的情况下才有用，比如当内存资源不足时调用 runtime.GC() ，这样会立即释放一大片内存，但是会造成程
序短时间的性能下降。

finalizer（终止器）是与对象关联的一个函数，通过runtime.SetFinalizer来设置，如果某个对象定义了finalizer，
当它被 GC 时候，这个 finalizer 就会被调用，以完成一些特定的任务，例如发信号或者写日志等。

在Go语言中 SetFinalizer 函数是这样定义的：
func SetFinalizer(x, f interface{})

参数说明如下：
参数 x 必须是一个指向通过 new 申请的对象的指针，或者通过对复合字面值取址得到的指针。
参数 f 必须是一个函数，它接受单个可以直接用 x 类型值赋值的参数，也可以有任意个被忽略的返回值。

SetFinalizer 函数可以将 x 的终止器设置为 f，当垃圾收集器发现 x 不能再直接或间接访问时，它会清理 x 并调用 f(x)。

另外，x 的终止器会在 x 不能直接或间接访问后的任意时间被调用执行，不保证终止器会在程序退出前执行，因此一般终止器只
用于在长期运行的程序中释放关联到某对象的非内存资源。例如，当一个程序丢弃一个 os.File 对象时没有调用其 Close 方法，
该 os.File 对象可以使用终止器去关闭对应的操作系统文件描述符。

终止器会按依赖顺序执行：如果 A 指向 B，两者都有终止器，且 A 和 B 没有其它关联，那么只有 A 的终止器执行完成，并且
A 被释放后，B 的终止器才可以执行。

如果 *x 的大小为 0 字节，也不保证终止器会执行。

此外，我们也可以使用SetFinalizer(x, nil)来清理绑定到 x 上的终止器。
提示：终止器只有在对象被 GC 时，才会被执行。其他情况下，都不会被执行，即使程序正常结束或者发生错误。
*/
package main

import (
	"log"
	"runtime"
	"time"
)

type Road int

func findRoad(r *Road) {

	log.Println("road:", *r)
}

func entry() {
	var rd Road = Road(999)
	r := &rd

	runtime.SetFinalizer(r, findRoad)
}

func main() {

	entry()
	for i := 0; i < 10; i++ {
		time.Sleep(time.Second)
		runtime.GC()
	}

}
