package main



type MyQueue struct {
	InputStack  []int //输入栈
	OutputStack []int //输出栈
}


/**

使用栈实现队列的下列操作：

push(x) -- 将一个元素放入队列的尾部。
pop() -- 从队列首部移除元素。
peek() -- 返回队列首部的元素。
empty() -- 返回队列是否为空。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/implement-queue-using-stacks
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
Initialize your data structure here. */

func Constructor() MyQueue {
	queue := MyQueue{[]int{}, []int{}}
	return queue
}


/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int)  {
	this.InputStack = append(this.InputStack, x)
}


/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
	if len(this.OutputStack) != 0 {
		x := this.OutputStack[len(this.OutputStack)-1]
		this.OutputStack =  this.OutputStack[:len(this.OutputStack)-1]
		return x
	}

	if len(this.InputStack) != 0 {
		for i := len(this.InputStack); i>0; i--{
			this.OutputStack = append(this.OutputStack, this.InputStack[i-1])
			this.InputStack = this.InputStack[:len(this.InputStack)-1]
		}
		x := this.OutputStack[len(this.OutputStack) -1]
		this.OutputStack = this.OutputStack[:len(this.OutputStack) -1]
		return x

	}
	return 0
}


/** Get the front element. */
func (this *MyQueue) Peek() int {
	if len(this.OutputStack) != 0 {
		return this.OutputStack[len(this.OutputStack) -1]
	}

	if len(this.InputStack) != 0 {
		for i := len(this.InputStack); i>0; i--{
			this.OutputStack = append(this.OutputStack, this.InputStack[i-1])
			this.InputStack = this.InputStack[:len(this.InputStack)-1]
		}
		return this.OutputStack[len(this.OutputStack) -1]
	}
	return 0
}


/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
	if len(this.InputStack)+len(this.OutputStack) == 0 {
		return true
	}
	return false
}


func main() {
	queue := Constructor()
	queue.Push(1)
	queue.Push(2)
	queue.Push(3)
	queue.Push(4)
	queue.Push(5)
	_ = queue.Peek()
	_ = queue.Pop()
	_ = queue.Pop()
	_ = queue.Pop()
	_ = queue.Pop()
	_ = queue.Empty()
}