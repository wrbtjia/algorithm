package main

import "fmt"

/**
53. 最大子序和
给定一个整数数组 nums ，找到一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。

示例:

输入: [-2,1,-3,4,-1,2,1,-5,4],
输出: 6
解释: 连续子数组 [4,-1,2,1] 的和最大，为 6。

 */


func maxSubArray(nums []int) int {
	if len(nums) < 1 {
		return 0;
	} else {
		dp:=nums[0]
		sum:=dp
		for _,v := range nums[1:] {
			if dp>0 {
				dp=dp+v
			}else {
				dp = v
			}
			if sum < dp{
				sum = dp
			}
		}
		return sum;
	}
}

func maxSubArray2(nums []int) int {
	if len(nums) < 1{
		return 0
	}
	dp:=nums[0]
	res:=0
	for i:=1;i<len(nums);i++ {
		dp = maxx(nums[i],dp+nums[i])
		res = maxx(dp,res)
	}
	return res
}

func maxx(i,j int) int {
	if i>j{
		return i
	}else {
		return j
	}

}

func main() {
	nums:=[]int{2,1,-3,4,-1,2,1,-5,4}
	fmt.Println(maxSubArray2(nums))
}