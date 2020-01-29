package main

import (
	"StudyGo/leet_code/Tree"
	"fmt"
)

func main() {
	tree := &Tree.TreeNode{
		Val:      3,
		Children: []*Tree.TreeNode{
			&Tree.TreeNode{
				Val:5,
				Children:[]*Tree.TreeNode{
					&Tree.TreeNode{
						Val:6,
						Children:[]*Tree.TreeNode{},
					},
					&Tree.TreeNode{
						Val:2,
						Children:[]*Tree.TreeNode{
							&Tree.TreeNode{
								Val:7,
								Children:[]*Tree.TreeNode{},
							},
							&Tree.TreeNode{
								Val:4,
								Children:[]*Tree.TreeNode{},
							},
						},
					},
				},
			},
			&Tree.TreeNode{
				Val:1,
				Children:[]*Tree.TreeNode{
					&Tree.TreeNode{
						Val:0,
						Children:[]*Tree.TreeNode{},
					},
					&Tree.TreeNode{
						Val:8,
						Children:[]*Tree.TreeNode{},
					},
				},
			},
		},
	}

	Tree.BreadthFirstSearch(tree)
	fmt.Println()
	Tree.DepthFirstSearch(tree)
	fmt.Println()
	Tree.DepthFirstSearch2(tree)
}
