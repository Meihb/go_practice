package main

import "fmt"

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

	var mapPointer = new(map[string][3]int)

	(*mapPointer)["one"] = [3]int{0, 1, 2}
	(*mapPointer)["two"] = [3]int{3, 4, 5}
	(*mapPointer)["three"] = [3]int{6, 7, 8}

}
