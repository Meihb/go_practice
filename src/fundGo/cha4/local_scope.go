package main

import "fmt"

var a = "G"

func main() {
	n()
	m1()
	//    m2()
	n()

	var ch int = '\u0041'
	var ch2 int = '\u03B2'
	var ch3 int = '\U00101234'
	fmt.Printf("%d - %d - %d\n", ch, ch2, ch3) // integer
	fmt.Printf("%c - %c - %c\n", ch, ch2, ch3) // character
	fmt.Printf("%X - %X - %X\n", ch, ch2, ch3) // UTF-8 bytes
	fmt.Printf("%U - %U - %U\n", ch, ch2, ch3)   // UTF-8 code point 



}

func n() { fmt.Println(a) }

func m1() {
	a := "O"
	fmt.Println(a)
}

func m2() int{
	a = "O"
	fmt.Println(a)

	x:=100
	if(x<10){
		return 1
	}else if x==10 {
		return 2
	}else{
		return 3
	}
}
