package main

//144. 二叉树的前序遍历
/**
给定一个二叉树，返回它的 前序 遍历。

 示例:

输入: [1,null,2,3]
   1
    \
     2
    /
   3

输出: [1,2,3]

 */
func preorderTraversal(root *TreeNode) []int {

	var sum []int
	echo(root,&sum)
	return sum
}

func echo(root *TreeNode,sum *[]int)  {
	if root == nil {
		return
	}
	*sum = append(*sum, root.Val)
	echo(root.Left,sum)
	echo(root.Right,sum)
}

func main() {

}