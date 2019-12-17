package main

import (
	"fmt"
	"unsafe"
)

/*
type 类型名 struct {
    字段1 字段1类型
    字段2 字段2类型
    …
}
*/

/* 佳句
黄梅戏《梁山伯与祝英台》中，梁山伯对女扮男装的祝英台起疑，于是有了下面的对白：“英台不是女儿身，因何耳上有环痕？ ”
“耳环痕有原因，梁兄何必起疑云，村里酬神多庙会，年年由我扮观音，梁兄做文章要专心，你前程不想想钗裙。 ”
“我从此不敢看观音。”
*/
type Point struct {
	X int;
	Y int;
}

var p Point //实例化

type MyStruct struct {
	field int
}

func (self MyStruct) modify_struct_value() {
	fmt.Println(unsafe.Pointer(&self))
	self.field = 2
}
func (self MyStruct) copy_my_self() MyStruct {
	fmt.Println(unsafe.Pointer(&self))
	return self
}

type testint int

func (t testint) double() {
	t = 2 * t
	fmt.Printf("double:%d\r\n", t)
}
func (t *testint) square() {
	*t = *t * *t
	fmt.Printf("square:%d\r\n", *t)
}

type TestReceiver struct {
	name string
}

func (t TestReceiver) func1() {
	t.name = "结构体非指针"
}
func (t *TestReceiver) PointerFunc1() {
	t.name = "结构体指针"
}

func main() {
	p.X = 1
	p.Y = 2

	//创建指针类型的结构体,new,但是为什么？？？为什么指针类型的struct可以直接访问字段,it just works
	//在Go语言中，访问结构体指针的成员变量时可以继续使用.，这是因为Go语言为了方便开发者访问结构体指针的成员变量，
	// 使用了语法糖（Syntactic sugar）技术，将 ins.Name 形式转换为 (*ins).Name(只是结构体哦目前来看,primitive不是)
	ins2 := new(Point)
	(*ins2).X = 2
	ins2.Y = 3
	//可推导的type 可omitted
	var ins3 = new(Point) //new的时候好像不能够赋值
	ins3.X = 2

	local_variable := MyStruct{1}
	fmt.Println(unsafe.Pointer(&local_variable))
	local_variable.modify_struct_value()
	fmt.Println(local_variable) // {1}
	copied := local_variable.copy_my_self()
	fmt.Println(unsafe.Pointer(&copied))

	//receiver为结构/结构指针的区别
	var ti1 testint = 2
	ti1.double()
	fmt.Println(ti1)
	ti1.square()
	fmt.Println(ti1)

	tr1 := TestReceiver{"默认名"}
	tr1.func1()
	fmt.Println(tr1) //copy的杂鱼呐
	tr1.PointerFunc1()
	fmt.Println(tr1) //修改了内部属性

	//使用键值对 初始化结构体
	type People struct {
		name  string
		child *People
	}

	relation := &People{
		name: "grandpa",
		//child: nil,//哦,原来编辑器自动给我零值化了
		child: &People{
			name:  "我",
			child: nil,
		}, //此处需要, ，不解
	}
	relation.child.child = new(People)
	relation.child.child.name = "大"

	//匿名结构体 和匿名函数有点像哦
	insa := struct {
		name string
		id   int
	}{
		"mhb",
		12,
	};
	fmt.Println(insa.name)
	//myfunc :=func()(int){
	//	return 1
	//}()


}
