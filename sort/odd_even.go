package main

import "fmt"

//奇偶排序

func main() {

	array:=[]int{1,9,2,8,3,7,4,6,5,10}


	isSorted := false
	for isSorted == false  {
		isSorted = true
		for i:=1;i<len(array)-1 ;i=i+2  {
			if array[i]<array[i+1] {
				array[i],array[i+1]=array[i+1],array[i]
				isSorted = false
			}
		}
		for i:=0;i<len(array)-1 ;i=i+2  {
			if array[i]<array[i+1] {
				array[i],array[i+1]=array[i+1],array[i]
				isSorted = false
			}
		}
	}
	fmt.Println(array)
}


