package main

import (
	"errors"
	"fmt"
)

/*
type error interface {
    Error() string
}
自定义错误
*/
var err = errors.New("this is an error")

type ParseError struct {
	FileName string //文件名
	Line     int
}

// 实现error接口，返回错误描述
func (e ParseError) Error() ( string) {
	return fmt.Sprintf("%s:%d", e.FileName,  e.Line)
}

//创建解析错误的函数
func newParseError(filename string, line int) error {
	return ParseError{filename, line}
}

func main() {
	var e error
	//创建一个错误实例
	e = newParseError("main.go", 1)
	//通过error接口查询错误描述
	fmt.Println(e.Error())
	//根据错误已接口具体的类型,获取详细错误信息
	switch detail := e.(type) {
	case ParseError: //这是一个解析错误
		fmt.Printf("Filename:%s Line:%d\n", detail.FileName, detail.Line)
	default:
		fmt.Printf("other error")
	}
}
