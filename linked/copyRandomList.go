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

	a.Next = b
	b.Next = c
	c.Next = d
	d.Next = e

	a.Random = c
	b.Random = d
	c.Random = c
	e.Random = d

	p :=CopyRandomList(a)
	Prints(p)
}
