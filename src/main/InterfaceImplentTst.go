/*
所以这里面体现的implent思想和php或者java不同,后两者语言其实是显示通过implents规范了class必须实现相关方法
go其实结构体或类似的结构是现在某方法,但其实并没有显示implents该Interface,只有当你把结构体实例赋值给interface之时
go编译器会检查赋值的结构体是否implents了此接口的方法,GOT IT?

*/
package main

import "fmt"

//调用器接口
type Invoker interface {
	//需要实现一个方法
	Call(interface{})
}

//结构体类型
type Struct struct {
}

//实现Invoker的Call s *Struct 为 receiver
func (s *Struct) Call(p interface{}) {
	fmt.Println("from struct", p)
}

//函数定义为类型
type FuncCaller func(interface{})

//实现Invoker的Call
func (f FuncCaller) Call(p interface{}) {
	//调用f函数本体
	f(p)
}

func main() {
	//声明接口变量
	var invoker Invoker
	//实例化结构体
	s := new(Struct)

	//将实例化的结构体复制到接口
	invoker = s //s类型是*Struct,已经实现了Call方法,因此可有赋值给invoker

	//使用接口调用实例化结构体的方法Struct.Call()
	invoker.Call("hello")

	//将匿名函数转为FuncCaller类型,再赋值给接口
	invoker = FuncCaller(func(v interface{}) {
		fmt.Println("from function", v)
	})
	//使用接口调用FuncCaller，内部会调用函数本体
	invoker.Call("hello1")

}
