package main

import "fmt"

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
