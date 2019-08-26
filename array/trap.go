package main

import "fmt"

/**
42. 接雨水
给定 n 个非负整数表示每个宽度为 1 的柱子的高度图，计算按此排列的柱子，下雨之后能接多少雨水。

上面是由数组 [0,1,0,2,1,0,1,3,2,1,2,1] 表示的高度图，在这种情况下，可以接 6 个单位的雨水（蓝色部分表示雨水）
示例:

输入: [0,1,0,2,1,0,1,3,2,1,2,1]
输出: 6
*/
func trap(height []int) int {
	max:=0
	maxIndex :=0

	for i:=0;i< len(height);i++  {
		if height[i] > max {
			max = height[i]
			maxIndex=i
		}
	}

	maxl :=0
	maxr :=0
	sum :=0

	for a:=0;a<maxIndex;a++{
		if maxl < height[a]{
			maxl = height[a]
		}else {
			sum +=maxl-height[a]
		}
	}

	for b:= len(height)-1;b> maxIndex;b--  {
		if maxr < height[b]{
			maxr = height[b]
		}else {
			sum +=maxr-height[b]
		}
	}
	return sum
}

func main() {
	var nums =[]int{0,1,0,2,1,0,1,3,2,1,2,1}
	fmt.Println(trap(nums))
}