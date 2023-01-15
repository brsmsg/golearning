package tree

import "fmt"

type Node struct {
	Value       int
	Left, Right *Node
}

// 所有参数传值
func (node Node) Print() {
	fmt.Println(node.Value)
}

func (node *Node) Traverse() {
	if node == nil {
		return
	}
	node.Left.Traverse()
	node.Print()
	node.Right.Traverse()
}

func CreateNode(value int) *Node {
	return &Node{Value: value}
}

// 使用指针才能改变结构内容
func (node *Node) SetValue(value int) {
	if node == nil {
		fmt.Println("setting value to nil")
		return
	}
	node.Value = value
}

// 改变内容 / 结构过大 使用指针接收者
// 建议一律使用指针接收者
