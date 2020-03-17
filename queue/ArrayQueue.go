package main

import "fmt"

type ArrayQueue struct {
	maxSize int
	array  [5]int
	head    int
	tail    int
}

func InitArray() *ArrayQueue  {
	return &ArrayQueue{
		maxSize:5,
		head:0,
		tail:0,
	}
}

func (a *ArrayQueue)IsFull() bool {
	return true
}

func (a *ArrayQueue)isTemp() bool  {
	return a.tail == a.head
}

func (a *ArrayQueue)Add(val int)  {

}

func (a *ArrayQueue)Get() int {
return 1
}

func main() {
	fmt.Println(1%5)
}