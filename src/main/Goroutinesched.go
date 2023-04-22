package main

import (
	"fmt"
	"runtime"
)



func say(s string) {

	for i := 0; i < 2; i++ {
		runtime.Gosched()//输出 hello world hello ,如果注释掉此句 则输出 hello helloo(当然是在单核cpu的情况下,
		// 多核会输出hello hello world world)
		//runtime.Gosched()用于让出CPU时间片。这就像跑接力赛，A跑了一会碰到代码runtime.Gosched()就把接力棒交给B了，
		// A歇着了，B继续跑。
		//类似于 php 协程的yield? 协程主动让出cpu 控制权
		fmt.Println(s)
	}

}



func main() {
	runtime.GOMAXPROCS(1)
	go say("world")
	say("hello")
}
