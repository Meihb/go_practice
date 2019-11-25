package main

import (
	"flag"
	"fmt"
)

//定义命令行参数
var mode  = flag.String("mode","","process mode")

func swap(a, b *int) {
	//去a指针的值,赋给临时变量t
	t := *a
	//取b指针的值,赋给a指针指向的变量
	*a = *b
	*b = t
}
func swap1(a, b *int) {
	b,a = a,b
	fmt.Printf("a:%d b:%d\r\n",*a,*b)
}
func main() {
	var cat int = 1
	var str string = "banana"
	fmt.Println("%p %p", &cat, &str)

	// 准备一个字符串类型
	var house = "Malibu Point 10880,90265"
	ptr := &house

	//打印ptr类型
	fmt.Printf("ptr type：%p\r\n", ptr)
	//打印Ptr地址
	fmt.Printf("address:%p\r\n", ptr)
	//打印Ptr自身地址
	fmt.Printf("address:%p\r\n", &ptr)
	value := *ptr
	//对指针进行取值操作
	fmt.Printf("value type:%T\r\n", value)

	fmt.Printf("value:%s\n", value)

	// 准备两个变量, 赋值1和2
	x, y := 1, 2
	// 交换变量值
	swap(&x, &y)
	// 输出变量值
	fmt.Println(x, y)

	x, y = 1, 2
	// 交换变量值
	swap1(&x, &y)//交换失败，很简单呐,ab函数变量互换影响不到xy的指向
	// 输出变量值
	fmt.Println(x, y)

	//解析命令行参数
	flag.Parse()
	fmt.Println(*mode)


	//new函数创建指针 new() 函数可以创建一个对应类型的指针，创建过程会分配内存，被创建的指针指向默认值
	//之前的所有操作都是基于所有内存已经创建并存储,当前情况会创建一个默认值的内存
	str2 := new(string)
	*str2 = "Go语言教程"
	fmt.Println(*str2)
}
