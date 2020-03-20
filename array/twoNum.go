package main

import (
	"fmt"
)

func twoSum(nums []int, target int) []int {
	tarMap := make(map[int]int)
	for a, b := range nums {
		othVal := target - b
		if v, ok := tarMap[othVal]; ok {

			return []int{v, a}
		} else {

			tarMap[b] = a
		}
	}
	return []int{}

}

func main() {
	var nums = []int{1, 3, 4, 2}
	s := twoSum(nums, 6)

	fmt.Println(s)

}
