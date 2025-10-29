package main

// func main() {
//
// }
func findBottomLeftValue(root *TreeNode) int {
	var maxDepth = -1
	var result int
	var traversal func (node *TreeNode,depth int)
	traversal = func (node *TreeNode,depth int){
		if node.Left == nil && node.Right == nil {
			// 判断
			if depth > maxDepth {
				maxDepth = depth // 此处直接赋值
				result = node.Val
			}
			return
		}

		if node.Left != nil {
			depth++
			traversal(node.Left,depth)
			depth-- // 回溯
		}
		if node.Right != nil {
			depth++
			traversal(node.Right,depth)
			depth--
		}
	}
	traversal(root,0)
	return result
}
