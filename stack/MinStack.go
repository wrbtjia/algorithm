package main



type MinStack struct {
	data []int
	min []int
}


/** initialize your data structure here.
设计一个支持 push，pop，top 操作，并能在常数时间内检索到最小元素的栈。

push(x) -- 将元素 x 推入栈中。
pop() -- 删除栈顶的元素。
top() -- 获取栈顶元素。
getMin() -- 检索栈中的最小元素。

示例
MinStack minStack = new MinStack();
minStack.push(-2);
minStack.push(0);
minStack.push(-3);
minStack.getMin();   --> 返回 -3.
minStack.pop();
minStack.top();      --> 返回 0.
minStack.getMin();   --> 返回 -2.

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/min-stack
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func Constructor1() MinStack {
	return  MinStack{data:[]int{},min:[]int{}}
}


func (this *MinStack) Push(x int)  {

	this.data = append(this.data, x)
	if len(this.min) == 0 {
		this.min = append(this.min, x)
	}else {
		res := this.min[len(this.min)-1]
		if res < x {
			this.min = append(this.min, res)
		}else {
			this.min = append(this.min, x)
		}

	}
}


func (this *MinStack) Pop()  {
	l := len(this.data)
	this.data = this.data[:l-1]
	this.min = this.min[:l-1]
}


func (this *MinStack) Top() int {
	res := this.data[len(this.data) -1 ]

	return res
}


func (this *MinStack) GetMin() int {
	res := this.min[len(this.min)-1]
	return  res
}

func main() {

}