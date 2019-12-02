/*
对数据的操作进行多步骤的处理被称为链式处理
*/
package main

import (
	"fmt"
	"strings"
)

func StringProcess(list []string, chain []func(string) string) {
	//遍历字符串
	for index, value := range list {
		result := ""
		//遍历处理链函数
		for _, proc := range chain {
			/*
				和php好像不一样啊,多次循环,内循环可以访问到外循环的变量,反之不同,
				go的内存管理相对严格,想必在编译期间就处理此问题
			*/
			result = proc(value)
		}

		list[index] = result
	}
}

func RemovePrefix(str string) string {
	return strings.TrimPrefix(str, "go")
}

func main() {
	// 待处理的字符串列表
	list := []string{
		"go scanner",
		"go parser",
		"go compiler",
		"go printer",
		"go formater",
	}
	// 处理函数链
	chain := []func(string) string{
		RemovePrefix,
		strings.TrimSpace,
		strings.ToUpper,
	}

	//处理字符串
	StringProcess(list, chain)

	// 输出处理好的字符串
	for _, str := range list {
		fmt.Println(str)
	}
}
