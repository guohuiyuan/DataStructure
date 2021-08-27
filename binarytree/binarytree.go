package binarytree

import (
	"DataStructure/structure"
	"fmt"
)

func PreorderTraversal(root *structure.TreeNode)  {
	if root==nil{
		return
	}
	fmt.Print(root.Val)
	PreorderTraversal(root.Left)
	PreorderTraversal(root.Right)
}
func PreorderTraversalImprove(root *structure.TreeNode) []int {
	// 非递归
	if root == nil{
		return nil
	}
	result:=make([]int,0)
	stack:=make([]*structure.TreeNode,0)

	for root!=nil || len(stack)!=0{
		for root !=nil{
			// 前序遍历，所以先保存结果
			result=append(result,root.Val)
			stack=append(stack,root)
			root=root.Left
		}
		// pop
		node:=stack[len(stack)-1]
		stack=stack[:len(stack)-1]
		root=node.Right
	}
	return result
}
func InorderTraversal(root *structure.TreeNode)  {
	if root==nil{
		return
	}
	InorderTraversal(root.Left)
	fmt.Print(root.Val)
	InorderTraversal(root.Right)
	
}
func PostorderTraversal(root *structure.TreeNode)  {
	if root==nil{
		return
	}
	PostorderTraversal(root.Left)
	PostorderTraversal(root.Right)
	fmt.Print(root.Val)
}
