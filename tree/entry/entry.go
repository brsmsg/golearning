package main

import (
	"fmt"
	"golearning/tree"
)

// compose
type MyTreeNode struct {
	node *tree.Node
}

func (myNode *MyTreeNode) postOrder() {
	if myNode == nil || myNode.node == nil {
		return
	}
	left := MyTreeNode{myNode.node.Left}
	left.postOrder()
	right := MyTreeNode{myNode.node.Right}
	right.postOrder()
	myNode.node.Print()
}

func main() {
	var root tree.Node
	root = tree.Node{Value: 3}
	root.Left = &tree.Node{}
	root.Right = &tree.Node{5, nil, nil}
	root.Right.Left = new(tree.Node)

	root.Left.Right = tree.CreateNode(2)

	nodes := []tree.Node{
		{Value: 3},
		{},
		{6, nil, &root},
	}
	fmt.Println(nodes)
	root.Print()
	root.Right.Left.SetValue(4)
	root.Right.Left.Print()
	// print(root)

	var pRoot *tree.Node
	// pRoot.print()
	pRoot.SetValue(200)
	pRoot = &root
	pRoot.SetValue(300)
	pRoot.Print()

	pRoot.Traverse()
	fmt.Println()
	myRoot := MyTreeNode{&root}
	myRoot.postOrder()
}
