package main

import "fmt"

/**
102. 二叉树的层次遍历
	https://leetcode-cn.com/problems/binary-tree-level-order-traversal/
题目描述：
	给定一个二叉树，返回其按层次遍历的节点值。 （即逐层地，从左到右访问所有节点）。
*/

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
	ret := levelOrder2(root)
	fmt.Println(ret)

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//利用广度优先遍历
func levelOrder(root *TreeNode) [][]int {
	var ret [][]int
	if root == nil {
		return ret
	}
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	for len(queue) > 0 {
		levelSize := len(queue)
		var currentLevel []int
		for i := 0; i < levelSize; i++ {
			curr := queue[0]
			currentLevel = append(currentLevel, curr.Val)
			if curr.Left != nil {
				queue = append(queue, curr.Left)
			}
			if curr.Right != nil {
				queue = append(queue, curr.Right)
			}
			queue = queue[1:]
		}
		ret = append(ret, currentLevel)
	}
	return ret
}

//利用深度优先遍历
func levelOrder2(root *TreeNode) [][]int {
	var ret = make([][]int, 0)
	ret = dfs(root, 0, ret)
	return ret
}

func dfs(root *TreeNode, level int, ret [][]int) [][]int {
	if root == nil {
		return ret
	}
	if len(ret) > level {
		ret[level] = append(ret[level], root.Val)
	} else {
		ret = append(ret, []int{root.Val})
	}
	ret = dfs(root.Left, level+1, ret)
	ret = dfs(root.Right, level+1, ret)
	return ret
}
