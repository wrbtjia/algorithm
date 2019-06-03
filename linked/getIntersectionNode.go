package main

import "fmt"

type ListNode struct {
	Data int
	Next *ListNode
}

/**
分隔链表
输入: head = 1->4->3->2->5->2, x = 3
输出: 1->2->2->4->3->5
 */
func partition(head *ListNode, x int) *ListNode {
	var less = &ListNode{Data:0}
	var mor = &ListNode{Data:0}

	var less_prt = less
	var mor_prt = mor
	for head !=nil  {
		if head.Data < x {
			less_prt.Next = head

			less_prt = less_prt.Next

		}else {
			mor_prt.Next = head
			mor_prt = mor_prt.Next
		}
		head = head.Next
	}
	less_prt.Next = mor.Next
	mor_prt.Next = nil
	return less.Next
}

/**
判断链表是否有环
 */
func detectCycle(head *ListNode) *ListNode {
	var slow = head
	var fast = head
	var meet *ListNode

	for fast != nil{
		fast = fast.Next
		slow = slow.Next
		if fast == nil {
			return &ListNode{}
		}
		fast = fast.Next
		if fast == slow{
			meet = fast
			break
		}
	}

	if meet == nil{
		return nil
	}

	for  {
		if meet == head{
			head.Next = nil
			return head
		}
		meet = meet.Next
		head = head.Next
	}
	return nil
	
}

/**
判断两个链表共同指向的节点
 */
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	maps := make(map[*ListNode]int)
	var rest *ListNode
	one := headA
	for one.Next != nil{
		maps[one]=one.Data
		one = one.Next

	}
	for headB.Next != nil {
		if _,ok := maps[headB]; ok{
			rest = headB
			break
		}
		headB = headB.Next
	}
	return rest
}

func CreateListNode(i int)  *ListNode{
	return &ListNode{Data:i}
}

func (l* ListNode)addListNode(i int)  {
	l.Next = &ListNode{Data:i}
}

func main() {
	a := CreateListNode(1)

	b := CreateListNode(2)
	c := CreateListNode(3)
	d := CreateListNode(4)
	e := CreateListNode(5)
	f := CreateListNode(6)

	a.Next = b
	b.Next = c
	c.Next=d
	d.Next =e
	e.Next = f

	a1 := CreateListNode(7)
	a2 := CreateListNode(8)
	a3 := CreateListNode(9)

	a1.Next=a2
	a2.Next=a3
	a3.Next =d


//	p := getIntersectionNode(a,a1)



	n3 := CreateListNode(1)

//	b2 := CreateListNode(2)
//	c2 := CreateListNode(3)
/*	d2 := CreateListNode(4)
	e2 := CreateListNode(5)
	f2 := CreateListNode(6)
	g2 := CreateListNode(7)
	h2 := CreateListNode(8)*/

	/*n3.Next = b2
	b2.Next = n3*/
/*	c2.Next = d2
	d2.Next = e2
	e2.Next = f2
	f2.Next = g2
	g2.Next = h2
	h2.Next = c2*/

	p:=detectCycle(n3)
	PrintListNode(p)


	/*p1 := CreateListNode(1)
	p2 := CreateListNode(4)
	p3 := CreateListNode(3)
	p4 := CreateListNode(2)
	p5 := CreateListNode(5)
	p6 := CreateListNode(2)

	p1.Next=p2
	p2.Next =p3
	p3.Next = p4
	p4.Next =p5
	p5.Next =p6

	p:= partition(p1,3)
	PrintListNode(p)*/
}


func PrintListNode(l *ListNode)  {
	for no := l; no != nil;{
		fmt.Println(no.Data)
		no = no.Next
	}
}