/*
Go语言标准库的 bufio 包中，实现了对数据 I/O 接口的缓冲功能。这些功能封装于接口
io.ReadWriter、io.Reader 和 io.Writer 中，
并对应创建了 ReadWriter、Reader 或 Writer 对象，在提供缓冲的同时实现了一些文本基本 I/O 操作功能。
*/
package main

import (
	"bufio"
	"bytes"
	"fmt"
)

func main() {
	data := []byte("Go语言入门教程")

	rd := bytes.NewReader(data)
	//NewReader() 函数的功能是按照缓冲区默认长度创建 Reader 对象，
	// Reader 对象会从底层 io.Reader 接口读取尽量多的数据进行缓存
	r := bufio.NewReader(rd)
	var buf [14]byte

	//Read 读取数据并存放到参数切片中,
	n, err := r.Read(buf[:]) //读取并存放
	fmt.Println(r.Buffered())
	fmt.Println(string(buf[:n]), n, err)

	//读取单个字节
	b, err := r.ReadByte()
	fmt.Println(b, err)


	var delim byte = ','
	//delimiter 参数注意一下,可以实现readline by delimiter \n
	line, err := r.ReadBytes(delim)
	fmt.Println(string(line), err)

	line, prefix, err := r.ReadLine()
	fmt.Println(string(line), prefix, err)


	/*
	writer
	 */
	wr := bytes.NewBuffer(nil)
	w := bufio.NewWriter(wr)
	p := []byte("C语言中文网")
	fmt.Println("写入前未使用的缓冲区为：", w.Available())
	w.Write(p)
	fmt.Printf("写入%q后，未使用的缓冲区为：%d\n", string(p), w.Available())
	fmt.Println("当前缓冲区:",w.Buffered())

	w.Flush()
	fmt.Printf("执行 Flush 后缓冲区输出 %q\n", string(wr.Bytes()))

}
