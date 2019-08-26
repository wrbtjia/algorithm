package main

import "fmt"

/**
209
给定一个含有 n 个正整数的数组和一个正整数 s ，找出该数组中满足其和 ≥ s 的长度最小的连续子数组。如果不存在符合条件的连续子数组，返回 0。
输入: s = 7, nums = [2,3,1,2,4,3]
输出: 2
解释: 子数组 [4,3] 是该条件下的长度最小的连续子数组。
 */

func minSubArrayLen(s int, nums []int) int {

	l:=0
	r:=-1
	sum:=0
	res := len(nums)+1

	for l< len(nums) {
		if r+1 < len(nums) && sum < s {
			r++
			sum += nums[r]
		}else {

			sum -= nums[l]
			l++
		}
		if sum >=s {
			res = min(res,r-l+1)
		}
	}
	if res == len(nums)+1 {
		return 0
	}
	return res
}
func min(a,b int)int{
	if a<b{
		return a
	}else{
		return b
	}
}
func main() {
	var nums = []int{2,3,1,2,4,3}
	fmt.Println(minSubArrayLen(7,nums))


}

