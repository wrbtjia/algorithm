package main

import "fmt"

//归并排序
//时间复杂度 O(n log n)

func main() {
	array:=[]int{1,9,2,8,3,7,4,6,5,10}

	fmt.Println(MergeSort(array))
}

func MergeSort(array []int) []int  {

	length:=len(array)
	if length <=1 {
		return array
	}
	mid:=len(array)/2
	leftarr:=MergeSort(array[:mid])
	rightarr:=MergeSort(array[mid:])

	return merge(leftarr,rightarr)

}

func merge(leftarr,rightarr []int)[]int  {
	leftindex:=0  //左索引
	rightindex:=0; //右索引
	lastarr :=[]int{}  //组合的新数组
	for leftindex < len(leftarr) && rightindex <len(rightarr){
		if leftarr[leftindex] < rightarr[rightindex]{
			lastarr = append(lastarr, leftarr[leftindex])
			leftindex++
		}else if leftarr[leftindex] > rightarr[rightindex] {
			lastarr = append(lastarr,  rightarr[rightindex])
			rightindex++
		}else{
			lastarr = append(lastarr, leftarr[leftindex])
			lastarr = append(lastarr,  rightarr[rightindex])
			leftindex++
			rightindex++
		}
	}
	for leftindex < len(leftarr) {
		lastarr = append(lastarr, leftarr[leftindex])
		leftindex++
	}
	for rightindex <len(rightarr) {
		lastarr = append(lastarr,  rightarr[rightindex])
		rightindex++
	}

	return lastarr
	
}

