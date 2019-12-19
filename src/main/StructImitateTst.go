package main

import "fmt"

//可飞行类
type Flying struct {
}

func (f *Flying) Fly() {
	fmt.Println("can fly")
}

//行走类
type Walkable struct {
}

func (w *Walkable) Walk() {
	fmt.Println("can walk")
}

//人类
type Human struct {
	Walkable
}
type Bird struct {
	Walkable
	Flying
}

func main() {
	// 实例化鸟类
	b := new(Bird)
	fmt.Println("Bird: ")
	b.Flying.Fly()
	b.Fly()
	b.Walk()
	// 实例化人类
	h := new(Human)
	fmt.Println("Human: ")
	h.Walk()

}
