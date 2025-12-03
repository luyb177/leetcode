package main

// func main() {
//
// }
func sortedArrayToBST(nums []int) *TreeNode {
	// 找中间元素进行分割 // -10 -3 0 5 9
	length := len(nums) // 5
	if length == 0 {
		return nil
	}

	var middle int
	if (length-1)%2 == 0 {
		middle = (length - 1) / 2 // 2
	} else {
		middle = (length-1)/2 + 1 // 取后面的那个元素
	}

	node := &TreeNode{Val: nums[middle]}

	node.Left = sortedArrayToBST(nums[:middle])
	node.Right = sortedArrayToBST(nums[middle+1:])
	return node
}
