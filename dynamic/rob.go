package main

import "fmt"

/**
198. 打家劫舍
你是一个专业的小偷，计划偷窃沿街的房屋。每间房内都藏有一定的现金，影响你偷窃的唯一制约因素就是相邻的房屋装有相互连通的防盗系统，如果两间相邻的房屋在同一晚上被小偷闯入，系统会自动报警。

给定一个代表每个房屋存放金额的非负整数数组，计算你在不触动警报装置的情况下，能够偷窃到的最高金额。

示例 1:

输入: [1,2,3,1]
输出: 4
解释: 偷窃 1 号房屋 (金额 = 1) ，然后偷窃 3 号房屋 (金额 = 3)。
     偷窃到的最高金额 = 1 + 3 = 4 。
示例 2:

输入: [2,7,9,3,1]
输出: 12
解释: 偷窃 1 号房屋 (金额 = 2), 偷窃 3 号房屋 (金额 = 9)，接着偷窃 5 号房屋 (金额 = 1)。
     偷窃到的最高金额 = 2 + 9 + 1 = 12 。
 */
func rob(nums []int) int {
	pre, cur := 0, 0
	for _, m := range nums {
	//	curMax, preMax = max(preMax + m, curMax), curMax
		temp := max(pre + m, cur)
		pre = cur
		cur =temp
	}
	return cur
}
func max(x, y int) int {
	if x > y {return x}
	return y
}

func rob2(nums []int) int {
	if len(nums) == 0{
		return 0
	}
//nums:=[]int{2,7,9,3,1}
	sum:=0  //表示下标从0到i-2的累计最大值nums[i]
	for i:=1;i<len(nums);i++ {
		nums[i] += sum
		if nums[i - 1] > sum {
			sum = nums[i - 1]  //在下标为i时更新i-1的max，下一轮循环i=i+1,使用的是i-1的max，这样才是正确的（不能连续，避免报警）
		}

	}


	return max(sum,nums[len(nums)-1])


/*	l := len(nums)
	max := 0 //max表示下标从0到i-2的累计最大值
	if l == 0 {
		return 0
	}
	if l == 1 {
		return nums[0]
	}
	for i:=1; i<l; i++ {
		nums[i] += max
		if nums[i-1] > max {
			max = nums[i-1] //在下标为i时更新i-1的max，下一轮循环i=i+1,使用的是i-1的max，这样才是正确的（不能连续，避免报警）
		}
	}
	if max > nums[l-1] {
		return max
	}else {
		return nums[l-1]
	}*/

}


func main() {
	nums:=[]int{2,7,9,3,1}
	fmt.Println(rob(nums))

}