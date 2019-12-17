package main

import (
	"fmt"
	"math"
)

/*
收器类型可以是（几乎）任何类型，不仅仅是结构体类型，任何类型都可以有方法，甚至可以是函数类型，可以是 int、bool、string 或数组的别名类型，
但是接收器不能是一个接口类型，因为接口是一个抽象定义，而方法却是具体实现，如果这样做了就会引发一个编译错误invalid receiver type…。
接收器也不能是一个指针类型，但是它可以是任何其他允许类型的指针，一个类型加上它的方法等价于面向对象中的一个类，一个重要的区别是，在Go语言中，
类型的代码和绑定在它上面的方法的代码可以不放置在一起，它们可以存在不同的源文件中，唯一的要求是它们必须是同一个包的。
类型 T（或 T）上的所有方法的集合叫做类型 T（或 T）的方法集。
因为方法是函数，所以同样的，不允许方法重载，即对于一个类型只能有一个给定名称的方法，但是如果基于接收器类型，是有重载的：具有同样名字的
方法可以在 2 个或多个不同的接收器类型上存在，比如在同一个包里这么做是允许的。
*/

//背包结构
type Bag struct {
	items []int
}

//将物品放入背包的过程
//面向过程风格
func InsertProcedureOriented(bag *Bag, itemid int) {
	bag.items = append(bag.items, itemid)
}

//面向过程风格
func (bag *Bag) InsertObjectOriented(itemid int) {
	bag.items = append(bag.items, itemid)
}

//矢量结构
type Vec2 struct {
	X, Y float32
}

//加法
func (v Vec2) add(other Vec2) Vec2 {
	return Vec2{
		X: v.X + other.X,
		Y: v.Y + other.Y,
	}
}

//减法
func (v Vec2) sub(other Vec2) Vec2 {
	return Vec2{
		X: v.X - other.X,
		Y: v.Y - other.Y,
	}
}

//乘
func (v Vec2) scale(other Vec2) Vec2 {
	return Vec2{
		X: v.X * other.X,
		Y: v.Y * other.Y,
	}
}

//距离
func (v Vec2) Distance(other Vec2) float32 {
	dx := v.X - other.X
	dy := v.Y - other.Y
	return float32(math.Sqrt(float64(dx*dx + dy*dy)))
}

//插值 标准化
func (v Vec2) Normalize() Vec2 {
	mag := v.X*v.X + v.Y*v.Y
	if mag > 0 {
		oneOverMag := 1 / float32(math.Sqrt(float64(mag)))
		return Vec2{
			X: v.X * oneOverMag,
			Y: v.Y * oneOverMag,
		}
	}
	return Vec2{
		X: 0,
		Y: 0,
	}
}

//玩家
type Player struct {
	currPos   Vec2    //当前位置
	targetPos Vec2    //目标位置
	speed     float32 //移动速度
}

func main() {
	bag := &Bag{}
	InsertProcedureOriented(bag, 12)
	fmt.Println(bag)
	bag.InsertObjectOriented(23)
	fmt.Println(bag)

}
