package main

import "fmt"

type Node struct {
	Val  int
	Rank int
}

type UnionFind struct {
	roots []*Node
}

func Construct(n int) *UnionFind {
	unionFind := &UnionFind{roots: make([]*Node, n, n)}
	for i := 0; i < n; i++ { //初始化每个节点为孤立节点，即该节点的父亲节点是其本身
		unionFind.roots[i] = &Node{
			Val:  i, //保存它的上级节点在 roots 切片里的索引
			Rank: 0,
		}
	}
	return unionFind
}

func (obj *UnionFind) Find(index int) int {
	if obj.roots[index].Val != index {
		obj.roots[index].Val = obj.Find(obj.roots[index].Val)
	}
	return obj.roots[index].Val
}

func (obj *UnionFind) Union(p, q int) {
	pRoot := obj.Find(p)
	qRoot := obj.Find(q)
	if pRoot != qRoot {
		obj.roots[qRoot].Val = obj.roots[pRoot].Val
	}
}

func (obj *UnionFind) Connected(p, q int) bool {
	if obj.Find(p) == obj.Find(q) {
		return true
	}
	return false
}

func main()  {
	unionFind := Construct(5)
	fmt.Println("init")
	for _, root := range unionFind.roots {
		fmt.Print(*root)
	}
	fmt.Println()

	fmt.Println()
	unionFind.Union(1, 2)
	fmt.Println("union 1, 2")
	for _, root := range unionFind.roots {
		fmt.Print(*root)
	}

	unionFind.Union(0, 1)
	fmt.Println("union 0, 1")
	for _, root := range unionFind.roots {
		fmt.Print(*root)
	}
}