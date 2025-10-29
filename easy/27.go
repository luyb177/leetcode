package main

//	func main() {
//		a := []int{3, 2, 2, 3}
//		fmt.Println(removeElement(a, 3))
//		fmt.Println(a)
//	}
func removeElement(nums []int, val int) int {
	fast := 0
	slow := 0
	for ; fast <= len(nums)-1; fast++ {
		if nums[fast] != val {
			nums[slow] = nums[fast]
			slow++
		}

	}
	return slow
}
