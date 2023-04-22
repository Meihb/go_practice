package main

import (
	"flag"
	"fmt"
)

var skillParam = flag.String("skill", "fly", "skill to perform")

func main() {

	/*
	申明时即可调用,thats the suspicious ()after the body
		或者复制给变量,从而调用
	*/
	func(data int) {
		fmt.Printf("Data:%d\n", data)
	}(100)

	f := func(data int) { fmt.Printf("Data:%d\n", data) }
	f(120)

	flag.Parse()
	fmt.Println(skillParam)

	var skill = map[string]func(){
		"fire": func() {
			fmt.Println("chicken fire")
		},
		"run": func() {
			fmt.Println("soldier run")
		},
		"fly": func() {
			fmt.Println("angel fly")
		},
	}
	if f, ok := skill[*skillParam]; ok {
		f()
	} else {
		fmt.Println("skill not found")
	}
}