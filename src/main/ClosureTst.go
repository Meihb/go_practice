/*
闭包是引用了自由变量的函数，被引用的自由变量和函数一同存在，即使已经离开了自由变量的环境也不会被释放或者删除，
在闭包中可以继续使用这个自由变量，因此，简单的说：
函数 + 引用环境 = 闭包
成为对应用环境的	"捕获"
*/
package main

import "fmt"

/*
闭包的记忆效应,有点像python3 的@function?
被捕获到闭包中的变量让闭包本身拥有了记忆效应，闭包中的逻辑可以修改闭包捕获的变量，变量会跟随闭包生命期一直存在，
闭包本身就如同变量一样拥有了记忆效应。
*/
func Accumulate(value int) func() int {
	//返回一个闭包
	return func() int {
		//累加
		value++
		return value
	}
}

/*生成器
 */
func playerGen(name string) func() (string, int) {
	//血量
	hp := 150
	return func() (s string, i int) {
		return name, hp
	}
}
func main() {
	/*
		闭包对它作用域上部的变量可以进行修改，修改引用的变量会对变量进行实际修改
		这个上部 其实是调用之时的上部呀好像,不对,声明必须在上部,但实际值其实是在调用上部
	*/
	str := "hello world"
	foo := func() {
		fmt.Println(str)//hello world2
		//匿名函数中访问str
		str = "hello dude"
		fmt.Println(str)//hello dude
	}
	str = "hello world2"
	//调用匿名函数
	foo()

	println(str)//hello dude

	//创建一个累加器,初始值为1
	accumulator := Accumulate(1)
	//累加1并打印
	fmt.Println(accumulator())
	fmt.Println(accumulator())
	//打印累加器的函数地址
	fmt.Printf("%p\n", accumulator)
	//创建累加器2
	accumulator2 := Accumulate(10)
	fmt.Println(accumulator2())
	//打印累加器的函数地址
	fmt.Printf("%p\n", accumulator2)

	//创建一个玩家生成器
	generator := playerGen("mhb")
	//返回玩家名字与血量
	name, hp := generator()
	fmt.Println(name,hp)
}
