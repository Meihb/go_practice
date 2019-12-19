package main

import "fmt"

/*
结构体可以包含一个或多个匿名（或内嵌）字段，即这些字段没有显式的名字，只有字段的类型是必须的，此时类型也就是字段的名
字。匿名字段本身可以是一个结构体类型，即结构体可以包含内嵌结构体。
可以粗略地将这个和面向对象语言中的继承概念相比较，随后将会看到它被用来模拟类似继承的行为。Go语言中的继承是通过内嵌或
组合来实现的，所以可以说，在 Go语言中，相比较于继承，组合更受青睐。
*/
type innerS struct {
	in1 int
	in2 int
	b   int
}
type outerS struct {
	b      int
	c      float32
	int    //anonymous field 类型即名字
	innerS //anonymous field
}

// 车轮
type Wheel struct {
	Size int
}

// 车
type Car struct {
	Wheel
	// 引擎
	Engine struct {
		Power int    // 功率
		Type  string // 类型
	}
}

func main() {
	outer := new(outerS)
	outer.b = 6
	outer.c = 7.5
	outer.int = 2
	//outer.innerS.in1 = 5
	outer.in2 = 5 //内嵌结构体的字段可以省略中间步骤,当然测试下来仅访问最短路径 但如存在路径相同的两个,会报错
	outer.innerS.b = 4
	outer.innerS.in2 = 4

	fmt.Println(outer)
	outer2 := outerS{
		b:      0,
		c:      0,
		int:    0,
		innerS: innerS{},
	}
	fmt.Println("outer2 is :", outer2)

	c := Car{
		// 初始化轮子
		Wheel: Wheel{
			Size: 18,
		},
		// 初始化引擎
		Engine: struct { //此处是struct 是因为上面定义的时候 直接将结构体定义在嵌入的结构体中
			Power int
			Type  string
		}{
			Type:  "1.4T",
			Power: 143,
		},
	}
	fmt.Printf("%+v\n", c)

}
