package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main(){
	counts := make(map[string]int)

	input:=bufio.NewScanner(os.Stdin)  
	input.Split(bufio.ScanWords)
	for input.Scan() {
		word:=input.Text()
		if _,ok:=counts[word];!ok{
			counts[word] = 1
		}else{ 
			counts[word]++
		}
		fmt.Println(word)
	} 
	fmt.Println(counts)
}