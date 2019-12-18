package main

//实例化一个通过字符串解析函数切片的map
var eventByName = make(map[string][]func(interface{}))

//注册时间,提供事件名和回调函数
func RegisterEvent(name string, callback func(interface{})) {
	
}
