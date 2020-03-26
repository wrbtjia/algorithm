package main

import "fmt"

/**
300. 最长上升子序列
给定一个无序的整数数组，找到其中最长上升子序列的长度。

示例:

输入: [10,9,2,5,3,7,101,18]
输出: 4
解释: 最长的上升子序列是 [2,3,7,101]，它的长度是 4。

 */


func lengthOfLIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	//[10,9,2,5,3,7,101,18]
	dp := make([]int, len(nums))
	res := dp[0]
	for i := 0; i < len(nums); i++ {
		last := 0
		for j := 0; j < i; j++ {
			if last < dp[j] && nums[j] < nums[i] {
				last = dp[j]
			}
		}
		dp[i] = last + 1
		if dp[i] > res {
			res = dp[i]
		}
	}
	return res
}

func main() {
	nums:=[]int{10,9,2,5,3,7,101,18}
	fmt.Println(lengthOfLIS(nums))
}