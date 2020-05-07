package main
/**
572. 另一个树的子树
	给定两个非空二叉树 s 和 t，检验 s 中是否包含和 t 具有相同结构和节点值的子树。s 的一个子树包括 s 的一个节点和这个节点的所有子孙。s 也可以看做它自身的一棵子树
*/

func main()  {
	
}

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func isSubtree(s *TreeNode, t *TreeNode) bool {
	if s == nil {
		return false
	}
	return recursion(s, t) || isSubtree(s.Left, t) || isSubtree(s.Right, t)
}

func recursion(s, t *TreeNode) bool {
	if s == nil && t == nil {
		return true
	}
	if s == nil || t == nil {
		return false
	}
	if s.Val != t.Val {
		return false
	}
	return recursion(s.Left, t.Left) && recursion(s.Right, t.Right)
}