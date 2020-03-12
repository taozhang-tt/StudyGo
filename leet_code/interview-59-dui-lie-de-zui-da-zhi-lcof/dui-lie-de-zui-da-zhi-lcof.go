package main

func main()  {
	
}

type MaxQueue struct {
	Set []int
}

func Constructor() MaxQueue {
	maxQueue := new(MaxQueue)
	maxQueue.Set = make([]int, 0)
	return *maxQueue
}


func (this *MaxQueue) Max_value() int {
	if len(this.Set) <= 0 {
		return -1
	}
	max := this.Set[0]
	for _, value := range this.Set {
		if value > max {
			max = value
		}
	}
	return max
}


func (this *MaxQueue) Push_back(value int)  {
	this.Set = append(this.Set, value)
}


func (this *MaxQueue) Pop_front() int {
	if len(this.Set) <= 0 {
		return -1
	}
	item := this.Set[0]
	this.Set = this.Set[1:]
	return item
}


/**
 * Your MaxQueue object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Max_value();
 * obj.Push_back(value);
 * param_3 := obj.Pop_front();
 */
