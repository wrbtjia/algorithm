package main

import "fmt"

func main() {
	//选择排序，从后面找个最小的放在前面的位置，从小到大排序,时间复杂度O(N^2)
	array := [11]int{56,45,9,16,2,89,78,34,102,56,99}
	length:= len(array)


	for i:=0;i< length-1;i++{
		position := i
		for j:=i+1;j<length-1;j++ {
			if array[j] < array[position] {
				position = j
			}
		}

		temp := array[i]

		array[i]= array[position]
		array[position] = temp
	}

	fmt.Println(array)
}