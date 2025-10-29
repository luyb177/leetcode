package main

//func main() {
//
//}

// 双指针
func reverseList(head *ListNode) *ListNode {
	cur := head
	var pre *ListNode = nil

	for cur != nil {
		temp := cur.Next //jilu
		//反转
		cur.Next = pre

		//改变变量
		pre = cur
		cur = temp

	}
	return pre
}

// 递归
func reverse(cur *ListNode, pre *ListNode) *ListNode {
	//递归终止条件
	if cur == nil {
		return pre
	}
	temp := cur.Next
	cur.Next = pre
	//进入下一个递归
	return reverse(temp, cur)

}
func reverseList1(head *ListNode) *ListNode {
	return reverse(head, nil)
}
