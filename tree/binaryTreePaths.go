package main

import "strconv"

/**
257
给定一个二叉树，返回所有从根节点到叶子节点的路径。

说明: 叶子节点是指没有子节点的节点。

示例:

输入:

   1
 /   \
2     3
 \
  5

输出: ["1->2->5", "1->3"]

解释: 所有根节点到叶子节点的路径为: 1->2->5, 1->3

 */

func binaryTreePaths(root *TreeNode) []string {
	var res = []string{""}

	if root == nil{
		return res
	}

	if root.Left == nil && root.Right == nil {
		res = append(res, strconv.Itoa(root.Val))
		return res
	}
	chang := binaryTreePaths(root.Left)
	for i:=0;i< len(chang);i++  {
		res = append(res, strconv.Itoa(root.Val)+"->"+chang[i])
	}

	chang = binaryTreePaths(root.Right)
	for i:=0;i< len(chang);i++  {
		res = append(res, strconv.Itoa(root.Val)+"->"+chang[i])
	}

	return res
}