package main

import (
	"io"
)

/*
type Writer interface {
	Write(p []byte) (n int, err error)
}
type Closer interface {
	Close() error
}
type WriteCloser interface {
	Writer//看起来就像是匿名结构体,应该是如果不是函数签名,就找寻同名interface
	Closer
}
*/
// 声明一个设备结构
type device struct {
}

// 实现io.Writer的Write()方法
func (d *device) Write(p []byte) (n int, err error) {
	return 0, nil
}

// 实现io.Closer的Close()方法
func (d *device) Close() error {
	return nil
}
func main() {
	// 声明写入关闭器, 并赋予device的实例
	var wc io.WriteCloser = new(device)
	// 写入数据
	wc.Write(nil)
	// 关闭设备
	wc.Close()
	// 声明写入器, 并赋予device的新实例
	var writeOnly io.Writer = new(device)
	// 写入数据
	writeOnly.Write(nil)
}
