package main

import (
	"flag"
	"fmt"
)

func main()  {
	name := flag.String("name", "张三", "姓名")
	age := flag.Int("age", 18, "年龄")
	married := flag.Bool("married", false, "婚否")
	delay := flag.Duration("d", 0, "时间间隔")

	fmt.Println(name,age,married,delay)
}