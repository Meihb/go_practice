package main

/*
type 类型名 struct {
    字段1 字段1类型
    字段2 字段2类型
    …
}
*/

type Point struct {
	X int;
	Y int;
}

var p Point //实例化

func main() {
	p.X = 1
	p.Y = 2


	//创建指针类型的结构体,new,但是为什么？？？为什么指针类型的struct可以直接访问字段
	ins2 := new(Point)
	ins2 .X = 2

}
