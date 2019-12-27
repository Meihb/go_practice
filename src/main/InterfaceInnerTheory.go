package main

/*
非空接口底层iface
//src/runtime/runtime2.go
type iface struct {
    tab *itab                //itab 存放类型及方法指针信息
    data unsafe.Pointer      //数据信息
}
itab：用来存放接口自身类型和绑定的实例类型及实例相关的函数指针，具体内容后面有详细介绍。
数据指针 data：指向接口绑定的实例的副本，接口的初始化也是一种值拷贝。

itab:
//src/runtime/runtime2.go
type itab struct {
    inter *interfacetype      //接口自身的静态类型
    _type *_type              //_type 就是接口存放的具体实例的类型（动态类型）
    //hash 存放具体类型的 Hash 值
    hash uint32               // copy of _type.hash. Used for type switches.
    _   [4]byte
    fun [1]uintptr            // variable sized. fun[0]==0 means _type does not implement inter.
}
inner：是指向接口类型元信息的指针。
_type：是指向接口存放的具体类型元信息的指针，iface 里的 data 指针指向的是该类型的值。一个是类型信息，
		另一个是类型的值。
hash：是具体类型的 Hash 值，_type 里面也有 hash，这里冗余存放主要是为了接口断言或类型查询时快速访问。
fun：是一个函数指针，可以理解为 C++ 对象模型里面的虚拟函数指针，这里虽然只有一个元素，实际上指针数组
		的大小是可变的，编译器负责填充，运行时使用底层指针进行访问，不会受 struct 类型越界检查的约束，这些指针指向的是具体类型的方法。

type
//src/runtime/type.go
type type struct {
    size uintptr     // 大小
    ptrdata uintptr  //size of memory prefix holding all pointers
    hash uint32      //类型Hash
    tflag tflag      //类型的特征标记
    align uint8      //_type 作为整体交量存放时的对齐字节数
    fieldalign uint8 //当前结构字段的对齐字节数
    kind uint8       //基础类型枚举值和反射中的 Kind 一致，kind 决定了如何解析该类型
    alg *typeAlg     //指向一个函数指针表，该表有两个函数，一个是计算类型 Hash 函
                     //数，另一个是比较两个类型是否相同的 equal 函数
    //gcdata stores the GC type data for the garbage collector.
    //If the KindGCProg bit is set in kind, gcdata is a GC program.
    //Otherwise it is a ptrmask bitmap. See mbitmap.go for details.
    gcdata *byte      //GC 相关信息
    str nameOff       //str 用来表示类型名称字符串在编译后二进制文件中某个 section
                      //的偏移量
                      //由链接器负责填充
    ptrToThis typeOff //ptrToThis 用来表示类型元信息的指针在编译后二进制文件中某个
                      //section 的偏移量
                      //由链接器负责填充
}
*/

type Caler interface {
	Add(a, b int) int
	Sub(a, b int) int
}
type Adder struct{ id int }

//go:noinline
func (adder Adder) Add(a, b int) int { return a + b }

//go:noinline
func (adder Adder) Sub(a, b int) int { return a - b }
func main() {
	var m Caler = Adder{id: 1234}
	m.Add(10, 32)
}
