package base
// 类接口
type Class interface {
	Do()
}
var (
	// 保存注册好的工厂信息
	factoryByName = make(map[string]func() Class)
)
// 注册一个类生成工厂
func Register(name string, factory func() Class) {//哇,其实是一个返回Class的func 想错了 go的函数签名,我这个phper还适应的不太行
	factoryByName[name] = factory
}
// 根据名称创建对应的类
func Create(name string) Class {
	if f, ok := factoryByName[name]; ok {
		return f()
	} else {
		panic("name not found")
	}
}