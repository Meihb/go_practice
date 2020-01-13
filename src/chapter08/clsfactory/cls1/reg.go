package cls1
import (
	base2 "chapter08/clsfactory/base"
	"fmt"
)
// 定义类1
type Class1 struct {
}
// 实现Class接口
func (c *Class1) Do() {
	fmt.Println("Class1")
}
func init() {
	// 在启动时注册类1工厂
	base2.Register("Class1", func() base2.Class {
		return new(Class1)
	})
}