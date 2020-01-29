package Tree

import (
	"fmt"
)

type TreeNode struct {
	Val int
	Children []*TreeNode
}

func BreadthFirstSearch(root *TreeNode) {
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	for len(queue) > 0 {
		curr := queue[0]
		fmt.Print(curr.Val, ", ")
		for _, item := range curr.Children {
			queue = append(queue, item)
		}
		queue = queue[1:]
	}
}
