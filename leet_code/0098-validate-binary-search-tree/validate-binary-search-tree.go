package main

import (
	"fmt"
	"math"
)

/**
98. 验证二叉搜索树
	https://leetcode-cn.com/problems/validate-binary-search-tree/
题目描述：
	给定一个二叉树，判断其是否是一个有效的二叉搜索树。
	假设一个二叉搜索树具有如下特征：
	1. 节点的左子树只包含小于当前节点的数。
	2. 节点的右子树只包含大于当前节点的数。
	3. 所有左子树和右子树自身必须也是二叉搜索树。
*/

//	二叉排序树的定义：
//	若左子树不空，则左子树上所有结点的值均小于它的根结点的值
//	若右子树不空，则右子树上所有结点的值均大于它的根结点的值
//	左、右子树也分别为二叉排序树
//	没有键值相等的结点

func main() {
	root := &TreeNode{
		Val: 10,
		Left: &TreeNode{
			Val: 5,
			Left: &TreeNode{
				Val:   3,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val: 8,
				Left: &TreeNode{
					Val:   6,
					Left:  nil,
					Right: nil,
				},
				Right: &TreeNode{
					Val:   12,
					Left:  nil,
					Right: nil,
				},
			},
		},
		Right: &TreeNode{
			Val:   15,
			Left:  nil,
			Right: nil,
		},
	}
	ret := isValidBST2(root)
	fmt.Println(ret)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//通过一次中序遍历储存所有节点到切片中，然后遍历一遍切片判断前一个值都小于后一个值
func isValidBST(root *TreeNode) bool {
	ret := inOrder(root)
	for i := 0; i < len(ret)-1; i++ {
		if ret[i].Val >= ret[i+1].Val {
			return false
		}
	}
	return true
}

func inOrder(root *TreeNode) []*TreeNode {
	ret := make([]*TreeNode, 0)
	if root != nil {
		ret = append(ret, inOrder(root.Left)...)
		ret = append(ret, root)
		ret = append(ret, inOrder(root.Right)...)
	}
	return ret
}

//中序遍历的优化版，每次遍历的时候保存前继节点，判断当前节点是否大于前继节点即可
//根据定义可知，中序遍历二叉排序树，每次遍历的节点值依此增大，那我们可以在遍历某个节点时，记录下它的值，在遍历下个节点时做比较，判断当前节点的值是否大于上一个节点
func isValidBST2(root *TreeNode) bool {
	intMax := int(^uint(0) >> 1)
	intMin := ^intMax
	//初始化一个默认的前继节点，节点的值为最小整数
	var prev = &TreeNode{
		Val:   intMin,
		Left:  nil,
		Right: nil,
	}
	return helper(root, prev)
}

func helper(root *TreeNode, prev *TreeNode) bool {
	if root == nil {
		return true
	}
	if !helper(root.Left, prev) {
		return false
	}
	//遍历了该节点，那该节点为下一个被遍历的节点的前继节点
	if prev != nil && prev.Val >= root.Val {
		return false
	}
	prev.Val = root.Val
	return helper(root.Right, prev)
}

func InOrder(root *TreeNode) {
	if root == nil {
		return
	}
	InOrder(root.Left)
	fmt.Println(root.Val)
	InOrder(root.Right)
}

//根据二叉排序树的定义，只需递归判断 左子树的最大值小于根结点的值且右子树的最小值大于根结点
func isValidBST3(root *TreeNode) bool {
	return isValidbst(root, math.Inf(-1), math.Inf(1))
}
func isValidbst(root *TreeNode, min, max float64) bool {
	if root == nil {
		return true
	}
	v := float64(root.Val)
	//isValidbst(root.Left, min, v) 限制左子树的最大值小于根结点
	//isValidbst(root.Right, v, max) 限制右子树的最小值大于根节点
	return v < max && v > min && isValidbst(root.Left, min, v) && isValidbst(root.Right, v, max)
}
