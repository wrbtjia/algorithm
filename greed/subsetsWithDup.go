package main

import (
	"fmt"
	"sort"
)

/**
给定一个可能包含重复元素的整数数组 nums，返回该数组所有可能的子集（幂集）。

说明：解集不能包含重复的子集。

输入: [1,2,2]
输出:
[
  [2],
  [1],
  [1,2,2],
  [2,2],
  [1,2],
  []
]
 */



func subsetsWithDup(nums []int) [][]int {
	var res [][]int
	var tmp []int

	if len(nums) == 0 {
		return res
	}
	//排序
	sort.Ints(nums)
	
	res= append(res, tmp)
	generate(0,nums,tmp,&res)
	return res
}

func generate(i int,nums []int, tmp []int, res *[][]int)  {
	if i >= len(nums){
		
		return
	}
	tmp = append(tmp, nums[i])

	*res = append(*res, tmp)




	generate(i+1,nums,tmp,res)

	tmp = tmp[1:]
	generate(i+1,nums,tmp,res)
}

func main() {
	var nums = []int{1,2,2}

	fmt.Println(subsetsWithDup(nums))


	 maps :=make(map[interface{}]int)

	 maps[nums] = 1

	 fmt.Println(maps)



}