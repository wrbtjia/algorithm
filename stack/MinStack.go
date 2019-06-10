package main



type MinStack struct {
	data []int
	min []int
}


/** initialize your data structure here. */
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