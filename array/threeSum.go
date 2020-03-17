package main

import (
	"fmt"
	"sort"
)

/**
15. 三数之和
给定一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？找出所有满足条件且不重复的三元组。

注意：答案中不可以包含重复的三元组。

例如, 给定数组 nums = [-1, 0, 1, 2, -1, -4]，

满足要求的三元组集合为：
[
  [-1, 0, 1],
  [-1, -1, 2]
]


 */
func threeSum(nums []int) [][]int {
	n, res := len(nums), make([][]int, 0)
	sort.Ints(nums)
	for i := 0; i < n; i += 1 {
		if i > 0 && nums[i] == nums[i-1] { // i=0的时候我们需要直接往下执行
			continue
		}
		l, r := i+1, n-1
		for l < r {
			tmp := nums[i] + nums[l] + nums[r]
			if tmp == 0 { // 如果三个数字加起来是0的话
				res = append(res, []int{nums[i], nums[l], nums[r]})
				l += 1
				r -= 1
				for l < r && nums[l] == nums[l-1] { // 重复数字我们不需要考虑
					l += 1
				}
				for l < r && nums[r] == nums[r+1] { // 重复数字我们不需要考虑
					r -= 1
				}
			} else if tmp > 0 { // 说明我们需要一个更小的数字
				r -= 1
			} else { // 说明我们需要一个更大的数字
				l += 1
			}
		}
	}
	return res
}

func main() {
	var nums =[]int{-1, 0, 1, 2, -1, -4}
//	sort.Ints(nums)
//	fmt.Println(nums)

	fmt.Println(threeSum(nums))
}