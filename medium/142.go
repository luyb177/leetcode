package main

//func main() {
//
//}
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func detectCycle(head *ListNode) *ListNode {
	//双指针
	slow := head
	fast := head
	index1 := head
	//有“圆”定会相遇
	for fast != nil && fast.Next != nil {
		//fast 每次走两步
		//slow 每次走一步
		fast = fast.Next.Next
		slow = slow.Next

		//判断是否相遇
		if fast == slow {
			//因为路程的原因
			//slow必然在没走完一圈就被追上
			//x = (n-1)(y + z) + z
			//x 是head到入环点
			//y 是相遇时候slow走的路程
			//z 是一圈剩下的路程
			//2(x + y) = x + y + n(y + z)

			index2 := fast
			for {
				if index1 == index2 {
					return index1
				}
				//相同的路程
				//相遇在入环点
				index1 = index1.Next
				index2 = index2.Next
			}
		}

	}
	return nil
}
