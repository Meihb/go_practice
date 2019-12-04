package main

import (
	"fmt"
	"time"
)

func fibonacci(n int) (res int) {
	if n <= 2 {
		return 1
	} else {
		return fibonacci(n-1) + fibonacci(n-2)
	}

}

const LIM = 41

var fibs [LIM]uint64

func fibonacciCache(n int) (res uint64) {
	//记忆化,检查数组中是否已知fibonacci(n)
	if fibs[n] != 0 {
		res = fibs[n]
	} else {
		if n <= 2 {
			res = 1
		} else {
			res = fibonacciCache(n-1) + fibonacciCache(n-2)
		}
		fibs[n] = res
	}
	return
}
func main() {
	result := 0
	start := time.Now()
	for i := 1; i <= 40; i++ {
		result = fibonacci(i)
		fmt.Printf("数列第%d位:%d\n", i, result)
	}
	end := time.Now()
	delta := end.Sub(start)
	fmt.Printf("程序的执行时间为:%s\n", delta)

	var uresult uint64 = 0
	start = time.Now()
	for i := 0; i <= 40; i++ {
		uresult = fibonacciCache(i)
		fmt.Printf("数列第%d位:%d\n", i, uresult)
	}
	end = time.Now()
	delta = end.Sub(start)
	fmt.Printf("缓存化程序执行时间为:%s\n", delta)

}
