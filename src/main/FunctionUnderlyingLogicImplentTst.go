package main

/*

sub
//开辟栈空间，压栈 BP 保存现场
SUBQ $x, SP    //为函数开辟裁空间
MOVQ BP, y(SP) //保存当前函数 BP 到 y(SP）位直， y 为相对 SP 的偏移量
LEAQ y(SP), BP //重直 BP，使其指向刚刚保存 BP 旧值的位置，这里主要是方便后续 BP 的恢复

//弹出栈，恢复 BP
MOVQ y(SP), BP //恢复 BP 的值为调用前的值
ADDQ $x, SP    //恢复 SP 的值为函数开始时的位
*/

/*

Go 编译器产生的汇编代码是一种中间抽象态，它不是对机器码的映射，而是和平台无关的一个中间态汇编描述，所以汇编代码中有些寄存器是真实的，有些是抽象的，几个抽象的寄存器如下：
SB (Static base pointer)：静态基址寄存器，它和全局符号一起表示全局变量的地址。
FP (Frame pointer)：栈帧寄存器，该寄存器指向当前函数调用栈帧的栈底位置。 栈底
PC (Program counter)：程序计数器，存放下一条指令的执行地址，很少直接操作该寄存器，一般是 CALL、RET 等指令隐式的操作。
SP (Stack pointer)：栈顶寄存器，一般在函数调用前由主调函数设置 SP 的值对栈空间进行分配或回收。 栈顶



"".swap STEXT nosplit size=39 args=0x20 locals=0x0
    0x0000 00000 (swap.go:4) TEXT  "".swap(SB), NOSPLIT, $0 - 32
    0x0000 00000 (swap.go:4) FUNCDATA  $0, gclocals.ff19ed39bdde8a01a800918ac3ef0ec7(SB)
    0x0000 00000 (swap.go:4) FUNCDATA  $1, gclocals.33cdeccccebe80329flfdbee7f5874cb(SB)
	初始化返回值 x 为 0
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

func main() {
}
