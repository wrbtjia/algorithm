package main

import "fmt"

type LNode struct {
	Data int
	Next *LNode
}

func create(i int) *LNode {
	return &LNode{Data:i}
}

func (n *LNode)add(i int)  {
	cur := n

	for  {
		if cur.Next != nil {
			cur = cur.Next
		}else {
			cur.Next=&LNode{}
			cur.Next.Data=i
			cur = cur.Next
			return
		}
	}


}
func print(n *LNode)  {
	for cur := n; cur != nil;{
		fmt.Println(cur.Data)
		cur = cur.Next
	}
}

func reverse(node *LNode) *LNode {
	if node == nil || node.Next == nil{
		return nil
	}
	var pre *LNode
	var curs *LNode


	for node != nil  {
		curs = node.Next
		node.Next = pre
		pre = node
		node = curs
	}
	return pre

}

func main() {
	n := create(1)
	n.add(2)
	n.add(3)
	n.add(4)
	n.add(5)
	n.add(6)
	n.add(7)
	print(n)

	d := reverse(n)

	print(d)
}