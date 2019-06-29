package main

import (
	"fmt"
	"sort"
)

/**
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

	sort.Ints(candidates)
	res := [][]int{}
	generate2(&res, candidates, []int{}, target, 0, 0)
	return res
}

func generate2(res *[][]int, candidates, newRow []int, target, start, sum int) {
	if sum == target {
		(*res) = append((*res), newRow)
		return
	} else if sum > target {
		return
	}
	length := len(candidates)
	for i := start; i < length; i++ {
		item := candidates[i]
		row := make([]int, len(newRow)+1)
		copy(row, newRow)
		row[len(row)-1] = item
		generate2(res, candidates, row, target, i+1, sum+item)

		next := i
		for candidates[next] == candidates[i] {
			next += 1
			if next == length {
				break
			}
		}
		i = next - 1
	}
	return
}

func main() {
	var candidates = []int{10,1,2,7,6,1,5}

	fmt.Println(combinationSum2(candidates,8))
}