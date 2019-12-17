package main

import "fmt"

/*
模拟构造函数
*/

//模拟构造函数 go 好像没有函数重载
type Cat struct {
	Color string
	Name  string
}

func NewCatTst(color string, name string) *Cat {
	return &Cat{
		color,
		name, //要是把}提上前来就可以不写逗号,好奇怪的feature
	}
}

//模拟父级结构调用

type BlackCat struct {
	Cat // 嵌入Cat, 类似于派生
}

func newCat(name string) *Cat {
	return &Cat{
		Name: name,
	}
}

func newBlackCat(color string) *BlackCat {
	cat := &BlackCat{}
	cat.Color = color
	return cat
}

type animal interface {
	pass()
}

func main() {

	//创建Cat
	cat := newCat("lisa")
	fmt.Printf("cat is %s\r\n", cat)

	blackcat := newBlackCat("black")
	fmt.Printf("blackcat is %s\r\n", blackcat)
}
