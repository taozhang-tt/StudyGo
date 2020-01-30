package main

import "fmt"

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
	fmt.Println(maxDepth2(root))
}

//广度优先搜索
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	max := 0
	for len(queue) > 0 {
		max ++
		levelSize := len(queue)
		for i:=0; i<levelSize; i++ {	//通过这层循环处理相同一层的节点
			curr := queue[0]
			if curr.Left != nil {
				queue = append(queue, curr.Left)
			}
			if curr.Right != nil {
				queue = append(queue, curr.Right)
			}
			queue = queue[1:]
		}
	}
	return max
}

//利用深度优先遍历
func maxDepth2(root *TreeNode) int {
	if root == nil {
		return 0
	}
	lLevel := maxDepth2(root.Left)
	rLevel := maxDepth2(root.Right)
	if lLevel > rLevel {
		return lLevel+1
	}
	return rLevel+1
}