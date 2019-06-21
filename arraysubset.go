package main

import "fmt"


/**
获取一个数组的所有子集

 */
func subset(i int,nums []int,item []int,resutl *[][]int)   {
	if i>= len(nums){
		return
	}
	item = append(item, nums[i])
	*resutl = append(*resutl, item)
	subset(i+1,nums,item,resutl)
}



func main() {
	num :=[]int{1,2,3}
	var item =[]int{}

	resutl := [][]int{}
	subset(0,num,item,&resutl)

	fmt.Println(resutl)
}