package main

import "fmt"

func main() {
	var a [3]int = [3]int{1, 2, 3} //居然初始化的时候不可以跨行,这一点要注意,go好多这个要点

	//切片运算,左闭右开
	/*
			从数组或切片生成新的切片拥有如下特性：
			取出的元素数量为：结束位置 - 开始位置；
			取出元素不包含结束位置对应的索引，切片最后一个元素使用 slice[len(slice)] 获取；
			当缺省开始位置时，表示从连续区域开头到结束位置；
			当缺省结束位置时，表示从开始位置到整个连续区域末尾；
			两者同时缺省时，与切片本身等效；
			两者同时为 0 时，等效于空切片，一般用于切片复位。

		要注意结束为止缺省时和len-1是不同的,缺少了连续内存的最后一个索引
	*/
	fmt.Println(a, a[1:2])
	fmt.Println(a, a[1:1]) //空切片
	fmt.Println(a, a[1:])
	fmt.Println(a, a[:2])
	fmt.Println(a, a[:])
	//fmt.Println(a,a[1:0]); must be low <=high

	/*
		直接声明切片
		注意和数组的区别,当var name []Type为切片,var name[number]Type或者var name[...]Type={...}这样为数组,简单的观察下来
		没有Number并且不指定为...则为切片,宣告了其自动扩展性
	*/
	fmt.Println("declare slice:")
	var strList []string       //声明字符串切片
	var numList []int          //整型切片
	var numListEmpty = []int{} //!=nil 为什么呢,因为其已经被分配了内存,只是没有元素罢了,上面两个都没有被分配内存
	numArr := [...]int{}
	//输出三个切片
	fmt.Println(strList, numList, numListEmpty, numArr)
	//输出三个切片大小
	fmt.Println(len(strList), len(numList), len(numListEmpty), len(numArr))
	//输出三个切片类型
	//输出类型为 []Type和[0]Type,前者为切片,后者为数组
	fmt.Printf("strList:%T,numList:%T,numListEmpty:%T,numArr:%T\n", strList, numList, numListEmpty, numArr)
	//输出三个切片是否为空
	fmt.Println(strList == nil, numList == nil, numListEmpty == nil)

	/*
		make([]Type,size,cap)
		capacity 容量 只是能提前分配空间，降低多次分配空间造成的性能问题
	*/
	a1 := make([]int, 2)
	a2 := make([]int, 2, 10)

	fmt.Println(a1, a2)
	fmt.Println(len(a1), len(a2))
	/*
		使用 make() 函数生成的切片一定发生了内存分配操作，但给定开始与结束位置（包括切片复位）的切片只是将新的切片结构指向已经分
		配好的内存区域，设定开始与结束位置，不会发生内存分配操作
	*/
	fmt.Println(make([]int, 0) == nil)

	/*
		append
	*/
	var numbers []int
	for i := 0; i < 10; i++ {
		numbers = append(numbers, i)
		fmt.Printf("len: %d  cap: %d pointer: %p\n", len(numbers), cap(numbers), numbers)
	}

	var aa1 = []int{1, 2, 3}
	aa1 = append([]int{0}, aa1...)          // 在开头添加1个元素
	aa1 = append([]int{-3, -2, -1}, aa1...) // 在开头添加1个切片

	//在第i个位置插入x
	i, x := 1, 17
	aa1 = append(aa1[:i], append([]int{x}, aa1[i:]...)...) //注意...的方式,类似于eachable的调用
	//在第i个位置插入切片
	aa1 = append(aa1[:i], append([]int{4, 5}, aa1[i:]...)...)

	/*
			切片复制 copy
			copy( destSlice, srcSlice []T) int
		其中 srcSlice 为数据来源切片，destSlice 为复制的目标（也就是将 srcSlice 复制到 destSlice），
		目标切片必须分配过空间且足够承载复制的元素个数，并且来源和目标的类型必须一致，copy() 函数的返回值表示实际发生复制的元素个数
	*/
	// 设置元素数量为1000
	const elementCount = 1000
	// 预分配足够多的元素切片
	srcData := make([]int, elementCount)
	// 将切片赋值
	for i := 0; i < elementCount; i++ {
		srcData[i] = i
	}
	// 引用切片数据
	refData := srcData
	// 预分配足够多的元素切片
	copyData := make([]int, elementCount)
	// 将数据复制到新的切片空间中
	copy(copyData, srcData)
	// 修改原始数据的第一个元素
	srcData[0] = 999
	// 打印引用切片的第一个元素
	fmt.Println(refData[0])
	// 打印复制切片的第一个和最后一个元素
	fmt.Println(copyData[0], copyData[elementCount-1])
	// 复制原始数据从4到6(不包含)
	copy(copyData, srcData[4:6])
	for i := 0; i < 5; i++ {
		fmt.Printf("%d ", copyData[i])
	}
}
