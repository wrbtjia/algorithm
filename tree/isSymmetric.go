package main


/**
101

给定一个二叉树，检查它是否是镜像对称的。

例如，二叉树 [1,2,2,3,4,4,3] 是对称的。

    1
   / \
  2   2
 / \ / \
3  4 4  3
但是下面这个 [1,2,2,null,3,null,3] 则不是镜像对称的:

    1
   / \
  2   2
   \   \
   3    3


 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return isSmae(root.Left,root.Right)
}

func isSmae(p *TreeNode,q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}

	if p !=nil && q != nil && p.Val == q.Val {
		return isSmae(p.Left,q.Right) && isSmae(p.Right,q.Left)
	}else {
		return false
	}

}