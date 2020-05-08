package main

import "fmt"

func main() {
	var a = "hello"
	switch a {
	case "hello":
		fmt.Println(1)
	case "world":
		fmt.Println(2)
	case "mum", "daddy": //一分多支
		fmt.Println("family")
	default:
		fmt.Println(0)
	}

	var r int = 11
	var s = "hello"
	switch {
	/*
	分支表达式,再次不要切不能写入r,自伤而下只能命中一次表达式,除非fallthrough(就和php中swtch case中不加入break)
	不对啊,好像fallthrough是无脑 执行下一个case,无论是否满足表达式exp
	 */
	case r > 10 && r < 20:
		fmt.Println(r)
		fallthrough
	case s == "hello":
		fmt.Println("hello")
		fallthrough
	case s == "hello1":
		fmt.Println("sss")
	case s != "world":
		fmt.Println("world")
	}
}
