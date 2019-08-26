package main

import "fmt"

/**
46
给定一个没有重复数字的序列，返回其所有可能的全排列。
输入: [1,2,3]
输出:
[
  [1,2,3],
  [1,3,2],
  [2,1,3],
  [2,3,1],
  [3,1,2],
  [3,2,1]
]

 */



func permute(nums []int) [][]int {
	st := make([]bool, len(nums))
	var res [][]int
	 item :=make([]int,len(nums))
	allpermute(nums,0,item,&res,st)
	return res
}


func allpermute(nums []int,index int,item []int,res *[][]int, st []bool){

	if index == len(nums){
		tmp := make([]int, len(nums))
		copy(tmp, item)
		*res = append(*res, tmp)
		return
	}

	for i:=0;i< len(nums); i++  {
		if !st[i]  {
			st[i] = true

			item[index] = nums[i]
			allpermute(nums,index+1,item,res,st)

			st[i] = false
		}

	}

}

func main() {
	var nums =[]int{5,4,6,2}
	fmt.Println(permute(nums))
}


