package main

import (
	"errors"
	"fmt"
	"math"
)

/*
// The error built-in interface type is the conventional interface for
// representing an error condition, with the nil value representing no error.
type error interface {
    Error() string
}

Error接口其实很简单,一个Error()string的签名函数实现之后,你的自定义类型就可以作为error类型返回值了
*/


func Sqrt(f float64) (float64, error) {
	if f < 0 {
		//简单错误类型返回
		return -1, errors.New("math: square root of negative number")
	}
	return math.Sqrt(f), nil
}

/*
自定义错误类型
*/
type dualError struct {
	Num     float64
	problem string
}

func (e dualError) Error() string {
	return fmt.Sprintf("Wrong!!!,because  \"%\f%\"is a negative number ", e.Num)
}
func Sqrt2(f float64) (float64, error) {
	if f < 0 {
		return -1, dualError{Num: f}
	}
	return math.Sqrt(f), nil
}

func main() {
	result, err := Sqrt(-13)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}

	result1, err1 := Sqrt2(-13)
	if err1 != nil {
		fmt.Println(err1)
	} else {
		fmt.Println(result1)
	}
}
