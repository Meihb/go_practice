package main

import (
	"fmt"
	"io"
	"os"
)

/*
类型断言的基本格式如下：
t := i.(T)

这里有两种可能。第一种，如果断言的类型 T 是一个具体类型，然后类型断言检查 i 的动态类型是否和 T 相同。
如果这个检查成功了，类型断言的结果是 i 的动态值，当然它的类型是 T。换句话说，具体类型的类型断言从它
的操作对象中获得具体的值。如果检查失败，接下来这个操作会抛出 panic。例如：
var w io.Writer
w = os.Stdout
f := w.(*os.File) // 成功: f == os.Stdout
c := w.(*bytes.Buffer) // 死机：接口保存*os.file，而不是*bytes.buffer

第二种，如果相反断言的类型 T 是一个接口类型，然后类型断言检查是否 i 的动态类型满足 T。如果这个检查成功了，
动态值没有获取到；这个结果仍然是一个有相同类型和值部分的接口值，但是结果有类型 T。换句话说，对一个接口类
型的类型断言改变了类型的表述方式，改变了可以获取的方法集合（通常更大），但是它保护了接口值内部的动态类型
和值的部分。

在下面的第一个类型断言后，w 和 rw 都持有 os.Stdout 因此它们每个有一个动态类型 *os.File，但是变量 w 是
一个 io.Writer 类型只对外公开出文件的 Write 方法，然而 rw 变量也只公开它的 Read 方法。
var w io.Writer
w = os.Stdout
rw := w.(io.ReadWriter) // 成功：*os.file具有读写功能
w = new(ByteCounter)
rw = w.(io.ReadWriter) // 死机：*字节计数器没有读取方法
*/

// 定义飞行动物接口
type Flyer interface {
	Fly()
}

// 定义行走动物接口
type Walker interface {
	Walk()
}

// 定义鸟类
type bird struct {
}

// 实现飞行动物接口
func (b *bird) Fly() {
	fmt.Println("bird: fly")
}

// 为鸟添加Walk()方法, 实现行走动物接口
func (b *bird) Walk() {
	fmt.Println("bird: walk")
}

// 定义猪
type pig struct {
}

// 为猪添加Walk()方法, 实现行走动物接口
func (p *pig) Walk() {
	fmt.Println("pig: walk")
}

//现金支付和alipay的例子
type ContainStolen interface {
	Stolen()
}
type ContainCanUseFaceID interface {
	UseFaceID()
}
type AliPay struct {
}

func (a *AliPay) UseFaceID() {
}

type CashPay struct {
}

func (c *CashPay) Stolen() {

}

func print(i interface{}) {
	switch i.(type) {
	case ContainCanUseFaceID:
		fmt.Printf("%T can use faceid\n", i)
	case ContainStolen:
		fmt.Printf("%T may be stolen\n", i)
	default:
		fmt.Println("nothing special")
	}
}

func main() {
	var w io.Writer
	w = os.Stdout
	f := w.(*os.File) // 成功: f == os.Stdout
	fmt.Println(f)
	//c := w.(*bytes.Buffer) // 死机：接口保存*os.file，而不是*bytes.buffer

	var w1 io.Writer
	w1 = os.Stdout
	rw1 := w1.(io.ReadWriter) // 成功：*os.file具有读写功能
	w1.Write([]byte{1, 2, 3})
	fmt.Println(rw1)
	rw1.Read([]byte{1, 2, 3}) //w1无法使用read方法,看出来了吗,通过断言可以更改方法集
	/*
		接口类型和接口 ,就是让面w1 和io.Writer的关系,接口类型可能实现了其他接口的方法,但是其本身只能显示
		调用此接口的方法集
		一个变量,其本身可能实现了多个方法,当其被赋值给了一个 接口,就变成了一个接口类型,其本身不在此接口的方法集就不能显示调用了
		但是通过断言,可以实现方法集的变化(通过断言assertion的第一个返回值)
	*/
	//w1 = new(ByteCounter) 编译报错
	//rw1 = w1.(io.ReadWriter) // 死机：*字节计数器没有读取方法

	// 创建动物的名字到实例的映射
	animals := map[string]interface{}{
		"bird": new(bird),
		"pig":  new(pig),
	}
	// 遍历映射
	for name, obj := range animals {
		// 判断对象是否为飞行动物
		f, isFlyer := obj.(Flyer)
		// 判断对象是否为行走动物
		w, isWalker := obj.(Walker)
		fmt.Printf("name: %s isFlyer: %v,addr:%p, isWalker: %v addr:%p\n", name, isFlyer,f, isWalker,w)
		// 如果是飞行动物则调用飞行动物接口
		if isFlyer {
			f.Fly()
		}
		// 如果是行走动物则调用行走动物接口
		if isWalker {
			w.Walk()
		}
	}

	p1 := new(pig)
	var a Walker = p1 //a是接口类型,因此只能支持Walker的方法集
	p2 := a.(*pig)    //通过断言,p2解放了其他方法的权限,但其实指针指向是一样的
	fmt.Printf("p1=%p p2=%p", p1, p2)

	//支付特点
	print(new(AliPay))
	print(new(CashPay))

	print(CashPay{})//哦豁,记得指针结构体的实现不能代表结构体的实现
}
