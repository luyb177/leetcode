package main

// func main() {
//
// }
type MyStack struct {
	Queue1 []int
	Queue2 []int
}

func Constructor() MyStack {
	return MyStack{
		Queue1: make([]int, 0),
		Queue2: make([]int, 0),
	}
}

func (this *MyStack) Push(x int) {
	this.Queue1 = append(this.Queue1, x)
}

func (this *MyStack) Pop() int {
	if len(this.Queue1) == 0 {
		return -1
	}
	// 首先将Queue1中的除了最后一个，其他放入queue2
	for len(this.Queue1) > 1 {
		this.Queue2 = append(this.Queue2, this.Queue1[0])
		this.Queue1 = this.Queue1[1:]
	}

	// 弹出最后一个
	val := this.Queue1[0]
	this.Queue1 = this.Queue1[1:]
	// 将queue2中的数据全部放入queue1中
	for len(this.Queue2) > 0 {
		this.Queue1 = append(this.Queue1, this.Queue2[0])
		this.Queue2 = this.Queue2[1:]
	}
	return val
}

func (this *MyStack) Top() int {
	if len(this.Queue1) == 0 {
		return -1
	}
	return this.Queue1[len(this.Queue1)-1]
}

func (this *MyStack) Empty() bool {
	return len(this.Queue1) == 0
}
