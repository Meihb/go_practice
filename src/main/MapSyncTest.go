/*
	sync map
*/
package main

import (
	"fmt"
	"sync"
)

/*
fatal error: concurrent map read and map write
错误信息显示，并发的 map 读和 map 写，也就是说使用了两个并发函数不断地对 map 进行读和写而发生了竞态问题，
map 内部会对这种并发操作进行检查并提前发现。
func main() {
	m := make(map[int]int)

	//goroutine ,go关键字
	go func() {
		//不停对map写入
		for {
			m[1] = 1
		}
	}()

	go func() {
		for {
			_ = m[1]
		}
	}()

	for {

	}
}

*/

func main()  {
	var scene sync.Map
	// 将键值对保存到sync.Map
	scene.Store("greece", 97)
	scene.Store("london", 100)
	scene.Store("egypt", 200)
	// 从sync.Map中根据键取值
	fmt.Println(scene.Load("london"))
	// 根据键删除对应的键值对
	scene.Delete("london")
	// 遍历所有sync.Map中的键值对,这个眼熟哦,java的eachable匿名函数
	scene.Range(func(k, v interface{}) bool {
		fmt.Println("iterate:", k, v)
		return true
	})


}
