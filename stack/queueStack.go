package main

import "fmt"

type MyStack struct {
	q1 []int

}


/** Initialize your data structure here. */
func Constructors() MyStack {

	stack := MyStack{q1:[]int{}}
	return stack
}


/** Push element x onto stack. */
func (this *MyStack) Push(x int)  {

	this.q1 = append(this.q1, x)
}


/** Removes the element on top of the stack and returns that element. */
func (this *MyStack) Pop() int {
	if len(this.q1) != 0{
		res := this.q1[len(this.q1)-1]
		this.q1 = this.q1[:len(this.q1)-1]
		return res
	}
	return 0
}


/** Get the top element. */
func (this *MyStack) Tops() int {
	if len(this.q1) != 0 {
		res := this.q1[len(this.q1)-1]
		return res
	}
	return 0
}


/** Returns whether the stack is empty. */
func (this *MyStack) Empty() bool {

	return len(this.q1)  == 0
}

func main() {
	obj := Constructors();
	 obj.Push(1);
	 obj.Push(2);
	q1 := obj.Tops()
	 obj.Push(3);
	q3 := obj.Tops()
	/* obj.Push(4);
	 obj.Push(5);
	 q := obj.Pop();

	 q2 := obj.Empty();*/

	 fmt.Println(q1,q3)
}