package main

import "fmt"

func main() {
	//if是一个特殊的语句,其内部的变量不在当前作用域呢,当前的写法可以把err的作用域尽可能缩小
	/*
	初始语句是在第一次循环前执行的语句，一般使用初始语句执行变量初始化，如果变量在此处被声明，其作用域将被局限在这个 for
	的范围内。
	仅当其在这个for的范围内啊
	 */
	if err := Connect(); err != nil {
		fmt.Println(err)
		return
	}

	if a2:=1;a2==2{

	}
}
