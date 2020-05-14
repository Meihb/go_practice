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

//面向对象风格
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
func (v Vec2) scale(s float32) Vec2 {
	return Vec2{
		X: v.X * s,
		Y: v.Y * s,
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

//获取玩家当前位置
func (p *Player) Pos() Vec2 {
	return p.currPos
}

//是否到达
func (p *Player) IsArrived() bool {
	//如果当前玩家为之与目标位置的距离不超过移动的步长,则认之为已到达目标
	return p.currPos.Distance(p.targetPos) < p.speed
}

//逻辑更新
func (p *Player) Update() {
	if !p.IsArrived() {
		//计算当前位置指向目标的朝向
		dir := p.targetPos.sub(p.currPos).Normalize()

		p.currPos = p.currPos.add(dir.scale(p.speed))
	}
}

// 移动到某个点就是设置目标位置
func (p *Player) MoveTo(v Vec2) {
	p.targetPos = v
}
func NewPlayer(speed float32) *Player {
	return &Player{
		currPos:   Vec2{},
		targetPos: Vec2{},
		speed:     speed,
	}
}
func main() {
	bag := &Bag{}
	InsertProcedureOriented(bag, 12)
	fmt.Println(bag)
	bag.InsertObjectOriented(23)
	fmt.Println(bag)

	// 实例化玩家对象，并设速度为0.5
	p := NewPlayer(0.5)
	// 让玩家移动到3,1点
	p.MoveTo(Vec2{3, 1})
	// 如果没有到达就一直循环
	for !p.IsArrived() {
		// 更新玩家位置
		p.Update()
		// 打印每次移动后的玩家位置
		fmt.Println(p.Pos())
	}
}
