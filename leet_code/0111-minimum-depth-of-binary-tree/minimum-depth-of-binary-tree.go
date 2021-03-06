package main

import "fmt"

/**
111. 二叉树的最小深度
	https://leetcode-cn.com/problems/minimum-depth-of-binary-tree/
题目描述：
	给定一个二叉树，找出其最小深度。
	最小深度是从根节点到最近叶子节点的最短路径上的节点数量。
	说明: 叶子节点是指没有子节点的节点。
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

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
			Val: 1,
			Left: &TreeNode{
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
	fmt.Println(minDepth2(root))
}

//利用广度优先搜索
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	min := 0
	for len(queue) > 0 {
		min++
		currLevelSize := len(queue)
		for i := 0; i < currLevelSize; i++ {
			curr := queue[0]
			if curr.Left == nil && curr.Right == nil {
				return min
			}
			if curr.Left != nil {
				queue = append(queue, curr.Left)
			}
			if curr.Right != nil {
				queue = append(queue, curr.Right)
			}
			queue = queue[1:]
		}
	}
	return min
}

//利用深度优先搜索
func minDepth2(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil {
		return minDepth2(root.Right) + 1
	}
	if root.Right == nil {
		return minDepth2(root.Left) + 1
	}
	lMin := minDepth2(root.Left)
	rMin := minDepth2(root.Right)
	if lMin < rMin {
		return lMin + 1
	}
	return rMin + 1
}
