package main

import (
	"fmt"

	datastructures "github.com/cecarbs/dsa/internal/data-structures"
)

func main() {
	root := datastructures.TreeNode{Val: 1}
	root.Left = &datastructures.TreeNode{Val: 2}
	root.Right = &datastructures.TreeNode{Val: 3}
	fmt.Printf("Val: %d, Left: %v, Right: %v", root.Val, root.Left.Val, root.Right.Val)
}
