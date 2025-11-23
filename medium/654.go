package main

//	func main() {
//		fmt.Printf("%#v", constructMaximumBinaryTree([]int{3, 2, 1, 6, 0, 5}))
//	}
func constructMaximumBinaryTree(nums []int) *TreeNode {
	var traversal func(num []int, left int, right int) *TreeNode

	traversal = func(num []int, left int, right int) *TreeNode { // [)
		if left >= right {
			return nil
		}

		// 最大值
		maxValue := num[left]
		maxValueIndex := left
		for i := left; i < right; i++ {
			if num[i] > maxValue {
				maxValue = num[i]
				maxValueIndex = i
			}
		}

		// 中
		mid := TreeNode{Val: maxValue}

		mid.Left = traversal(num, left, maxValueIndex)
		mid.Right = traversal(num, maxValueIndex+1, right)
		return &mid
	}

	return traversal(nums, 0, len(nums))
}
