package main

import (
	"fmt"
)

func main() {
	var a [3]int             // 定义三个整数的数组
	fmt.Println(a[0])        // 打印第一个元素
	fmt.Println(a[len(a)-1]) // 打印最后一个元素

	// 打印索引和元素 这个range 可真是膈应 key ,value 简短模式赋值 range variable,emmmmm
	for i, v := range a {
		fmt.Printf("%d %d\n", i, v)
	}

	// 仅打印元素
	for _, v := range a {
		fmt.Printf("%d\n", v)
	}
	var r [3]int = [3]int{1, 2}
	fmt.Println(r[2]) // "0"

	//这个地方不使用简短模式好像不行,或者说认为...这个只能作为类型推导？
	q := [...]int{1, 2, 3}
	fmt.Printf("%T\n", q) // "[3]int"

	// 声明一个二维整型数组，两个维度的长度分别是 4 和 2
	var array [4][2]int
	// 使用数组字面量来声明并初始化一个二维整型数组
	array = [4][2]int{{10, 11}, {20, 21}, {30, 31}, {40, 41}}
	// 声明并初始化数组中索引为 1 和 3 的元素
	array = [4][2]int{1: {20, 21}, 3: {40, 41}}
	// 声明并初始化数组中指定的元素
	array = [4][2]int{1: {0: 20}, 3: {1: 41}}

	for k, v := range array {
		fmt.Printf("k:%v,v:%v\n", k, v)
		for k1, v1 := range v {
			fmt.Printf("\t\tk1:%v,v1:%v \n", k1, v1)
		}
	}
}
