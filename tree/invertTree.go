package main


/**
226
翻转一棵二叉树。
示例：

输入：

     4
   /   \
  2     7
 / \   / \
1   3 6   9
输出：

     4
   /   \
  7     2
 / \   / \
9   6 3   1

 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func invertTree(root *TreeNode) *TreeNode {
	if root == nil{
		return root
	}

	invertTree(root.Left)
	invertTree(root.Right)

	root.Left,root.Right = root.Right,root.Left
  	return root

}