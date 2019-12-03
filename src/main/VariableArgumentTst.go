package main

import (
	"fmt"
)

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
	for _, arg := range args {
		switch arg.(type) {
		case int:
			fmt.Println(arg, "is an int value")
		case string:
			fmt.Println(arg, "is a string value")
		case int64:
			fmt.Println(arg, "is an int64 value")
		default:
			fmt.Println(arg, "is a unknown type")
		}

	}
}

func main() {
	var v1 int = 1
	var v2 int64 = 234
	var v3 string = "hello"
	var v4 float32 = 1.234
	/*
		注意观察传参方式,和py3不一样,not *list 或者**dict
		golang 传参要么多个值传入,要么切片/数组...
	*/
	MyPrintf(v1, v2, v3, v4)
}
