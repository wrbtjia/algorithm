package main

import "fmt"

/**
222
给出一个完全二叉树，求出该树的节点个数。
输入:
    1
   / \
  2   3
 / \  /
4  5 6

输出: 6
 */

var g int = 0
func countNodes(root *treeNoded) int {
	if root == nil || root.Right == nil || root.Left == nil{
		return 0
	}
	g = g+1
	countNodes(root.Left)
	countNodes(root.Right)

	return g

}

type treeNoded struct {
	Val int
	Left *treeNoded
	Right *treeNoded
}

func main() {
	t := &treeNoded{
		Val:1,
	}

	fmt.Println(countNodes(t))
}