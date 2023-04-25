package main

import (
	"fmt"
)

// s :="string1";//简短申明只可用在函数内
func main() {

	a, b := 1, 2
	fmt.Println(a, b)
	a, b = 2, 1  //这不是申明变量,而是变量赋值
	a, c := 1, 3 //简短申明可以批量,但至少有一个是申明,若全是赋值则报错
	c = c + 1
	//简短申明 只有在相同词法域申明过的变量才和赋值等价,如果该变量之前在外部词法域申明,则简短申明将会在当前词法域申明一个新变量

	var i int = 0
	fmt.Println(i)
	for i := 2; i < 4; i++ {
		fmt.Println(i)
	}
	fmt.Println(i) //0 也就是说如果在次级词法域中简短申明一个在外部词法域申明过的变量,他不是赋值操作而是申明操作,生命周期在次级词法域中 
}
