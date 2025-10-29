package main

//func main() {
//
//}
// go中没有栈，首先要实现一个栈
// 使用切片实现一个栈

type Mystack []int

func (s *Mystack) Push(v int) {
	*s = append(*s, v)
}

func (s *Mystack) Pop() int {
	v := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return v
}

func (s *Mystack) Peek() int {
	v := (*s)[len(*s)-1]
	return v
}

func (s *Mystack) Size() int {
	return len(*s)
}

func (s *Mystack) Empty() bool {
	return len(*s) == 0
}

// 队列

type MyQueue struct {
	StackIn  *Mystack // 输入栈
	StackOut *Mystack // 输出栈
}

//func Constructor() MyQueue {
//	return MyQueue{
//		StackIn:  &Mystack{},
//		StackOut: &Mystack{},
//	}
//}

func (this *MyQueue) Push(x int) {
	this.StackIn.Push(x)
}

func (this *MyQueue) Pop() int {
	// 查看输出栈是否为空
	if this.StackOut.Empty() {
		// 空的话，将输入栈全部输出到输出栈
		for !this.StackIn.Empty() {
			v := this.StackIn.Pop()
			this.StackOut.Push(v)
		}
	}
	return this.StackOut.Pop()
}

func (this *MyQueue) Peek() int {
	// 先pop 再push
	v := this.Pop()
	this.StackOut.Push(v)
	return v
}

func (this *MyQueue) Empty() bool {
	return this.StackIn.Empty() && this.StackOut.Empty()
}
