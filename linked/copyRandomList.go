package main

import "fmt"

type LMode struct {
	Val int
	Next *LMode
	Random *LMode
}

/**
复制带随机指针的链表
 */
func  CopyRandomList(head *LMode) *LMode {
	maps := make(map[*LMode]int)
	arr := make([]*LMode,0)
	prt := head
	var i int = 0
	for prt != nil{
		arr = append(arr, &LMode{Val:prt.Val})
		maps[prt] = i
		i++
		prt = prt.Next
		}
	prt = head
	i = 0
	for prt != nil{
		if prt.Next != nil {
			arr[i].Next = arr[i+1]
		}

		if prt.Random != nil {

			id,_ := maps[prt.Random]
			arr[i].Random = arr[id]
		}
		prt = prt.Next
		i++
	}


return arr[0]
}

func removeElements(head *LMode, val int) *LMode {
	for head != nil && head.Val == val {
		head = head.Next
	}
	var pre = head
	for pre != nil && pre.Next != nil {
		if pre.Next.Val == val{
			var cur = pre.Next.Next
			pre.Next = cur
		}
		pre = pre.Next
	}
	return head
}


func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	var res =&ListNode{Data:0}
	curr := res
	var p = l1
	var q = l2
	carry := 0;
	for p != nil || q != nil {
		var x int=0
		var y int =0
		if p != nil {
			x = p.Data
		}
		if q != nil {
			y = q.Data
		}
		sum := carry + x +y
		carry = sum /10
		curr.Next = &ListNode{Data:sum%10}
		curr = curr.Next
		if p != nil  {
			p = p.Next
		}
		if q != nil {
			q = q.Next
		}

	}
	if carry >0{
		curr.Next = &ListNode{Data:carry}
	}
	

	return res.Next
}

func CreateLMode(i int) *LMode {
	return &LMode{Val:i}
}

func Prints(head *LMode)  {
	for head != nil{
		fmt.Println("val =",head.Val)
		head = head.Next
	}
}

func main() {
	a:=CreateLMode(1)
	b:=CreateLMode(2)
	c:=CreateLMode(3)
	d:=CreateLMode(4)
	e:=CreateLMode(5)
	w:=CreateLMode(1)

	a.Next = b
	b.Next = c
	c.Next = d
	d.Next = e

	removeElements(w,1)

	a.Random = c
	b.Random = d
	c.Random = c
	e.Random = d

	p :=CopyRandomList(a)
	Prints(p)
}
