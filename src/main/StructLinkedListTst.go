package main

import "fmt"

type Node struct {
	Data int
	Next *Node
}

func ShowNode(p *Node) {
	for p != nil {
		fmt.Println(*p)
		p = p.Next
	}
}

func main() {
	var head = new(Node)
	head.Data = 0
	var tail *Node //tail用于记录头结点的地址，刚开始tail的的指针指向头结点
	tail = head

	/*
		插入节点 头插法和尾插法
	*/
	//头插法
	for i := 1; i < 10; i++ {
		var node = Node{
			Data: i,
			Next: nil,
		}
		node.Next = tail
		tail = &node
	}

	ShowNode(tail)

	var head1 = new(Node)
	head1.Data = 0
	tail = head1 //tail指向尾部节点的地址
	//尾插法
	for i := 1; i < 10; i++ {
		node := Node{
			Data: i,
			Next: nil,
		}
		tail.Next = &node
		tail = &node
	}
	ShowNode(head1)
}
