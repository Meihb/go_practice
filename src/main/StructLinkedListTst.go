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

func main()  {
	var head =new(Node)
	head.Data = 1
	var node1 = new(Node)
	node1.Data=2
	head.Next=node1

	var node2 =new(Node)
	node2.Data = 3
	node1.Next=node2

	ShowNode(head)
}