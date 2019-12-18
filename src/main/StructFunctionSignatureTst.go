package main

/*
函数签名（或者类型签名，抑或方法签名）定义了 函数或方法的输入与输出。

签名可包含以下内容：

参数 及参数的 类型
一个的返回值及其类型
可能会抛出或传回的异常 go无视这点？
该方法在 面向对象程序中的可用性方面的信息（如public、static或prototype）。 go无视此点

无论是普通函数还是结构体的方法，只要它们的签名一致，与它们签名一致的函数变量就可以保存普通函数或是结构体方法。

*/
import "fmt"

type class struct {
	name string
}

//结构体方法Do
func (c *class) Do(v int) {
	fmt.Println(c.name, "call method do:", v) //method 面向对象
}
//结构体方法Do
func (c class) Do1(v int) {
	fmt.Println(c.name, "call method do:", v) //method 面向对象
}
//普通函数的Do
func Do(v int) {
	fmt.Println("call function do:", v) //function 是面向过程的
}

func main() {
	//申明一个函数回调
	var delegate func(int)
	//创建结构体实例
	c := new(class) //妹的这个new果然和其他语言用法一样,虽然本质不同
	c.name = "sss"
	// 将回调设为c的Do方法
	delegate = c.Do //疑似 会复制(?)实例的数据,诶呦不是复制,是引用哦,还是因为receiver是引用类型？
	c.name="sss1"
	//delegate = class.Do 必须用实例而非struct,想当然也
	//调用
	delegate(100)
	//将回调设为
	delegate = Do
	//调用
	delegate(100)

	delegate= c.Do1//哦原理和receiver类型是一样的,引用型receiver表现其引用性质,否则则是复制性质
	c.name="sss2"
	//调用
	delegate(100)

	/*
		一个事件系统拥有如下特性：
		能够实现事件的一方，可以根据事件 ID 或名字注册对应的事件。
		事件发起者，会根据注册信息通知这些注册者。
		一个事件可以有多个实现方响应。
	*/
}
