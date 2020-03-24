package main

import "sort"

/**
39. 组合总和
给定一个无重复元素的数组 candidates 和一个目标数 target ，找出 candidates 中所有可以使数字和为 target 的组合。

candidates 中的数字可以无限制重复被选取。

说明：

所有数字（包括 target）都是正整数。
解集不能包含重复的组合。 
示例 1:

输入: candidates = [2,3,6,7], target = 7,
所求解集为:
[
  [7],
  [2,2,3]
]
示例 2:

输入: candidates = [2,3,5], target = 8,
所求解集为:
[
  [2,2,2,2],
  [2,3,3],
  [3,5]
]


后续练习 216  78   90  401
 */
func combinationSum(candidates []int, target int) [][]int {
	sort.Ints(candidates)  //先排序
	var res [][]int
	var temp []int
	combinationSumDfs(candidates,temp,0,target,&res) //深度优先
	return res
}

func combinationSumDfs(candidates ,temp []int,left ,target int,res *[][]int)  {
	if target == 0{
		tmp:=make([]int,len(temp))
		copy(tmp,temp)
		*res = append(*res,tmp)
		return
	}
	for i:=left;i<len(candidates) ;i++  {  // left限定，形成分支
		if target < candidates[i]{  //剪枝
			return
		}
		combinationSumDfs(candidates,append(temp,candidates[i]),i,target -candidates[i],res)
	}
}