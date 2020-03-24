
package main

import "fmt"

/**
216. 组合总和 III
找出所有相加之和为 n 的 k 个数的组合。组合中只允许含有 1 - 9 的正整数，并且每种组合中不存在重复的数字。

说明：

所有数字都是正整数。
解集不能包含重复的组合。 
示例 1:

输入: k = 3, n = 7
输出: [[1,2,4]]
示例 2:

输入: k = 3, n = 9
输出: [[1,2,6], [1,3,5], [2,3,4]]
 */

func combinationSum3(k int, n int) [][]int {
	var res [][]int
	var temp []int

	combinationSum3Dfs(k,1,n,temp,&res)

	return  res
}

func combinationSum3Dfs(k int,left,n int,temp []int,res *[][]int)  {
	if n==0 && len(temp) == k{
		ms :=make([]int,len(temp))
		copy(ms,temp)
		*res = append(*res,ms)
		return
	}

	for i:=left;i<=9;i++ {

		combinationSum3Dfs(k,i+1,n-i,append(temp,i),res)
	}

}

func main() {
	fmt.Println(combinationSum3(3,9))

}