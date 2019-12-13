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

func main() {
	p.X = 1
	p.Y = 2

	//创建指针类型的结构体,new,但是为什么？？？为什么指针类型的struct可以直接访问字段
	ins2 := new(Point)
	ins2.X = 2

	local_variable := MyStruct{1}
	fmt.Println(unsafe.Pointer(&local_variable))
	local_variable.modify_struct_value()
	fmt.Println(local_variable) // {1}
	copied := local_variable.copy_my_self()
	fmt.Println(unsafe.Pointer(&copied))
}
