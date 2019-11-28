package main

import (
	"container/list"
	"fmt"
)
/*
  variableName:=list.New()
or
  var variableName list.List
 */
func main()  {
	l:= list.New()
	l.PushBack("first")
	l.PushFront(67)

	// 尾部添加后保存元素句柄
	element := l.PushBack("fist")
	// 在fist之后添加high
	l.InsertAfter("high", element)
	// 在fist之前添加noon
	l.InsertBefore("noon", element)
	// 使用
	l.Remove(element)
	
	
	//遍历链表
	for i:=l.Front(); i!=nil;i=i.Next()  {
		fmt.Println()
	}
}