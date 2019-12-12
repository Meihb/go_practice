package demo

import "testing"

/*
  不能通过package,必须直接到文件夹地址

编写测试用例有以下几点需要注意：
测试用例文件不会参与正常源码的编译，不会被包含到可执行文件中；
测试用例的文件名必须以_test.go结尾；
需要使用 import 导入 testing 包；
测试函数的名称要以Test或Benchmark开头，后面可以跟任意字母组成的字符串，但第一个字母必须大写，例如 TestAbc()，一个测试用例文件中可以包含多个测试函数；
单元测试则以(t *testing.T)作为参数，性能测试以(t *testing.B)做为参数；
测试用例文件使用go test命令来执行，源码中不需要 main() 函数作为入口，所有以_test.go结尾的源码文件内以Test开头的函数都会自动执行。

Go语言的 testing 包提供了三种测试方式，分别是单元（功能）测试、性能（压力）测试和覆盖率测试。
*/

/*
单元测试 功能测试
go test -v

=== RUN   TestGetArea
--- PASS: TestGetArea (0.00s)
PASS
ok      demo    0.204s

*/
func TestGetArea(t *testing.T) {
	area := GetArea(40, 50)
	if area != 2000 {
		t.Error("测试失败")
	}
}

/*
性能测试 即压力测试
go test -bench="."

goos: windows
goarch: amd64
pkg: demo
BenchmarkGetArea-8      1000000000               0.307 ns/op
PASS
ok      demo    0.525s

*/
func BenchmarkGetArea(t *testing.B) {
	for i := 0; i < t.N; i++ {
		GetArea(40, 50)
	}
}

/*
覆盖率测试
覆盖率测试能知道测试程序总共覆盖了多少业务代码（也就是 demo_test.go 中测试了多少 demo.go 中的代码），可以的话最好是覆盖100%。
go test -cover
PASS
coverage: 100.0% of statements
ok      demo    0.145s

*/
