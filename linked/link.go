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
func (n *LNode)print()  {
	for cur := n; cur != nil;{
		fmt.Println(cur.Data)
		cur = cur.Next
	}
}

func main() {
	n := create(1)
	n.add(2)
	n.add(3)
	n.add(4)
	n.print()

}