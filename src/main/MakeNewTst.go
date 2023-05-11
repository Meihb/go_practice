package main

import "fmt"

func main() {
	type Student struct {
		name string
		age  int
	}
	var s *Student
	s = new(Student) //分配空间
	s.name = "dequan"
	fmt.Println(s)

	/*
		var 或者简短模式的申明,其操作是创建内存,如果有初试值则赋值否则进行zerod 即零值化(0、0.0、nil等)
		new 分配内存,返回指针
		make 仅作用于channel、map、slice,分配内存后返回的即是类型本身而非指针类型,因为前三个类型都是引用类型

		make 只能用来分配及初始化类型为 slice、map、chanel 的数据。new 可以分配任意类型的数据；
		new 分配返回的是指针，即类型 *Type。make 返回引用，即 Type；
		new 分配的空间被清零。make 分配空间后，会进行初始化；
	*/

	var i int;
	i = 2
	fmt.Printf("int var value:%d\n", i)

	/*
			但是遇到指针会有问题
			打印的分别是p自身的存储地址、p存储的目标指针地址、p目标指针地址存储的内容
			如果不使用p=new(int),guess what happends after var p*int,即使说申明一个int指针类型p,那么由于没有显示初始化,p=nil
			这代表p指向的是一个nil，nil代表着系统还没有被分配地址,因此如果屏蔽此句就会报错,new语句则为覆盖了其Nil值为一个被分
			配过的内存地址,那就代表着*p的确能访问
			invalid memory address or nil pointer dereference
		不同类型的nil起始他们的地址都是一样的,但是不能进行==计算
	*/
	var p *int
	p = new(int)//runtime error: invalid memory address or nil pointer dereference ;var 好像并没有创建内存呀
	*p = 2
	fmt.Printf("p address:%p,p pointer:%p,p value:%d\n", &p, p, *p)//p address:0xc0000ba020,p pointer:0xc0000b6018,p value:2

	type makeP struct {
		name string
		id   int
	}
	var mp makeP
	mp.name = "12"
	fmt.Println(mp)

	mp2 := new(makeP)
	(*mp2).name = "22"
	fmt.Println(*mp2)

	//people := new(map[string]string)
	//*people["name"]="112"
	//p22 := *people
	//p22["name"] = "Kalan" // panic: assignment to entry in nil map

	var ma *(*makeP)
	ma = new(*makeP)//如果注释该行,将会报错 panic: runtime error: invalid memory address or nil pointer dereference 所以new 也分配了内存,申明变量并不能分配内存
	*ma = &makeP{};
	fmt.Println(*ma)
	fmt.Println(ma)
}
