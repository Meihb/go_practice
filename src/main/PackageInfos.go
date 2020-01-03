package main

import F "fmt" //引用别名 use namespace as
import . "net/http"//省略引用
import  _ "net/http/httptest"//匿名引用 只执行包初始化的init函数

func main() {
	F.Println("hehe")
	resp, error := Head("http://baidu.com")
	if error != nil {
		F.Println(error)
	}
	println(resp)
}
