package main

/*

sub、lea区别
LEA:“load effective address”的缩写，简单的说，lea指令可以用来将一个内存地址直接赋给目的操作数，
例如：lea eax,[ebx+8]就是将ebx+8这个值直接赋给eax，而不是把ebx+8处的内存地址里的数据赋给eax。
mov指令则恰恰相反，例如：mov eax,[ebx+8]则是把内存地址为ebx+8处的数据赋给eax。


//开辟栈空间，压栈 BP 保存现场
SUBQ $x, SP    //为函数开辟裁空间，栈底是0,栈顶是最大值,所以开辟空间是减法运算
MOVQ BP, y(SP) //保存当前函数 BP 到 y(SP）位直， y 为相对 SP 的偏移量
LEAQ y(SP), BP //重直 BP，使其指向刚刚保存 BP 旧值的位置，这里主要是方便后续 BP 的恢复

//弹出栈，恢复 BP
MOVQ y(SP), BP //恢复 BP 的值为调用前的值
ADDQ $x, SP    //恢复 SP 的值为函数开始时的位



bp、sp区别
SS, SP, BP 三个寄存器

SS:存放栈的段地址；
SP:堆栈寄存器SP(stack pointer)存放栈的偏移地址;
BP: 基数指针寄存器BP(base pointer)是一个寄存器，它的用途有点特殊，是和堆栈指针SP联合使用的，作为SP校准使用的，只有在寻找堆栈里的数据和使
用个别的寻址方式时候才能用到
比如说，堆栈中压入了很多数据或者地址，你肯定想通过SP来访问这些数据或者地址，但SP是要指向栈顶的，是不能随便乱改的，这时候你就需要使用BP，
把SP的值传递给BP，通过BP来寻找堆栈里数据或者地址．一般除了保存数据外,可以作为指针寄存器用于存储器寻址,此时它默认搭配的段寄存器是SS-堆栈段
寄存器.BP是16位的,再扩充16位就是EBP,用于32位编程环境的.一般高级语言的参数传递等等,转换为汇编后经常由BP/EBP来负责寻址\处理.
SP,BP一般与段寄存器SS 联用，以确定堆栈寄存器中某一单元的地址，SP用以指示栈顶的偏移地址，而BP可 作为堆栈区中的一个基地址，用以确定在堆栈中的操作数地址。

(下面这个像Win32汇编中的)
bp为基址寄存器，一般在函数中用来保存进入函数时的sp的栈顶基址
每次子函数调用时，系统在开始时都会保存这个两个指针并在函数结束时恢复sp和bp的值。像下面这样：
在函数进入时：
push bp     // 保存bp指针
mov bp,sp  // 将sp指针传给bp，此时bp指向sp的基地址。
                  // 这个时候，如果该函数有参数，则[bp + 2*4]则是该子函数的第一个参数，[bp+3*4]则是该子函数的 第二个参数，以此类推，有多少个参数则[bp+(n-1)*4]。
.....
.....
函数结束时：
mov sp,bp  // 将原sp指针传回给sp
pop bp       // 恢复原bp的值。
ret              // 退出子函数



Go 编译器产生的汇编代码是一种中间抽象态，它不是对机器码的映射，而是和平台无关的一个中间态汇编描述，所以汇编代码中有些寄存器是真实的，有些是抽象的，几个抽象的寄存器如下：
SB (Static base pointer)：静态基址寄存器，它和全局符号一起表示全局变量的地址。
FP (Frame pointer)：栈帧寄存器，该寄存器指向当前函数调用栈帧的栈底位置。 当前函数即主调函数 栈底
PC (Program counter)：程序计数器，存放下一条指令的执行地址，很少直接操作该寄存器，一般是 CALL、RET 等指令隐式的操作。
SP (Stack pointer)：栈顶寄存器，一般在函数调用前由主调函数设置 SP 的值对栈空间进行分配或回收。 栈顶


*/

/*
package main
func swap (a, b int) (x int, y int) {
    x = b
    y = a
    return
}
func main() {
    swap(10, 20)
}

//- S 产生汇编的代码
//- N 禁用优化
//- 1 禁用内联

GOOS=linux GOARCH=amd64 go tool compile -1 -N -S swap.go >swap.s 2>&1

所以所以,movq 好像和Mov不一样,not mov dest src,而是Movq src dest!!!!

.functionName
"".swap STEXT nosplit size=39 args=0x20 locals=0x0
    0x0000 00000 (swap.go:4) TEXT  "".swap(SB), NOSPLIT, $0 - 32
    0x0000 00000 (swap.go:4) FUNCDATA  $0, gclocals.ff19ed39bdde8a01a800918ac3ef0ec7(SB)
    0x0000 00000 (swap.go:4) FUNCDATA  $1, gclocals.33cdeccccebe80329flfdbee7f5874cb(SB)
	初始化返回值 x 为 0 为什么是 x+24呢,也许是因为args=0x20哦,不对(其实是对的啊,0x20就是32啊),是因为有两个返回值在栈顶,其次才是入参,
	所以上面TEXT才有$0-32,32表示32bit长度的返回值和入参长度
    0x0000 00000 (swap.go:4)  MOVQ  $0, "".x+24(SP)
	初始化返回值 y 为 0
    0x0009 00009 (swap.go:4)  MOVQ  $0, "".y+32(SP)
	取第 2 个参数赋值给返回值 x
    0x0012 00018 (swap.go:5)  MOVQ  "".b+16(SP), AX
    0x0017 00023 (swap.go:5)  MOVQ  AX, "".x+24(SP)
	取第 1 个参数赋值给返回值 y
    0xOO1c 00028 (swap.go:6)  MOVQ  "".a+8(SP), AX
    0x0021 00033 (swap.go:6)  MOVQ  AX, "".y+32(SP)
	函数返回，同时进行栈回收，FUNCDATA 和垃圾收集可以忽略。
    0x0026 00038 (swap.go:7)  RET



	 main 函数堆栈初始化：开辟栈空间，保存 BP 寄存器。
"".main STEXT size=68 args=0x0 locals=0x28
    0x0000 00000 (swap.go:10) TEXT "".main(SB), $40 - 0
    0x0000 00000 (swap.go:10) MOVQ (TLS), CX
    0x0009 00009 (swap.go:10) CMPQ SP, 16(CX)
    0x000d 00013 (swap.go:10) JLS 61
    0x000f 00015 (swap.go:10) SUBQ $40, SP
    0x0013 00019 (swap.go:10) MOVQ BP, 32 (SP)
    0x0018 00024 (swap.go:10) LEAQ 32(SP), BP
    0x001d 00029 (swap.go:10) FUNCDATA $0, gclocals ·33cdeccccebe80329flfdbee7f5874cb(SB)
    0x001d 00029 (swap.go:10) FUNCDATA $1, gclocals ·33cdeccccebe80329flfdbee7f5874cb(SB)
	初始化 add 函数的调用参数 1 的值为 10
    0x001d 00029 (swap.go:11) MOVQ $10, (SP)
	初始化 add 函数的调用参数 2 的值为 20
    0x0025 00037 (swap.go:11) MOVQ $20 , 8 (SP)
    0x002e 00046 (swap.go:11) PCDATA $0 , $0
	调用 swap 函数，注意 call 隐含一个将 swap 下一条指令地址压栈的动作，即 sp=sp+8
    0x002e 00046 (swap.go:11) CALL "". swap(SB)
    0x0033 00051 (swap.go:12) MOVQ 32(SP), BP
	恢复措空间
    0x0038 00056 (swap.go:12) ADDQ $40, SP
    0x003c 00060 (swap.go:12) RET
    0x003d 00061 (swap.go:12) NOP
    0x003d 00061 (swap.go:10) PCDATA $0, $ - 1

从汇编的代码得知：
函数的调用者负责环境准备，包括为参数和返回值开辟栈空间。
寄存器的保存和恢复也由调用方负责。
函数调用后回收栈空间，恢复 BP 也由主调函数负责。

*/

/*
源码
package main


func g(p int) int {
     return p+1;
}

func main() {
     c := g(4) + 1
     _ = c
}

GOOS=linux GOARCH=386 go tool compile -S main.go >> main.S

"".g t=1 size=16 value=0 args=0x10 locals=0x0
     0x0000 00000 (main.go:4)     TEXT     "".g(SB), $0-16
     0x0000 00000 (main.go:4)     NOP
     0x0000 00000 (main.go:4)     NOP
     0x0000 00000 (main.go:4)     FUNCDATA     $0, gclocals·23e8278e2b69a3a75fa59b23c49ed6ad(SB)
     0x0000 00000 (main.go:4)     FUNCDATA     $1, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
     0x0000 00000 (main.go:5)     MOVQ     "".p+8(FP), BX
     0x0005 00005 (main.go:5)     INCQ     BX
     0x0008 00008 (main.go:5)     MOVQ     BX, "".~r1+16(FP)
     0x000d 00013 (main.go:5)     RET
     0x0000 48 8b 5c 24 08 48 ff c3 48 89 5c 24 10 c3        H.\$.H..H.\$..
"".main t=1 size=16 value=0 args=0x0 locals=0x0
     0x0000 00000 (main.go:8)     TEXT     "".main(SB), $0-0
     0x0000 00000 (main.go:8)     NOP
     0x0000 00000 (main.go:8)     NOP
     0x0000 00000 (main.go:8)     FUNCDATA     $0, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
     0x0000 00000 (main.go:8)     FUNCDATA     $1, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
     0x0000 00000 (main.go:9)     MOVQ     $4, BX
     0x0007 00007 (main.go:9)     INCQ     BX
     0x000a 00010 (main.go:9)     INCQ     BX
     0x000d 00013 (main.go:11)     RET
     0x0000 48 c7 c3 04 00 00 00 48 ff c3 48 ff c3 c3        H......H..H...
"".init t=1 size=80 value=0 args=0x0 locals=0x0
     0x0000 00000 (main.go:11)     TEXT     "".init(SB), $0-0
     0x0000 00000 (main.go:11)     MOVQ     (TLS), CX
     0x0009 00009 (main.go:11)     CMPQ     SP, 16(CX)
     0x000d 00013 (main.go:11)     JLS     62
     0x000f 00015 (main.go:11)     NOP
     0x000f 00015 (main.go:11)     NOP
     0x000f 00015 (main.go:11)     FUNCDATA     $0, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
     0x000f 00015 (main.go:11)     FUNCDATA     $1, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
     0x000f 00015 (main.go:11)     MOVBQZX     "".initdone·(SB), BX
     0x0016 00022 (main.go:11)     CMPB     BL, $0
     0x0019 00025 (main.go:11)     JEQ     47
     0x001b 00027 (main.go:11)     MOVBQZX     "".initdone·(SB), BX
     0x0022 00034 (main.go:11)     CMPB     BL, $2
     0x0025 00037 (main.go:11)     JNE     40
     0x0027 00039 (main.go:11)     RET
     0x0028 00040 (main.go:11)     PCDATA     $0, $0
     0x0028 00040 (main.go:11)     CALL     runtime.throwinit(SB)
     0x002d 00045 (main.go:11)     UNDEF
     0x002f 00047 (main.go:11)     MOVB     $1, "".initdone·(SB)
     0x0036 00054 (main.go:11)     MOVB     $2, "".initdone·(SB)
     0x003d 00061 (main.go:11)     RET
     0x003e 00062 (main.go:11)     CALL     runtime.morestack_noctxt(SB)
     0x0043 00067 (main.go:11)     JMP     0

Analysis:
TEXT定义函数,分三个部分
0x0000 00000 (main.go:4)     TEXT     "".g(SB), $0-16
0x0000 00000 (main.go:8)     TEXT     "".main(SB), $0-0
0x0000 00000 (main.go:11)     TEXT     "".init(SB), $0-0
这个"". 代表的是这个函数的命名空间。
g(SB) 这里就有个SB的伪寄存器。全名未Static Base，代表g这个函数地址，0−16中的0代表局部变量字节数总和，0表示不存在局部变量。-16代表在0的
地址基础上空出16的长度作为传入和返回对象。这个也就是golang如何实现函数的多返回值的方法了。它在定义函数的时候，开辟了一定空间存储传入和传
出对象。

FUNCDATA是golang编译器自带的指令，plan9和x86的指令集都是没有的。它用来给gc收集进行提示。提示0和1是用于局部函数调用参数，需要进行
回收。

	0x0000 00000 (main.go:5)     MOVQ     "".p+8(FP), BX
 	0x0005 00005 (main.go:5)     INCQ     BX
	0x0008 00008 (main.go:5)     MOVQ     BX, "".~r1+16(FP)
这里有一个FP寄存器，FP是frame pointer，是指向栈底，SP是指向栈顶。BX是一个临时寄存器，那么上面的句子是代表把FP＋8这个位置的数据（参数p），
保存到BX中。FP+8代表的是参数
INCQ是自增算法，BX里面的数自加1，然后把BX里面的数存储到FP+16，代表的是返回值。
下面就是RET，直接返回。


*/
func main() {
}
