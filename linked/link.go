package main

import "fmt"

type LNode struct {
	Data int
	Next *LNode
}

func create(i int) *LNode {
	return &LNode{Data: i}
}

func (n *LNode) add(i int) {
	cur := n

	for {
		if cur.Next != nil {
			cur = cur.Next
		} else {
			cur.Next = &LNode{}
			cur.Next.Data = i

			return
		}
	}

}
func print(n *LNode) {
	for cur := n; cur != nil; {
		fmt.Println(cur.Data)
		cur = cur.Next
	}
}

/*func reverseList(head *ListNode) *ListNode {
	cur := head
	var pre *ListNode

	for cur != nil{
		pre, cur, cur.Next = cur, cur.Next, pre
	}
	return pre
}*/

func reverse(node *LNode) *LNode {
	if node == nil || node.Next == nil {
		return nil
	}
	var pre *LNode
	var curs = new(LNode)

	for node != nil {
		curs = node.Next

		node.Next = pre
		pre = node
		node = curs
	}

	return pre

}

/**
通过栈的 先进后出的方式 反转
*/
func reverseStack(node *LNode) *LNode {
	if node == nil || node.Next == nil {
		return node
	}

	newNode := reverseStack(node.Next)

	temp := node.Next
	temp.Next = node
	node.Next = nil

	return newNode

}

/**
以K为节点翻转链表
*/
func reverseK(node *LNode, k int) *LNode {

	var pre = node
	begin := pre
	var end *LNode
	var pNext *LNode

	for begin != nil {
		end = begin
		for i := 1; i < k; i++ {
			if end.Next != nil {
				end = end.Next

			}
			if end.Next == nil {
				break
			}
		}

		pNext = end.Next
		end.Next = nil

		pre = reverse(begin)

		begin.Next = pNext
		begin = pNext

	}
	return pre

}

func reverseKGroup(head *LNode, k int) *LNode {
	/*	var n = head
		for i:=0;i<k;i++ {
			if n == nil {
				return head
			}
			n = n.Next
		}
		//头插法
		var cp = head
		var cur = new(LNode)
		var last = cur
		for i:=0;i<k;i++ {
			cp = head.Next
			head.Next = cur.Next
			cur.Next = head
			head = cp
			if last.Next != nil {
				last = last.Next
			}
		}
		res := reverseKGroup(head,k)
		last.Next = res
		return cur.Next*/

	temp := &LNode{}
	temp.Next = head
	pre := temp
	tail := temp
	for {
		head = pre.Next
		for i := k; i > 0 && tail != nil; i-- {
			tail = tail.Next
		}
		if tail == nil {
			break
		}
		for pre.Next != tail {
			cur := pre.Next
			pre.Next = cur.Next
			cur.Next = tail.Next
			tail.Next = cur
		}
		pre = head
		tail = pre
	}
	return temp.Next
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

	/*	d := reverse(n)
		print(d)
	*/
	fmt.Println()

	d := reverseK(n, 4)
	print(d)

	/*	d := reverseKGroup(n,3)
		print(d)
	*/

	/*d:=reverseStack(n)
	print(d)*/

}
