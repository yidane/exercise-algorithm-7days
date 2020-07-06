package main

type Node struct {
	Val      int
	Children []*Node
}

func preorder(root *Node) []int {
	if root == nil {
		return nil
	}

	return appendArr(root)
}

func appendArr(node *Node) []int {
	if node == nil {
		return nil
	}

	var arr []int
	arr = append(arr, node.Val)
	for i := 0; i < len(node.Children); i++ {
		arr = append(arr, appendArr(node.Children[i])...)
	}

	return arr
}