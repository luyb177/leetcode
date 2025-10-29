package main

//初版没对

//type MyLinkedList struct {
//	Val  int
//	Next *MyLinkedList
//}
//
//func Constructor() MyLinkedList {
//	myList := new(MyLinkedList)
//	return *myList
//}
//
//func (this *MyLinkedList) Get(index int) int {
//	size := 0
//	if index < 0 {
//		return -1
//	} else {
//		cur := this
//		for ; index > 0; index-- {
//			if cur.Next != nil {
//				cur = cur.Next
//				size++
//			} else {
//				return -1
//			}
//		}
//		return cur.Val
//	}
//}
//
//func (this *MyLinkedList) AddAtHead(val int) {
//	dummy := new(MyLinkedList)
//	newNode := &MyLinkedList{Val: val, Next: this}
//	dummy.Next = newNode
//
//}
//
//func (this *MyLinkedList) AddAtTail(val int) {
//	dummy := new(MyLinkedList)
//	dummy.Next = this
//	cur := dummy
//	newNode := MyLinkedList{Val: val, Next: nil}
//	for cur.Next != nil {
//		cur = cur.Next
//	}
//	cur.Next = &newNode
//}
//
//func (this *MyLinkedList) AddAtIndex(index int, val int) {
//	if index < 0 {
//		return
//	}
//	dummy := new(MyLinkedList)
//	dummy.Next = this
//	newNode := &MyLinkedList{Val: val}
//	cur := dummy
//	for i := 0; i < index; i++ {
//		if cur.Next == nil {
//			return
//		}
//		cur = cur.Next
//	}
//	newNode.Next = cur.Next
//	cur.Next = newNode
//}
//func (this *MyLinkedList) DeleteAtIndex(index int) {
//	if index < 0 {
//		return // 索引无效，直接返回
//	}
//	dummy := new(MyLinkedList)
//	dummy.Next = this
//	cur := dummy
//	for i := 0; i < index; i++ {
//		if cur.Next == nil {
//			return
//		}
//		cur = cur.Next
//	}
//	if cur.Next != nil {
//		cur.Next = cur.Next.Next
//	}
//}

//func main() {
//	myLinkedList := Constructor()
//
//	myLinkedList.AddAtTail(1)
//	myLinkedList.AddAtTail(2)
//	myLinkedList.AddAtTail(3)
//	myLinkedList.AddAtTail(4)
//	//fmt.Println(myLinkedList.Get())
//	myLinkedList.AddAtIndex(3, 6)
//	fmt.Println(myLinkedList.Get(3))
//	myLinkedList.DeleteAtIndex(0)
//	fmt.Println(myLinkedList.Get(0))
//}

type SingleNode struct {
	Val  int
	Next *SingleNode //下一个节点的指针
}

type MyLinkedList struct {
	dummyHead *SingleNode //虚拟头节点
	Size      int         //链表的长度
}

func Constructor() MyLinkedList {
	NewNode := &SingleNode{9999, nil}
	return MyLinkedList{NewNode, 0}
}
func (this *MyLinkedList) Get(index int) int {
	if this == nil || index < 0 || index >= this.Size {
		return -1
	}
	cur := this.dummyHead.Next //头节点
	for i := 0; i < index; i++ {
		cur = cur.Next
	}
	return cur.Val
}

func (this *MyLinkedList) AddAtHead(val int) {
	NewNode := &SingleNode{Val: val}
	NewNode.Next = this.dummyHead.Next
	this.dummyHead.Next = NewNode
	this.Size++
}
func (this *MyLinkedList) AddAtTail(val int) {
	NewNode := &SingleNode{Val: val}
	cur := this.dummyHead
	for cur.Next != nil {
		cur = cur.Next
	}
	cur.Next = NewNode
	this.Size++
}
func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if index < 0 {
		index = 0 //索引小于0就在前面插入
	} else if index > this.Size {
		return
	}

	newNode := &SingleNode{Val: val}
	cur := this.dummyHead
	for i := 0; i < index; i++ {
		cur = cur.Next
	}
	newNode.Next = cur.Next
	cur.Next = newNode
	this.Size++
}
func (this *MyLinkedList) DeleteAtIndex(index int) {
	if index < 0 || index >= this.Size {
		return
	}
	cur := this.dummyHead
	for i := 0; i < index; i++ {
		cur = cur.Next
	}
	if cur.Next != nil {
		cur.Next = cur.Next.Next
	}
	this.Size--
}
