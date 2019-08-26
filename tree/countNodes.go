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


func countNodes(root *treeNoded) int {
	if root == nil{
		return 0
	}

	return countNodes(root.Left) + countNodes(root.Right) +1

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