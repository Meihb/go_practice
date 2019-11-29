package main

import "fmt"

func main() {

Outerloop:
	for x := 0; x < 3; x++ {
		fmt.Printf("x:%d\n", x)
		for y := 0; y < 5; y++ {
			fmt.Printf("\t\ty:%d\n", y)
			switch {
			case y == 2 && x == 0:
				fmt.Println(x, y)
				continue Outerloop
			case y == 3:
				fmt.Println(x, y)
				break Outerloop
			}
		}
	}
}
