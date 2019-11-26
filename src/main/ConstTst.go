package main

import "fmt"

/*
常量的值必须是能够在编译时就能够确定的，可以在其赋值表达式中涉及计算过程，但是所有用于计算的值必须在编译期
间就能获得。
如果是批量声明的常量，除了第一个外其它的常量右边的初始化表达式都可以省略，如果省略初始化表达式则表示使用前
面常量的初始化表达式，对应的常量类型也是一样的。
当然Go语言中的变量、函数、常量名称的首字母也可以大写，如果首字母大写，则表示它可以被其它的包
访问（类似于 Java 中的 public）；如果首字母小写，则表示它只能在本包中使用 (类似于 Java 中
private）。
*/
const (
	ca = 1
	cb
	cc = 2
	cd
)

type Weekday int

/*
常量声明可以使用 iota 常量生成器初始化，它用于生成一组以相似规则初始化的常量，但是不用每行都写一遍初始化表达式。
在一个 const 声明语句中，在第一个声明的常量所在的行，iota 将会被置为 0，然后在每一个有常量声明的行加一
*/
const (
	Sunday Weekday = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

func main() {
	fmt.Println(a, b, c, d);
	fmt.Println(Sunday, Monday, Tuesday, Wednesday, Thursday, Friday, Saturday);
}
