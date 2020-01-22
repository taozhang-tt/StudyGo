package main

import "fmt"

type MyQueue struct {
	Stack []int
	Queue []int
}

/** Initialize your data structure here. */
func Constructor() MyQueue {
	var s1 = make([]int, 0)
	var s2 = make([]int, 0)
	return MyQueue{
		Stack:s1,
		Queue:s2,
	}
}


/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int)  {
	this.Stack = append(this.Stack, x)
}


/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
	if len(this.Queue) == 0 {
		this.FromStackToQueue()
	}
	pop := this.Queue[len(this.Queue)-1]
	this.Queue = this.Queue[:len(this.Queue)-1]
	return pop
}


/** Get the front element. */
func (this *MyQueue) Peek() int {
	if len(this.Queue) == 0 {
		this.FromStackToQueue()
	}
	return this.Queue[len(this.Queue)-1]
}


/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
	if len(this.Stack) + len(this.Queue) == 0 {
		return true
	}
	return false
}

func (this *MyQueue) FromStackToQueue() {
	for len(this.Stack) > 0 {
		item := this.Stack[len(this.Stack)-1]
		this.Queue = append(this.Queue, item)
		this.Stack = this.Stack[:len(this.Stack)-1]
	}
}

func main()  {
	obj := Constructor()
	obj.Push(1)
	obj.Push(2)
	obj.Push(3)

	param2 := obj.Pop()
	fmt.Println("param_2: ", param2)
	param3 := obj.Peek()
	fmt.Println("param_3: ", param3)
	param4 := obj.Empty()
	fmt.Println("param_4: ", param4)
}