package main

import "fmt"

/*
...type语法糖
interface{} 表达的是任意类型
*/
func myFunc(args ...int) () {
	for _, arg := range args {
		fmt.Println(arg)
	}
}

func MyPrintf(args ...interface{}) {
	
}
