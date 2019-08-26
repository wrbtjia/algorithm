package main

import "fmt"


/**
239. 滑动窗口最大值
给定一个数组 nums，有一个大小为 k 的滑动窗口从数组的最左侧移动到数组的最右侧。你只可以看到在滑动窗口 k 内的数字。滑动窗口每次只向右移动一位。

返回滑动窗口最大值。

 */
func maxSlidingWindow(nums []int, k int) []int {

	result := []int{}
	if len(nums) == 0 {
		return result
	}

	window := []int{}
	for i, x := range nums {
		if i >= k && window[0] <= i-k {
			window = window[1:]
		}
		for len(window) != 0 && nums[window[len(window)-1]] <= x {
			window = []int{}
		}
		window = append(window, i)
		if i >= k-1 {
			result = append(result, nums[window[0]])
		}
	}
	return result
}

func main() {
	var nums = []int{1, 3, 1, 2, 0, 5}
	var k int = 3
	res := maxSlidingWindow(nums, k)

	fmt.Println(res)
}
