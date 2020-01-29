package Tree

import "fmt"

//递归写法
func DepthFirstSearch(root *TreeNode) {
	if root == nil {
		return
	}
	fmt.Print(root.Val, ", ")
	for _, item := range root.Children {
		DepthFirstSearch(item)
	}
}

//非递归写法
func DepthFirstSearch2(root *TreeNode) {
	stack := make([]*TreeNode, 0)
	stack = append(stack, root)
	for len(stack) > 0 {
		curr := stack[len(stack)-1] //取栈顶元素
		stack = stack[:len(stack)-1] //弹出栈顶元素
		fmt.Print(curr.Val, ", ")
		for i := len(curr.Children) - 1; i >= 0; i-- { //孩子节点依次入栈
			stack = append(stack, curr.Children[i])
		}
	}
}
