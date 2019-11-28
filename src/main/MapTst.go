package main

import (
	"fmt"
	"sort"
)

func main() {
	//这个定义方式好不喜欢啊,为什么不是 var name [KeyType]ValueType,这个map开起来不美
	var mapLit map[string]int
	//var mapCreated map[string]float32
	var mapAssigned map[string]int
	mapLit = map[string]int{"one": 1, "two": 2}

	mapCreated := make(map[string]float32)
	mapAssigned = mapLit
	mapCreated["key1"] = 4.5
	mapCreated["key2"] = 3.14159
	mapAssigned["two"] = 3
	fmt.Printf("Map literal at \"one\" is: %d\n", mapLit["one"])
	fmt.Printf("Map created at \"key2\" is: %f\n", mapCreated["key2"])
	fmt.Printf("Map assigned at \"two\" is: %d\n", mapLit["two"])
	fmt.Printf("Map literal at \"ten\" is: %d\n", mapLit["ten"])

	/*
	new 可以用來初始化泛型，並且返回儲存位址。所以通常我們會用指標變數來接 new 過後的型別。特別要注意的是，
	new 會自動用 zeroed value 來初始化型別，也就是字串會是""，number 會是 0，channel, func, map, slice 等等則會是 nil


	因為初始化的 map 會是 nil map，不像其他的 primitive type 一樣有預設值。
		var mapPointer = new(map[string][ ]int)
		mapPointerObj := *mapPointer
		fmt.Printf("type:%T",mapPointerObj)
		mapPointerObj["one"] = [ ]int{0, 1, 2}
		mapPointerObj["two"] = [ ]int{3, 4, 5}
		mapPointerObj["three"] = [ ]int{6, 7, 8}
		fmt.Println(mapPointerObj)




		type Person struct {
		  Name string
		  Age  int
		}

		func main() {
		    p := &Person{}
		    p := new(Person)
		}
		好處是上面的 Person 也可以根據自己想要傳入的值額外再做設定，但 new 則是全部的 field 都會直接塞 zeroed value。
	*/
	scene := make(map[string]int)
	// 准备map数据
	scene["route"] = 66
	scene["brazil"] = 4
	scene["china"] = 960
	// 声明一个切片保存map数据
	var sceneList []string
	// 将map数据遍历复制到切片中
	for k := range scene {
		sceneList = append(sceneList, k)
	}
	// 对切片进行排序
	sort.Strings(sceneList)
	// 输出
	fmt.Println(sceneList)



}
