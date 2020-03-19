package main

import "fmt"

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func addTree(tree *TreeNode,val int ) *TreeNode {

	if tree == nil {

		return &TreeNode{Val:val}
	}

	if tree.Val > val {
		tree.Left = addTree(tree.Left,val)
	}
	if tree.Val <val {
		tree.Right = addTree(tree.Right,val)
	}
	return tree

}

func printT(tree *TreeNode)  {
	if tree == nil {
		return
	}
	fmt.Println(tree.Val)
	printT(tree.Left)
	printT(tree.Right)
}

func findMin(t *TreeNode) *TreeNode  {
	if t.Left == nil {
		return t
	}
	return findMin(t.Left)
}

func delMin(t *TreeNode) *TreeNode {
	if t.Left == nil {
		newNode := t.Right
		t.Right = nil
		return newNode
	}

	t.Left = delMin(t.Left)
	return t

}

func delTree(t *TreeNode,val int) *TreeNode  {
	if t == nil{
		return t
	}
	if t.Val > val {
		t.Left = delTree(t.Left,val)
		return t
	}else if t.Val < val {
		t.Right = delTree(t.Right,val)
		return t
	}else {
		if t.Left == nil {
			newNode := t.Right
			t.Right = nil
			return newNode
		}else if t.Right == nil {
			newNode := t.Left
			t.Left = nil
			return newNode
		} else{
			newNode := findMin(t.Right)
			newNode.Right = delMin(t.Right)
			newNode.Left = t.Left
			t.Right,t.Left = nil,nil
			return newNode
		}
	}


}

func main() {

	t := &TreeNode{Val:33}

	t = addTree(t,41)
	t = addTree(t,50)
	t = addTree(t,12)
	t = addTree(t,20)
	t = addTree(t,6)
	t = addTree(t,34)
	t = addTree(t,43)
	t = addTree(t,60)


	printT(t)

	fmt.Println("..........")

	printT(delTree(t,41))


}