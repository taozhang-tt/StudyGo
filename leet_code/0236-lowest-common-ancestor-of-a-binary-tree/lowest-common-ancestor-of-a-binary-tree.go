package main

import "fmt"

func main() {
	root := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 5,
			Left: &TreeNode{
				Val:   6,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val: 2,
				Left: &TreeNode{
					Val:   7,
					Left:  nil,
					Right: nil,
				},
				Right: &TreeNode{
					Val:   4,
					Left:  nil,
					Right: nil,
				},
			},
		},
		Right: &TreeNode{
			Val:   1,
			Left:  &TreeNode{
				Val:   0,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   8,
				Left:  nil,
				Right: nil,
			},
		},
	}
	p := &TreeNode{
		Val:   5,
		Left:  nil,
		Right: nil,
	}
	q := &TreeNode{
		Val:   4,
		Left:  nil,
		Right: nil,
	}
	ret := lowestCommonAncestor(root, p, q)
	fmt.Println(ret.Val)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//分别查找 p 和 q 的路径，然后判断路径上重合的节点
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	pPath := findPath(root, p)
	qPath := findPath(root, q)
	var long int
	if len(pPath) < len(qPath) {
		long = len(pPath)
	} else {
		long = len(qPath)
	}
	for i:=0; i<long; i++ {
		if pPath[i].Val != qPath[i].Val {
			return pPath[i-1]
		}
	}
	return pPath[long-1]
}

func findPath(root *TreeNode, node *TreeNode) []*TreeNode {
	ret := make([]*TreeNode, 0)
	if root == nil {
		return ret
	}
	if root.Val == node.Val {
		return append(ret, root)
	}
	left := findPath(root.Left, node)
	if len(left) > 0 {
		ret = append(ret, root)
		return append(ret, left...)
	}
	right := findPath(root.Right, node)
	if len(right) > 0 {
		ret = append(ret, root)
		return append(ret, right...)
	}
	return ret
}

//这道题目之所以能用这种递归解法，是因为有两个前提条件：1 所有节点的值都是唯一的；2 p、q 为不同节点且均存在于给定的二叉树中
func lowestCommonAncestor2(root, p, q *TreeNode) *TreeNode {
	//1，如果是一棵空树，那么直接返回
	//2，假设树的深度为 2，且 p 或 q 两者中有一个存在于根节点，那么两者的最近公共祖先为根节点
	if root == nil || root == p || root == q {
		return root
	}
	//如果 p 或 q 均不存在于根节点，那要查寻左子树和右子树是否包含 p、q
	left := lowestCommonAncestor2(root.Left, p, q)
	right := lowestCommonAncestor2(root.Right, p, q)
	//如果左子树既不包含p也不包含q，那说明 p、q 存在于右子树，返回右子树查询到的最近公共祖先即可
	if left == nil {
		return right
	}
	//如果右子树既不包含p也不包含q，那说明 p、q 存在于左子树，返回左子树查询到的最近公共祖先即可
	if right == nil {
		return left
	}
	return root
}
