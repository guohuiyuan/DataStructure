package main

import (
	"DataStructure/binarytree"
	"DataStructure/structure"
	"fmt"
)

func main()  {
	var root *structure.TreeNode
	root = structure.GetExampleTree()
	binarytree.PreorderTraversal(root)
	fmt.Println()
	binarytree.InorderTraversal(root)
	fmt.Println()
	binarytree.PostorderTraversal(root)
	fmt.Println()
	fmt.Println(binarytree.PreorderTraversalImprove(root))

}
