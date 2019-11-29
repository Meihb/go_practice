package main

import "fmt"

func main() {
	arr := [...]int{21, 32, 12, 33, 34, 34, 87, 24}
	var n = len(arr)

	for i := 0; i < n; i++ {
		fmt.Printf("第%d次循环\n", i)
		for j := n - 1; j > i; j-- {
			fmt.Printf("\t\t比较 %d 和 %d\n", arr[j], arr[j-1])
			if (arr[j] > arr[j-1]) {
				tmp := arr[j]
				arr[j] = arr[j-1]
				arr[j-1] = tmp
			}
		}
		fmt.Println(arr)
	}
}
