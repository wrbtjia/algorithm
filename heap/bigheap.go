package main

import "fmt"

type BigHeap struct {
	data []int
}

func NewBigHeap() *BigHeap{
	return &BigHeap{
		data:[]int{},
		
	}
}

func (b *BigHeap ) GetSize() int {
	return len(b.data)
}

func (b *BigHeap ) isEmpty() bool {
	return len(b.data) == 0
}


// 返回完全二叉树的数组表示中，一个索引所表示的元素的父亲节点的索引
func (b *BigHeap ) parent(index int) int {
	return (index -1)/2
}


// 返回完全二叉树的数组表示中，一个索引所表示的元素的左孩子节点的索引
func (b *BigHeap ) leftChild(index int) int {
	return index * 2 + 1
}

func (b *BigHeap ) rightChild (index int) int {
	return index * 2 + 2;
}

func (b *BigHeap ) swap(i,j int)  {
	if i<0 || i > len(b.data) {
		return
	}

	b.data[i],b.data[j] = b.data[j],b.data[i]
}

func (b *BigHeap ) add(index int)  {
	b.data = append(b.data, index)
	b.siftUp(len(b.data) - 1);

}

func (b *BigHeap ) siftUp(i int)  {
	//特性2：比较插入值和其父结点的大小关系，小于父结点则用父结点替换当前值，index位置上升为父结点
	// 当上浮元素大于父亲，继续上浮。并且不能上浮到0之上
	// 直到i 等于 0 或 比 父亲节点小了。
	for i > 0 && b.data[i] < b.data[b.parent(i)] {
		b.swap(i,b.parent(i))
		i = b.parent(i)
	}

}

func (b *BigHeap) finMin() int {
	return b.data[0]
}

func (b *BigHeap) extractMin() int {
	min := b.finMin()
	b.swap(0, len(b.data)-1)

	b.data = b.data[:len(b.data)-1]

	b.siftDown(0)
	return min
}

func (b *BigHeap)siftDown(i int)  {
	for b.leftChild(i) < len(b.data) {
		j := b.leftChild(i)
		if j+1 < len(b.data) && b.data[j+1] < b.data[j] {
				j++
		}
		if b.data[i]  <= b.data[j] {
			break
		}

		b.swap(i,j)
		i = j
	}
}


func main() {
	var arr =[]int{8,4,6,1,7,3,5}

	head :=NewBigHeap()
	for _,e := range arr {
		head.add(e)
	}


	fmt.Println(head.data)

	head.extractMin()

	fmt.Println(head.data)


}
