package main

import "fmt"


/*
多个defer的执行顺序为“后进先出”；

defer、return、返回值三者的执行逻辑应该是：return最先执行，return负责将结果写入返回值中；接着defer开始执行
一些收尾工作；最后函数携带当前返回值退出。
 */
func f0() int {//无名返回值
	var a int
	defer func() {
		fmt.Println("in defer")
		a++
	}()
	return a//无名返回值感觉就是一个copy,因此defer中再不能获得返回值的地址,有名返回值则不同
}
func f1() (result int) {//有名返回值
	defer func() {
		fmt.Println("in defer")
		result++
	}()
	return 0 //1
}

func f2() (r int) {
	t := 5
	defer func() {
		fmt.Println("in defer")
		fmt.Println("before", t)
		t = t + 5
		fmt.Println("after", t)
	}()
	t++
	fmt.Println("during defer and return")
	return t
}

func f3() (t int) {
	t = 5
	defer func() {
		fmt.Println("in defer")
		t = t + 5
	}()
	return t
}
func f4() (r int) {
	defer func(r int) {
		fmt.Println("in defer")
		r = r + 5
	}(r)
	return 1
}

func f5()  {

	defer func() {
		if v := recover();v == 11 {
			fmt.Printf("v: %#v\n",v)
		}
		fmt.Printf("defer1...\n")
	}()

	defer func() {
		fmt.Printf("defer2...\n")
	}()

	array := [2]int{1,2}
	fmt.Println("array: ",array[1])
	panic(11)

}

func main() {
	fmt.Println(f0())
	fmt.Println(f1())
	fmt.Println(f2())
	fmt.Println(f3())
	fmt.Println(f4())

	f5()
}
