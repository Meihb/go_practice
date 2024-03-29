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
type Struct2 struct {
	name string
	age uint32
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

func FuncCaller2(interface{}) {

}

//函数的声明不能直接实现接口，需要将函数定义为类型后，使用类型实现结构体，当类型方法被调用时，还需要调用函数本体。
// 所以这里是区别,interface method 必须是type类型才能,单独设置一个func是不可以的,必须是type定义一下
//func (f FuncCaller2 ) Call(p interface{}) {
//	f(p)
//}
type FuncCalle3 func(interface{})

func main() {
	//声明接口变量
	var invoker Invoker
	//实例化结构体
	s2 := new(Struct2)
	s := new(Struct)
	fmt.Println(s2)
	s3:= Struct2{age:12,name:"asdsa"}
	fmt.Println(s3)

	//将实例化的结构体复制到接口
	invoker = s //s类型是*Struct,已经实现了Call方法,因此可有赋值给invoker
	// invoker = s2//无法通过编译,因为s2未实现 Call() method

	//使用接口调用实例化结构体的方法Struct.Call()
	invoker.Call("hello")
	s.Call("hello")

	//将匿名函数转为FuncCaller类型,再赋值给接口
	invoker = FuncCaller(func(v interface{}) {
		fmt.Println("from function", v)
	})
	FuncCaller_3 := FuncCaller(func(v interface{}) {
		fmt.Println("this is function ", v)
	})
	fmt.Printf("%t", FuncCaller_3) //上面难道就是function type的字面值申明,好奇怪
	//使用接口调用FuncCaller，内部会调用函数本体
	invoker.Call("hello1")

}
