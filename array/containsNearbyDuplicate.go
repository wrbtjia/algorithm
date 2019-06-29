package main

import "fmt"

/**
219
给定一个整数数组和一个整数 k，判断数组中是否存在两个不同的索引 i 和 j，使得 nums [i] = nums [j]，并且 i 和 j 的差的绝对值最大为 k。

输入: nums = [1,2,3,1], k = 3
输出: true

输入: nums = [1,0,1,1], k = 1
输出: true

输入: nums = [1,2,3,1,2,3], k = 2
输出: false
 */
func containsNearbyDuplicate(nums []int, k int) bool {
		set := make(map[int]int)

	for i:=0;i< len(nums);i++  {
		if _,ok := set[nums[i]]; ok {
			return true
		}
		set[nums[i]]= i

		if len(set) == k+1 {
			delete(set,nums[i-k])
		}
	}

	return false
}

func main() {
	var nums = []int{1,2,3,1}

	fmt.Println(containsNearbyDuplicate(nums,3))
}