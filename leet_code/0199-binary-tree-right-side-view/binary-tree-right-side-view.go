package main

import "fmt"

func main() {
	var root *TreeNode
	var ret []int
	root = &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Right: &TreeNode{
				Val:  5,
				Left: nil,
			},
		},
		Right: &TreeNode{
			Val: 3,
			Right: &TreeNode{
				Val:  4,
				Left: nil,
			},
		},
	}
	ret = rightSideView2(root)
	fmt.Println(ret)

	root = new(TreeNode)
	ret = rightSideView2(root)
	fmt.Println(ret)

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
先序遍历
	先序遍历的顺序是：中，左，右；
	在二叉树的遍历过程中，每一层的元素使用切片中的一个元素做记录，最先放进切片的元素，会被同一层后放进切片的元素覆盖掉，这刚好也是右视图的结果
*/
func rightSideView(root *TreeNode) []int {
	var ret = make([]int, 0)
	dfs(root, 0, &ret)
	return ret
}

func dfs(root *TreeNode, level int, ret *[]int) {
	if root == nil {
		return
	}
	if level >= len(*ret) { //将该节点的值放到对应的切片位置存储
		*ret = append(*ret, root.Val)
	} else {
		(*ret)[level] = root.Val
	}
	dfs(root.Left, level+1, ret) //先放左节点，这样如果右节点有值的话会覆盖左节点的值
	dfs(root.Right, level+1, ret)
}

/**
层序遍历
	通过层序遍历，层序遍历的时候每次先访问右侧孩子节点，记录每一层的第一个节点的值
*/
func rightSideView2(root *TreeNode) []int {
	ret := make([]int, 0)
	if root == nil {
		return ret
	}
	queue, level := []*TreeNode{root}, 0
	for len(queue) > 0 {
		length := len(queue)
		for length > 0 {
			length--
			if len(ret) == level {
				ret = append(ret, queue[0].Val)
			}
			if queue[0].Right != nil {
				queue = append(queue, queue[0].Right)
			}
			if queue[0].Left != nil {
				queue = append(queue, queue[0].Left)
			}
			queue = queue[1:]
		}
		level++
	}
	return ret
}
