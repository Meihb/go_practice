/*
定义一个延迟函数
把这个函数放到栈上,在其外部的包含函数即将return,返回参数到调用方法之前调用,即运行到最外层方法"}"时调用,可以用来做一些资源
释放,如文件io的关闭等
其实是return之前,不是}之前哦,有区别,def里面修改外部函数返回值的话,外外部调用方法获得的回参也会被改变
类似try()finally{}???
*/
package main

import (
	"fmt"
	"os"
)

func doSomething(fileName string) {
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()
}

/*
就像闭包一样，如果不是defer函数方法内的变量会向上一层函数访问变量，重新做计算。
*/
func doSomething1() {
	v := 10
	defer func() { //可以读写外部变量,类似于闭包函数
		fmt.Println(v)
		v++
		fmt.Println(v)
	}()
	v += 5
}

/*
defer 读写命名的返回值
以我看来,返回值需要约束类型这是肯定的,如果同时约束了返回值的名字,相当于在函数栈中初始化了一个同名类型值以待返回
但是有一个好处,比如次函数中,return 5 其实等于rev=5;return rev,这个5实际上被赋值给了rev
所以这里其实和doSomething1函数一样,其实就是读了外部变量而已

多个defer,后进先出
*/
func doSth2() (rev int) {
	defer func() {
		rev++
		fmt.Println(rev)
	}()
	defer fmt.Println(1)
	defer fmt.Println(2);
	return 5
}

/*
捕获异常
*/
func doSth3() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Print(err)
		}

	}()

	fmt.Println("Running...")
	panic("run error")
}
func main() {
	doSomething1()
	fmt.Println(doSth2())
	
	doSth3()
}
