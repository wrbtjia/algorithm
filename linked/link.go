package main

import (
	"fmt"
	"time"
)

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



    func reverseList(head *ListNode) *ListNode {
		cur := head
		var pre *ListNode
		for cur != nil{
			pre, cur, cur.Next = cur, cur.Next, pre
		}
		return pre
	}


func reverse(node *LNode) *LNode {
	if node == nil || node.Next == nil {
		return nil
	}
	var pre *LNode
	var curs *LNode

	for node !=nil {
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

	if node == nil {
		return node
	}

	var tail = node

	var head = node

	for  {
		count :=0
		i :=1
		for i<=k  {
			if tail != nil {
				tail = tail.Next
				count ++
			}
			i++
		}
		if count<k{
			return node
		}

		var newNode *LNode
		for head != tail {
			cur := head.Next
			head.Next = newNode
			newNode = head
			head = cur
		}

		node.Next = reverseK(head,k)

		return newNode

	}


	return node

}
/**
反转链表中间段
 */
func twostage(head *LNode, m,n int) *LNode {
	var rest = head
	var pre *LNode
	for i:=1;i<m;i++ {
		pre = head
		head = head.Next
	}
	var be = head
	var newhead *LNode
	for j:=0;j<n-m+1;j++ {
		news := head.Next
		head.Next = newhead
		newhead = head
		head = news
	}
	if pre != nil {
		pre.Next = newhead

	}else {
		rest = newhead
	}
	be.Next = head
	return rest
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
	/*	n.add(6)
		n.add(7)*/
	print(n)


/*	m := create(1)
	m.add(3)
	m.add(4)

	print(m)
*/


		d := reverse(n)
		print(d)

	fmt.Println()

/*	d:=twostage(n,1,4)
	print(d)*/

//	d := reverseK(n, 3)
//	print(d)

	/*	d := reverseKGroup(n,3)
		print(d)*/


/*	d:=reverseStack(n)
	print(d)*/


	/*d:=mergeTwoLists(n,m)
	print(d)*/

}
/**
合并两个有序链表
输入：1->2->4, 1->3->4
输出：1->1->2->3->4->4
 */
func mergeTwoLists(l1 *LNode, l2 *LNode) *LNode {

	begin := time.Now()

	var rest = &LNode{Data:0}
	var head = rest
	for l1 != nil && l2 != nil {
		if l1.Data > l2.Data {
			head.Next = l2
			l2 = l2.Next
		}else {
			head.Next = l1
			l1 = l1.Next
		}
		head = head.Next
	}

	if l1 != nil {
		head.Next = l1
	}
	if l2 != nil {
		head.Next = l2
	}

	end := time.Now()
	fmt.Println(end.Sub(begin))

	return rest.Next
}

/**
合并K个排序链表
输入:
[
  1->4->5,
  1->3->4,
  2->6
]
输出: 1->1->2->3->4->4->5->6
 */
func mergeKLists(lists []*LNode) *LNode {
	if len(lists)  == 0 {
		return nil
	}
	if len(lists)  == 1 {
		return lists[0]
	}
	if len(lists)  == 2 {
		return mergeTwoLists(lists[0],lists[1])
	}
	mid := len(lists) /2
	var sub1 []*LNode
	var sub2 []*LNode
	for i := 0;i<mid;i++ {
		sub1 = append(sub1, lists[i])
	}

	for j:=mid;j<len(lists);j++   {
		sub2 = append(sub2, lists[j])
	}

	l1 := mergeKLists(sub1)

	l2 := mergeKLists(sub2)

	return mergeTwoLists(l1,l2)
}


/**
给定一个链表，旋转链表，将链表每个节点向右移动 k 个位置，其中 k 是非负数。
输入: 1->2->3->4->5->NULL, k = 2
输出: 4->5->1->2->3->NULL
 */
/*func rotateRight(head *ListNode, k int) *ListNode {

	var start = &ListNode{Data:0}
	var next = head
	for next != nil {
		if next {

		}
		next = next.Next
	}
	return nil
}
*/

