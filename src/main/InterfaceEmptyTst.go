package main

import "fmt"

/*
空接口是接口类型的特殊形式，空接口没有任何方法，因此任何类型都无须实现空接口。从实现的角度看，任何值都满足这
个接口的需求。因此空接口类型可以保存任何值，也可以从空接口中取出原值。
提示
空接口类型类似于 C# 或 Java 语言中的 Object、C语言中的 void*、C++ 中的 std::any。在泛型和模板出现前，空接
口是一种非常灵活的数据抽象保存和使用的方法。
空接口的内部实现保存了对象的类型和指针。使用空接口保存一个数据的过程会比直接用数据对应类型的变量保存稍慢。因此
在开发中，应在需要的地方使用空接口，而不是在所有地方使用空接口。
 */

func main()  {
	var any interface{}
	any = 1
	fmt.Println(any)
	any = "hello"
	fmt.Println(any)
	any = false
	fmt.Println(any)

	//任何值都可以赋予空接口不意味着空接口可以赋予给任何值,只能用断言
	var i int
	any = 1
	//i = any
	i,ok :=any.(int)
	fmt.Println(i,ok)


}