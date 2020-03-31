package main

import (
	"algorithm/tools"
	"fmt"
)


/**
11. 盛最多水的容器
给你 n 个非负整数 a1，a2，...，an，每个数代表坐标中的一个点 (i, ai) 。
在坐标内画 n 条垂直线，垂直线 i 的两个端点分别为 (i, ai) 和 (i, 0)。
找出其中的两条线，使得它们与 x 轴共同构成的容器可以容纳最多的水。

 */


func maxArea(height []int) int {
	val ,l:=0,0
	r := len(height)-1
	for l < r  {
		val = tools.MaxInter(val,tools.MinInter(height[l],height[r]) * (r-l))
		if height[l] < height[r] {
			l++
		}else {
			r--
		}
	}
	return val
}

func main() {
	nums:=[]int{1,8,6,2,5,4,8,3,7}
	fmt.Println(maxArea(nums))
}

