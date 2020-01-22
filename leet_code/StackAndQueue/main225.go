package main

import "fmt"

type MyStack struct {
	Queue []int
	Temp []int
}


/** Initialize your data structure here. */
func Constructor225() MyStack {
	var s = make([]int, 0)
	var q = make([]int, 0)
	return MyStack{
		Temp:s,
		Queue:q,
	}
}


/** Push element x onto stack. */
func (this *MyStack) Push(x int)  {
	this.Queue = append(this.Queue, x)
}


/** Removes the element on top of the stack and returns that element. */
func (this *MyStack) Pop() int {
	//先把队列Queue里的元素一个一个弹出，并暂存到 Temp 里
	for len(this.Queue) > 1 {
		item := this.Queue[0]
		this.Temp = append(this.Temp, item)
		this.Queue = this.Queue[1:]
	}
	pop := this.Queue[0]
	this.Queue = this.Temp
	this.Temp = nil
	return pop
}


/** Get the top element. */
func (this *MyStack) Top() int {
	top := this.Pop()
	this.Queue = append(this.Queue, top)
	return top
}


/** Returns whether the stack is empty. */
func (this *MyStack) Empty() bool {
	if len(this.Queue) > 0 {
		return false
	}
	return true
}

func main()  {
	obj := Constructor225()
	obj.Push(1)
	obj.Push(2)
	param2 := obj.Pop()
	fmt.Println("param2: ", param2)
	top := obj.Top()
	fmt.Println("top: ", top)
	empty := obj.Empty()
	fmt.Println("empty: ", empty)
}
