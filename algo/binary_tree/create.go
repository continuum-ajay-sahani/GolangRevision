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

func inorder(root *Node) {
	if root == nil {
		return
	}
	inorder(root.Left)
	fmt.Print(root.Data)
	inorder(root.Right)
}

func preorder(root *Node) {
	if root == nil {
		return
	}
	fmt.Print(root.Data)
	preorder(root.Left)
	preorder(root.Right)
}

func postorder(root *Node) {
	if root == nil {
		return
	}
	postorder(root.Left)
	postorder(root.Right)
	fmt.Print(root.Data)
}

func main() {
	println("main")
	root := create()
	fmt.Printf("%+v\n", *root)
	fmt.Println("-------------------")
	fmt.Println("----Inorder----")
	inorder(root)
	fmt.Println("----Preorder----")
	preorder(root)
	fmt.Println("---Postorder----")
	postorder(root)
	fmt.Println("-------------------")
}
