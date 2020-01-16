package main

import "fmt"

/*
多个defer的执行顺序为“后进先出”；

defer、return、返回值三者的执行逻辑应该是：return最先执行，return负责将结果写入返回值中；接着defer开始执行
一些收尾工作；最后函数携带当前返回值退出。
*/
func f0() int { //无名返回值
	var a int
	defer func() {
		fmt.Println("in defer 1")
		a++
		fmt.Println(a)
	}()
	defer func(i int) {
		fmt.Println("in defer 2")
		fmt.Println(i)//0 为啥呢,也就是说函数调用的参数,他是copy的,不会引用传值
		a++//但是函数内部可以引用传值的概念读取外部变量,理解其中差异
		i++ //所以这个i++是局部变量,不影响a

		/*
		所以之前有个疑点,到底能够读取外部变量和保存参数是何意,就在于此了
		函数创造时copy,函数运行时是point
		 */
	}(a)
	a++
	return a //无名返回值感觉就是一个copy,因此defer中再不能获得返回值的地址,有名返回值则不同
}
func f1() (result int) { //有名返回值
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

func f5() {

	defer func() {
		if v := recover(); v == 11 {
			fmt.Printf("v: %#v\n", v)
		}
		fmt.Printf("defer1...\n")
	}()

	defer func() {
		fmt.Printf("defer2...\n")
	}()

	array := [2]int{1, 2}
	fmt.Println("array: ", array[1])
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
