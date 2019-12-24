/*
type 接口类型名 interface{
    方法名1( 参数列表1 ) 返回值列表1
    方法名2( 参数列表2 ) 返回值列表2
    …
}
对各个部分的说明：
接口类型名：使用 type 将接口定义为自定义的类型名。Go语言的接口在命名时，一般会在单词后面添加 er，如有写操作的接口叫
Writer，有字符串功能的接口叫 Stringer，有关闭功能的接口叫 Closer 等。
方法名：当方法名首字母是大写时，且这个接口类型名首字母也是大写时，这个方法可以被接口所在的包（package）之外的代码访
问。
参数列表、返回值列表：参数列表和返回值列表中的参数变量名可以被忽略
*/
package main

import "fmt"

// 一个服务需要满足能够开启和写日志的功能
type Service interface {
	Start()     // 开启服务
	Log(string) // 日志输出
}

// 日志器
type Logger struct {
}

// 实现Service的Log()方法
func (g *Logger) Log(l string) {
}

// 游戏服务
type GameService struct {
	Logger // 嵌入日志器
}

// 实现Service的Start()方法
func (g *GameService) Start() {
}

/*
嵌套结构体来实现接口
一个接口的方法，不一定需要由一个类型完全实现，接口的方法可以通过在类型中嵌入其他类型或者结构体来实现。也就是说，
使用者并不关心某个接口的方法是通过一个类型完全实现的，还是通过多个结构嵌入到一个结构体中拼凑起来共同实现的。

Service 接口定义了两个方法：一个是开启服务的方法（Start()），一个是输出日志的方法（Log()）。使用 GameService
结构体来实现 Service，GameService 自己的结构只能实现 Start() 方法，而 Service 接口中的 Log() 方法已经被一个
能输出日志的日志器（Logger）实现了，无须再进行 GameService 封装，或者重新实现一遍。所以，选择将 Logger 嵌入到
GameService 能最大程度地避免代码冗余，简化代码结构。详细实现过程如下：
*/
func main() {
	var s Service = new(GameService)
	s.Start()
	s.Log("hello")

	/*
		接口断言 Type Assertion 是一个使用在接口值上的操作，用于检查接口类型变量所持有的值是否实现了期望的接口或
		者具体的类型。

	其中，x 表示一个接口的类型，T 表示一个具体的类型（也可为接口类型）。

	该断言表达式会返回 x 的值（也就是 value）和一个布尔值（也就是 ok），可根据该布尔值判断 x 是否为 T 类型：
	如果 T 是具体某个类型，类型断言会检查 x 的动态类型是否等于具体类型 T。如果检查成功，类型断言返回的结果是 x
	的动态值，其类型是 T。
	如果 T 是接口类型，类型断言会检查 x 的动态类型是否满足 T。如果检查成功，x 的动态值不会被提取，返回值是一个
	类型为 T 的接口值。
	无论 T 是什么类型，如果 x 是 nil 接口值，类型断言都会失败。

	*/
	var x interface{}
	x = 10
	value, ok := x.(int)
	fmt.Println(value, ",", ok)

	value1,ok1 := s.(Service)
	fmt.Println(value1, ",", ok1)


	var a int
	a = 10
	getType(a)
}

func getType(a interface{}) {
	switch a.(type) {
	case int:
		fmt.Println("the type of a is int")
	case string:
		fmt.Println("the type of a is string")
	case float64:
		fmt.Println("the type of a is float")
	default:
		fmt.Println("unknown type")
	}
}