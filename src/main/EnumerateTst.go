package main

import "fmt"

// 声明芯片类型
type ChipType int

const (
	None ChipType = iota
	CPU    // 中央处理器
	GPU    // 图形处理器
)

//func (Type) funcName() (return type)??
//所以这个定义的是ChipType类型的方法,当格式化输出类型为%s即字符串类型时,会调用此方法
func (c ChipType) String() string {
	switch c {
	case None:
		return "None"
	case CPU:
		return "CPU"
	case GPU:
		return "GPU"
	}

	return "N/A"
}

func main() {

	// 输出CPU的值并以整型格式显示
	fmt.Printf("%s %d", CPU, CPU)
}
