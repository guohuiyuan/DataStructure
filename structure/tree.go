package structure

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func GetExampleTree() *TreeNode {
	t3:=&TreeNode{3,nil,nil}
	t5:=&TreeNode{5,nil,nil}
	t4:=&TreeNode{4,t3,t5}

	t0:=&TreeNode{0,nil,nil}
	t2:=&TreeNode{2,t0,t4}

	t7:=&TreeNode{7,nil,nil}
	t9:=&TreeNode{9,nil,nil}
	t8:=&TreeNode{8,t7,t9}

	t6:=&TreeNode{6,t2,t8}
	return t6
}
