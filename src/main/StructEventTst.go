package main

import "fmt"

//实例化一个通过字符串解析函数切片的map
var eventByName = make(map[string][]func(interface{}))

//注册时间,提供事件名和回调函数
func RegisterEvent(name string, callback func(interface{})) {
	//通过名字查找事件列表
	list := eventByName[name]
	//在切片中添加函数
	list = append(list, callback)
	eventByName[name] = list
}

//调用事件
func CallEvent(name string, param interface{}) {
	//通过名字查找事件列表
	list := eventByName[name]
	//遍历所有回调
	for _, callback := range list {
		callback(param)
	}

}

//角色结构体
type Actor struct {
}

//为角色添加一个事件处理函数
func (a *Actor) OnEvent(param interface{}) {
	fmt.Println("actor event:", param)
}

//全局事件
func GlobalEvent(param interface{}) {
	fmt.Println("global event:", param)
}
func main() {
	//实例化角色
	actor := new(Actor)
	//注册回调
	RegisterEvent("OnSkill", actor.OnEvent)
	//在OnSkill上注册全局事件
	RegisterEvent("OnSkill", GlobalEvent)

	//调用事件
	CallEvent("OnSkill", 100)
}
