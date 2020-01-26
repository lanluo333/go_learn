package tree

import "fmt"

type Node struct {
	Value int
	Left, Right *Node
}

func (node Node) Print(){
	fmt.Println(node.Value)
}

func (node *Node) SetValue(Value int){
	node.Value = Value
}

// 返回局部变量的函数
func CreateNode(Value int) *Node {
	return &Node{Value:Value}
}


