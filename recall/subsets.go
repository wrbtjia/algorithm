package main

import "fmt"

/**
78. 子集
给定一组不含重复元素的整数数组 nums，返回该数组所有可能的子集（幂集）。
说明：解集不能包含重复的子集。
示例:
输入: nums = [1,2,3]
输出:
[
  [3],
  [1],
  [2],
  [1,2,3],
  [1,3],
  [2,3],
  [1,2],
  []
]
 */

func subsets(nums []int) [][]int {
	var res [][]int
	if len(nums)==0 {
		return res
	}
	var temp []int
	getSubsets(&res,nums,temp,0)
	return res
}
func getSubsets(res *[][]int,nums []int,temp []int,n int){
	if n==len(nums) {
		cop:=make([]int,len(temp))
		copy(cop,temp)
		*res= append(*res, cop)
		return
	}
	temp= append(temp, nums[n])
	getSubsets(res,nums,temp,n+1)
	temp2:=temp
	temp2=temp2[:len(temp2)-1]
	temp= temp2
	getSubsets(res,nums,temp,n+1)
}

func main() {
	nums:=[]int{1,2,3}
	fmt.Println(subsets(nums))
}
