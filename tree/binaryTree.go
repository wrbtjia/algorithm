package main

import (
	"fmt"
)

type tree struct {
	data int
	left,right *tree
}

func CreateTree(val int) *tree{
	return &tree{data:val}
}


func Insert(e *tree,val int) *tree {

	if e == nil {
		e = &tree{data:val}
		return e
	}

	if e.data > val {
		e.left = Insert(e.left,val)
	}else if e.data < val{
		e.right = Insert(e.right,val)
	}
	return e
}
//查询时记录下标
var i int = 0
func toArrayNode(t *tree,datas []int)  {

	if t.left != nil{

		toArrayNode(t.left,datas)
	}
	datas[i] = t.data
	i++

	if t.right != nil {

		toArrayNode(t.right,datas)
	}
	fmt.Println(t.data)
}


func Show()  {
	fmt.Println()
}

func main() {
	t := CreateTree(6)

	t = Insert(t,8)
	t = Insert(t,2)
	t = Insert(t,7)
	t = Insert(t,1)
	t = Insert(t,9)
	t = Insert(t,4)
	t = Insert(t,5)



	var datas = make([]int, 8)
	toArrayNode(t,datas)

	fmt.Println(datas)

	
}
