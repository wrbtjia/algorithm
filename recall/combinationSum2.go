package main

import (
	"fmt"
	"sort"
)

/**
40. 组合总和 II
给定一个数组 candidates 和一个目标数 target ，找出 candidates 中所有可以使数字和为 target 的组合。

candidates 中的每个数字在每个组合中只能使用一次。

输入: candidates = [10,1,2,7,6,1,5], target = 8,
所求解集为:
[
  [1, 7],
  [1, 2, 5],
  [2, 6],
  [1, 1, 6]
]

 */
func combinationSum2(candidates []int, target int) [][]int {

	sort.Ints(candidates) //快排，懒得写
	res := [][]int{}
	generate2(candidates, nil, 0,target,&res) //深度优先
	return res

}

func generate2 (candidates, newRow []int, left,target int,res *[][]int) {
	if target == 0{
		ms :=make([]int,len(newRow))
		copy(ms,newRow)
		*res = append(*res,ms)
		return
	}
	for i:=left;i<len(candidates);i++ {
		if i !=left && candidates[i] == candidates[i-1] {
			continue
		}
		if target < candidates[i]{
			return
		}
		generate2(candidates,append(newRow,candidates[i]),i+1,target - candidates[i],res)
	}
}

func main() {
	var candidates = []int{10,1,2,7,6,1,5}

	fmt.Println(combinationSum2(candidates,8))
}