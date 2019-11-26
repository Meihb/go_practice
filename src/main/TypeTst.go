package main

import (
	"fmt"
	"reflect"
	"time"
)


//类型定义
type NewInt int

//类型别名 区别于类型定义只有一个 = 符号,但实际上区别较大
//类型别名实际上的类型依旧是本来类型
type intAlias = int
type MyDuration = time.Duration

/*
non-local 非本地类型不能够定义方法
不能在一个非本地的类型 time.Duration 上定义新方法，非本地类型指的就是 time.Duration 不是在 main 包中定义的，
而是在 time 包中定义的，与 main 包不在同一个包中，因此不能为不在一个包中的类型定义方法。
func (m MyDuration)EasySet(a string){

}
*/
type MyDuration2 time.Duration

//除了在local package定义类型方法以外,只能通过类型定义来处理以上问题,make them be us
func (m MyDuration2) EasySet(s string) {

}

//定义结构 商标,once 我想为什么不能够在func里面type呢,可是目前来看go不能再func里面嵌入定义func,如果在main里面或其他
//func里面type 或者type alias,那也无法定义此type的方法,如此本地化即无意义,总不能为了别名而别名吧
type Brand struct {
}

func (b Brand) show() {
	fmt.Printf("Type is %T\n", b)
}

//为Brand取别名
type FakeBrand = Brand

//定义车辆结构，这是默认同名吗？
//难道别名的意义就着？配合次语法省略代码吗
type Vehicle struct {
	//嵌入两个结构
	fb    FakeBrand
	brand Brand
	FakeBrand
	Brand
}

func main() {

	//结果显示 a 的类型是 main.NewInt，表示 main 包下定义的 NewInt 类型，a2 类型是 int，
	// IntAlias 类型只会在代码中存在，编译完成时，不会有 IntAlias 类型。
	var a NewInt
	fmt.Printf("a type: %T\r\n", a)
	var a2 intAlias
	fmt.Printf("a2 type:%T\n", a2)

	var vehicle Vehicle
	//指定调用FakeBrand的show方法
	vehicle.fb.show()
	vehicle.FakeBrand.show()
	//指定调用Brand的show方法
	vehicle.brand.show()
	vehicle.Brand.show()
	//取vehicle的类型反射对象
	ta := reflect.TypeOf(vehicle)

	//遍历打印
	for i := 0; i < ta.NumField(); i++ {
		//成员信息
		f := ta.Field(i)
		fmt.Printf("FieldName:%v,FieldType:%+v\r\n", f.Name, f.Type)
	}

}
