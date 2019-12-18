package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"
)

type MyInt int

func (m MyInt) IsZero() bool {
	return m == 0
}

func (m MyInt) MyAdd(other int) int {
	return other + int(m)
}

func main() {
	var b MyInt
	fmt.Println(b.IsZero())
	b = 1
	fmt.Println(b.MyAdd(2))

	//http包
	client := &http.Client{}

	//创建http请求
	req, err := http.NewRequest("POST", "http://www.baidu.com/", strings.NewReader("key=value"))
	//发现错误则打印并且推出
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
		return
	}
	//为header添加消息
	req.Header.Add("User-Agent", "myClient")

	//开始请求
	resp, err := client.Do(req)
	//发现错误则打印并且推出
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
		return
	}
	data,err :=ioutil.ReadAll(resp.Body)
	fmt.Println(string(data))
	defer resp.Body.Close()

	fmt.Println(time.Second.String())
}
