package main

//func main() {
//
//}

// 二叉树的定义

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 二叉树的递归遍历

// PreorderTraversal 前序遍历 中 左 右
func PreorderTraversal(root *TreeNode) []int {
	var res []int
	var traversal func(node *TreeNode)
	traversal = func(node *TreeNode) {
		if node == nil {
			return
		}
		// 中
		res = append(res, node.Val)
		// 左
		traversal(node.Left)
		// 右
		traversal(node.Right)
	}
	traversal(root)
	return res
}

// InorderTraversal 中序遍历 左 中 右
func InorderTraversal(root *TreeNode) []int {
	var res []int
	var traversal func(node *TreeNode)
	traversal = func(node *TreeNode) {
		if node == nil {
			return
		}
		// 左
		traversal(node.Left)
		// 中
		res = append(res, node.Val)
		// 右
		traversal(node.Right)
	}
	traversal(root)
	return res
}

// PostorderTraversal 后序遍历 左 右 中
func PostorderTraversal(root *TreeNode) []int {
	var res []int
	var traversal func(node *TreeNode)
	traversal = func(node *TreeNode) {
		if node == nil {
			return
		}
		// 左
		traversal(node.Left)
		// 右
		traversal(node.Right)
		// 中
		res = append(res, node.Val)
	}
	traversal(root)
	return res
}

// 二叉树的迭代遍历
//
//// 前序遍历
//func preorderTraversal(root *TreeNode) []int {
//	var res []int
//
//	if root == nil {
//		return res
//	}
//
//	// 创建一个栈
//	stack := list.New()
//	// 添加一个根节点
//	stack.PushBack(root)
//
//	for stack.Len() > 0 {
//		// 移出
//		node := stack.Remove(stack.Back()).(*TreeNode)
//
//		res = append(res, node.Val)
//		// 这里是右先入栈，然后左再入栈 可以实现左先出栈，右后出栈
//		if node.Right != nil {
//			stack.PushBack(node.Right)
//		}
//		if node.Left != nil {
//			stack.PushBack(node.Left)
//		}
//	}
//	return res
//}
//
//// 中序遍历
//
//func inorderTraversal(root *TreeNode) []int {
//	var res []int
//	if root == nil {
//		return res
//	}
//	stack := list.New()
//	current := root // 使用指针来深度遍历
//	for current != nil || stack.Len() > 0 {
//		if current != nil {
//			stack.PushBack(current)
//			current = current.Left // 深度遍历到 左边最底部
//		} else {
//			node := stack.Remove(stack.Back()).(*TreeNode)
//			res = append(res, node.Val) // 中
//			current = node.Right        // 右
//		}
//	}
//	return res
//}
//
//// 后序遍历  前序遍历 是 中左右  改变顺序为 中右左   然后再反转结果为 左右中
//
//func postorderTraversal(root *TreeNode) []int {
//	var res []int
//
//	if root == nil {
//		return res
//	}
//
//	stack := list.New()
//	stack.PushBack(root)
//	for stack.Len() > 0 {
//		node := stack.Remove(stack.Back()).(*TreeNode)
//		res = append(res, node.Val) // 中
//		if node.Left != nil {
//			stack.PushBack(node.Left)
//		}
//		if node.Right != nil {
//			stack.PushBack(node.Right)
//		}
//	} // 出来顺序为 中 右 左
//
//	// 反转
//	for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
//		res[i], res[j] = res[j], res[i]
//	}
//	return res
//}

// 二叉树迭代法——统一风格
// 有点难——暂且放下
