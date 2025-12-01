package main

//func main() {
//
//}

// 二叉搜索树自带的有序性
// 第一次遇见在 [p:q] 的区间的节点就是最近公共祖先节点

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	if root.Val > p.Val && root.Val > q.Val {
		// 走左子树
		left := lowestCommonAncestor(root.Left, p, q)
		if left != nil {
			return left
		}
	}

	if root.Val < p.Val && root.Val < q.Val {
		// 右子树
		right := lowestCommonAncestor(root.Right, p, q)
		if right != nil {
			return right
		}
	}

	// root.Val 在 [p:q] 或者[q:p]
	return root
}
