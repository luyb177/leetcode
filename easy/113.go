package main

//func main() {
//
//}

func pathSum(root *TreeNode, targetSum int) [][]int {
	var result [][]int
	if root == nil {
		return result
	}

	var traversal func(node *TreeNode, curPath []int, count int)
	traversal = func(node *TreeNode, curPath []int, count int) {
		if node.Left == nil && node.Right == nil && count == 0 {
			// 这里要复制一份 curPath，否则会因为共享底层数组被修改
			tmp := make([]int, len(curPath))
			copy(tmp, curPath)
			result = append(result, curPath)
			return
		}
		if node.Left == nil && node.Right == nil {
			return
		}

		if node.Left != nil {
			curPath = append(curPath, node.Left.Val)
			count -= node.Left.Val
			traversal(node.Left, curPath, count)
			curPath = curPath[:len(curPath)-1]
			count += node.Left.Val
		}

		if node.Right != nil {
			curPath = append(curPath, node.Right.Val)
			count -= node.Right.Val
			traversal(node.Right, curPath, count)
			curPath = curPath[:len(curPath)-1]
			count += node.Right.Val
		}
	}
	curPath := []int{root.Val}
	traversal(root, curPath, targetSum-root.Val)
	return result
}
