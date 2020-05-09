package main

import "fmt"

/*
配合错误代码处理
*/
func main() {
	for x := 0; x < 4; x++ {
		fmt.Printf("x:%d\n", x)
		for y := 0; y < 4; y++ {
			fmt.Printf("\t\ty:%d\n", y)
			if y == 2 {
				fmt.Printf("goto\n")
				//跳转到goto标签
				goto breakHere
			}
		}
	}
	//手动返回 ,避免执行goto标签,
	return

//breakHere2:
//	fmt.Println("not done")
	//label,跳转之后即离开之前的循环区域了
	//另外此标签如果之前的代码不退出的话是会被执行的
breakHere:
	fmt.Println("done")
}
