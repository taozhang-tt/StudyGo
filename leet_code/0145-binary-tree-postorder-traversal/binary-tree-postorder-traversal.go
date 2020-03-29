package main

import "fmt"

/**
145. 二叉树的后序遍历
	https://leetcode-cn.com/problems/binary-tree-postorder-traversal/
题目描述：
	给定一个二叉树，返回它的 后序 遍历。
	进阶: 递归算法很简单，你可以通过迭代算法完成吗？
*/

func main()  {
	var root = &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 4,
			},
			Right: &TreeNode{
				Val: 5,
			},
		},
		Right: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val: 6,
			},
			Right: &TreeNode{
				Val: 7,
			},
		},
	}
	ret := postorderTraversal(root)
	fmt.Println(ret)
}

/**
借助栈
	1. 初始：将跟节点入栈
	2. 取栈顶元素，
		2.1 判断该元素是否有右孩子，如果没有则继续下一步；如果有，判断右孩子是否被访问过？被访问过则继续下一步；未被访问，将右孩子入栈；
		2.2 判断该元素是否有左孩子，如果没有则继续下一步；如果有，判断左孩子是否被访问过？被访问过则继续下一步；未被访问，将左孩子入栈；
		2.3 判断 2.2 或 2.3 是否有成立的情况，如果有则直接继续下一步；如果均未成立，说明该节点无左右孩子或者是左右孩子均被访问过，则访问该节点、将该节点出栈、标记该节点已被访问过；
	3. 循环取栈顶元素进行 步骤2	
*/
func postorderTraversal(root *TreeNode) []int {
	ret := make([]int, 0)
	if root == nil {
		return ret
	}
	var stack, visited = make([]*TreeNode, 0), make(map[*TreeNode]bool, 0)
	stack = append(stack, root)
	for len(stack) > 0 {
		curr := stack[len(stack)-1]
		sign := true
		if curr.Right != nil && !visited[curr.Right] {
			stack = append(stack, curr.Right)
			sign = false
		}
		if curr.Left != nil && !visited[curr.Left] {
			stack = append(stack, curr.Left)
			sign = false
		}
		if sign {
			ret  = append(ret, curr.Val)
			visited[curr] = true
			stack = stack[0:len(stack)-1]
		}
	}
	return ret
}

 type TreeNode struct {
     Val int
     Left *TreeNode
     Right *TreeNode
 }
 