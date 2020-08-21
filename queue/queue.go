package main

import (
	"errors"
	"fmt"
)

type Queue struct {
	maxSize int
	array [5]int
	front int
	rear int
}

func (q *Queue) AddQueue(val int) (err error) {
	if q.maxSize -1 == q.rear {
		return errors.New(" queue is full")
	}
	q.rear++
	q.array[q.rear] = val
	return
}

func (q *Queue)GetQueue () (val int ,err error)  {
	if q.front == q.rear {
		return -1, errors.New("queue is temp")
	}
	q.front ++
	val = q.array[q.front]
	return val,nil
}

func (q *Queue)ShowQueue()  {
	for i:=q.front +1;i<=q.rear;i++ {
		fmt.Printf("array[%d]=%d \t",i,q.array[i])
	}
}

func main() {
	queue := &Queue{
		maxSize:5,
		front:-1,
		rear:-1,
	}

	var value int


	queue.AddQueue(1)
	queue.AddQueue(2)
	queue.AddQueue(3)
	queue.AddQueue(4)
	queue.AddQueue(5)

	queue.ShowQueue()
	fmt.Println()
	value,_=queue.GetQueue()
fmt.Println(value)
	value,_=queue.GetQueue()
	fmt.Println(value)
	value,_=queue.GetQueue()
	fmt.Println(value)
	queue.ShowQueue()

}