package main

import "fmt"

//鸡尾酒排序  又称双向冒泡
func main() {
	arr:=[]int {1,9,2,8,3,7,4,6,5,10}
	for i:=0;i<len(arr)/2 ;i++  { //每次循环正向一次 反向一次
		left :=0
		right :=len(arr) -1
		for left < right {   //结束条件
			if arr[left] > arr[left+1] {
				arr[left],arr[left+1]=arr[left+1],arr[left]
			}
			left++
			if arr[right] < arr[right-1]{
				arr[right] , arr[right-1] = arr[right-1],arr[right]
			}
			right --
		}

		
	}

	fmt.Println(arr)
}

