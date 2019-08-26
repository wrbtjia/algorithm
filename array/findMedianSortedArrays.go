package main

import (
	"sort"
)

/**
4. 寻找两个有序数组的中位数
给定两个大小为 m 和 n 的有序数组 nums1 和 nums2。

请你找出这两个有序数组的中位数，并且要求算法的时间复杂度为 O(log(m + n))。

你可以假设 nums1 和 nums2 不会同时为空。

示例 1:

nums1 = [1, 3]
nums2 = [2]

则中位数是 2.0
示例 2:

nums1 = [1, 2]
nums2 = [3, 4]

则中位数是 (2 + 3)/2 = 2.5

 */
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	total := len(nums1) + len(nums2)
	nums := make([]int, 0, total)
	i, j := 0, 0
	for len(nums) != total {
		if j >= len(nums2) {
			nums = append(nums, nums1[i])
			i++
		} else if i >= len(nums1) {
			nums = append(nums, nums2[j])
			j++
		} else if nums1[i] < nums2[j] {
			nums = append(nums, nums1[i])
			i++
		} else if nums1[i] > nums2[j] {
			nums = append(nums, nums2[j])
			j++
		} else {
			nums = append(nums, nums1[i], nums2[j])
			i++
			j++
		}
	}

	return float64(nums[(total+1)/2-1]+nums[(total+2)/2-1]) / 2
}


func findMedianSortedArrays2(nums1 []int, nums2 []int) float64 {

	nums1 = append(nums1, nums2...)
	sort.Ints(nums1)
	bln := len(nums1)%2
	val :=0.0
	if bln != 0 {
		val = float64(nums1[len(nums1)/2])
	}else {
		val = float64(nums1[len(nums1)/2]+nums1[len(nums1)/2-1])/2
	}

	return val

	/**
	执行结果：通过 显示详情
	执行用时 : 28 ms, 在所有 Go 提交中击败了
	62.81%
	的用户
	内存消耗 :
	6 MB
	, 在所有 Go 提交中击败了
	26.40%
	的用户
	 */
}


func main() {
	var nums1 =[]int{1,3}
	var nums2 =[]int{2}
	findMedianSortedArrays2(nums1,nums2)
}