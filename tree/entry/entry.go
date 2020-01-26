package main

import (
	"fmt"
	"learn/tree"
)

type myTreeNode struct {
	node *tree.Node
}

func (myNode *myTreeNode) postOrder(){
	if myNode == nil || myNode.node == nil {
		return
	}

	left := myTreeNode{myNode.node.Left}
	right := myTreeNode{myNode.node.Right}
	left.postOrder()
	right.postOrder()
	myNode.node.Print()
}

func main() {
	var root tree.Node

	root = tree.Node{Value:3}
	root.Left = &tree.Node{}
	root.Right = &tree.Node{5, nil, nil}
	root.Right.Left = new(tree.Node)
	root.Left.Right = tree.CreateNode(2)
	root.Right.Left.SetValue(4)

	/*
		nodes := []Node{
			{ Value:7 },
			{},
			{ 6, nil, &root },
		}
		fmt.Println(nodes)

		root.right.left.setValue(9)
		root.right.left.print()
		fmt.Println()*/

	root.Traverse()

	//fmt.Println()
	//myTree := myTreeNode{&root}
	//myTree.postOrder()

	nodeCount := 0
	root.TraverseFunc(func(node *tree.Node) {
		nodeCount++
	})
	fmt.Println("Node Count:",nodeCount)

	c := root.TraverseWithChannel()
	maxNode := 0
	for node := range c{
		if node.Value > maxNode {
			maxNode = node.Value
		}
	}
	fmt.Println("Node MaxCount:",maxNode)
}
