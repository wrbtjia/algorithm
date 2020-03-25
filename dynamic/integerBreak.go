package main

import "fmt"

/**
343. 整数拆分
给定一个正整数 n，将其拆分为至少两个正整数的和，并使这些整数的乘积最大化。 返回你可以获得的最大乘积。

示例 1:

输入: 2
输出: 1
解释: 2 = 1 + 1, 1 × 1 = 1。
示例 2:

输入: 10
输出: 36
解释: 10 = 3 + 3 + 4, 3 × 3 × 4 = 36。
 */

func integerBreak(n int) int {
	return method_dp(n)
}

/*
原因：n <=3 时 最大乘积dp[n] < n，因此需要使用max(n, dp[n])
dp[i] = getMax(dp[i], i > j ==> j * max(i-j, dp[i-j]))

后续练习 279  91  62  63
*/
func method_dp(n int) int {
	dp := make([]int, n+1)
	dp[1] = 1
	dp[2] = 1
	for i := 3; i <=n; i++ {
		for j := 1; j < i; j++ {
			dp[i] = getMax(dp[i], j * getMax(i-j, dp[i-j]))
		}
	}
	// fmt.Println(dp)
	return dp[n]
}

func getMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func main() {
	fmt.Println(integerBreak(10))
}
