package main

import "fmt"

/**
543. 二叉树的直径
	https://leetcode-cn.com/problems/diameter-of-binary-tree/
题目描述：
	给定一棵二叉树，你需要计算它的直径长度。一棵二叉树的直径长度是任意两个结点路径长度中的最大值。这条路径可能穿过根结点。
*/
func main() {
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val:   4,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   5,
				Left:  nil,
				Right: nil,
			},
		},
		Right: &TreeNode{
			Val:   3,
			Left:  nil,
			Right: nil,
		},
	}
	fmt.Println(diameterOfBinaryTree(root))
}

/**
思路：树的最大直径 = max(左子树的最大直径, 右子树的最大直径, 左子树的最大深度+右子树的最大深度)
*/
func diameterOfBinaryTree(root *TreeNode) int {
	_, diam := Depth(root)
	return diam
}

//深度优先搜索返回以 root 为根节点，它的最大深度和最大直径
func Depth(root *TreeNode) (dep int, dia int) {
	if root == nil {
		return 0, 0
	}
	leftDepth, leftDiameter := Depth(root.Left)
	rightDepth, rightDiameter := Depth(root.Right)
	dep = leftDepth + 1
	if rightDepth > leftDepth {
		dep = rightDepth + 1
	}
	dia = leftDepth + rightDepth
	if rightDiameter > dia {
		dia = rightDiameter
	}
	if leftDiameter > dia {
		dia = leftDiameter
	}
	return dep, dia
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
