package main

import "fmt"

/*
变量申明
整型和浮点型变量的默认值为 0 和 0.0。
字符串变量的默认值为空字符串。
布尔型变量默认为 bool。
切片、函数、指针变量的默认为 nil。
 */
var aa int
var (
	a int
	b string
	c []float32
	d func() bool
	e struct{
		x int
	}
)
/*
 第二个括号(int,int)是返回类型
 */
func getData() (int,int){
	return 100,200
}
func main() {
	/*
	简短模式 的变量申明,必须显示初始化、不提供数据类型、必须在函数内部
	:=是申明符号?
	注意：由于使用了:=，而不是赋值的=，因此推导声明写法的左值变量必须是没有定义过的变量。若定义过，将会发生编译错误。
	注意：在多个短变量声明和赋值中，至少有一个新声明的变量出现在左值中，即便其他变量名可能是重复声明的，编译器也不会报错
	*/
	a := 1
	a,s := a+1,"let go"
	//s :="dsa"//s已经被右值推导的类型,会报错
	fmt.Println(s)


	//http.Handle("/", http.FileServer(http.Dir(".")))
	//http.ListenAndServe(":8090", nil)


	/*
	多重复值
	多重赋值时，变量的左值和右值按从左到右的顺序赋值。
	 */
	var ai,bi  = 100,200
	ai,bi = bi,ai
	fmt.Println(ai,bi)

	/*
	匿名变量
	 */
	var an,_=getData()
	var _,bn=getData()
	fmt.Println(an,bn,aa)



}


