package main

import (
	"fmt"
)

// Node node information
type Node struct {
	Left  *Node
	Data  int
	Right *Node
}

func create() *Node {
	var data int
	node := &Node{}
	fmt.Printf("enter data (-1 for no data): ")
	fmt.Scanf("%d\n", &data)
	fmt.Println(data)
	if data == -1 {
		return nil
	}
	node.Data = data
	fmt.Println("enter left child")
	node.Left = create()
	fmt.Println("enter right child")
	node.Right = create()
	return node
}

func main() {
	println("main")
	root := create()
	fmt.Printf("%+v\n", *root)
}
