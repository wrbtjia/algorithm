package main


import (
	"errors"
	"fmt"
)

type loop struct {
	maxSize int
	array [5]int
	head  int
	tail   int
}

func CreateQueue() *loop  {
	return &loop{maxSize:5,head:0,tail:0}
}

func AddQueue(q *loop ,val int) error {
	if IsFull(q) {
		return errors.New("队列已满")
	}
	q.array[q.tail] = val
	q.tail = (q.tail+1)%q.maxSize
	return nil
}

func GetQueue(q *loop) (int,error) {
	if IsTemp(q) {
		return -1,errors.New("队列已空")
	}
	val := q.array[q.head]
	q.head = (q.head+1)%q.maxSize
	return val,nil
}

func IsTemp(q *loop) bool {
		return q.head == q.tail

}

func IsFull(q *loop) bool {
	return (q.tail+1)%q.maxSize == q.head
}

func Size(q *loop)  int{
	return (q.tail+q.maxSize-q.head)%q.maxSize
}

func Show(q *loop)  {
	size:=Size(q)
fmt.Println(size)
	tem :=q.head
	for i:=0; i<size;i++  {
		fmt.Printf("array[%d]=%d \t",tem,q.array[tem])
		tem =(tem+1)%q.maxSize
	}
}

func main() {
	q:=CreateQueue()

	AddQueue(q,1)
	AddQueue(q,2)
	AddQueue(q,3)
	AddQueue(q,4)


	Show(q)

	var res int
	res ,_= GetQueue(q)
	res ,_= GetQueue(q)

	fmt.Println(res)
	AddQueue(q,6)
	AddQueue(q,8)

	res ,_= GetQueue(q)
	res ,_= GetQueue(q)
	res ,_= GetQueue(q)

	AddQueue(q,5)
	AddQueue(q,6)
	Show(q)
}
