package main

//func main() {
//
//}

func deleteNode(root *TreeNode, key int) *TreeNode {
	// 1.未找到
	if root == nil {
		return nil
	}

	if root.Val == key {
		if root.Left == nil && root.Right == nil {
			// 2. 左右节点都为空
			return nil
		} else if root.Left == nil {
			// 3. 左节点为空，右节点不为空
			return root.Right
		} else if root.Right == nil {
			// 4. 左节点不为空，右节点为空
			return root.Left
		} else {
			// 5. 左右节点都不为空
			// 将左节点放在右子树最左侧的节点上
			var cur *TreeNode
			cur = root.Right
			for cur.Left != nil {
				cur = cur.Left
			}
			cur.Left = root.Left
			return root.Right
		}
	}

	if root.Val > key {
		root.Left = deleteNode(root.Left, key)
	}
	if root.Val < key {
		root.Right = deleteNode(root.Right, key)
	}
	return root
}
