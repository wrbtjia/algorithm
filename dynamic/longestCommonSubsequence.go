package main

import "fmt"

/**
1143. 最长公共子序列
给定两个字符串 text1 和 text2，返回这两个字符串的最长公共子序列。

一个字符串的 子序列 是指这样一个新的字符串：它是由原字符串在不改变字符的相对顺序的情况下删除某些字符（也可以不删除任何字符）后组成的新字符串。
例如，"ace" 是 "abcde" 的子序列，但 "aec" 不是 "abcde" 的子序列。两个字符串的「公共子序列」是这两个字符串所共同拥有的子序列。

若这两个字符串没有公共子序列，则返回 0。

 

示例 1:

输入：text1 = "abcde", text2 = "ace"
输出：3
解释：最长公共子序列是 "ace"，它的长度为 3。
示例 2:

输入：text1 = "abc", text2 = "abc"
输出：3
解释：最长公共子序列是 "abc"，它的长度为 3。
示例 3:

输入：text1 = "abc", text2 = "def"
输出：0
解释：两个字符串没有公共子序列，返回 0。

 */

func longestCommonSubsequence(text1 string, text2 string) int {
	dp := make([]int, len(text2)+1)

	for i := 1; i <= len(text1); i++ {
		last := 0
		for j := 1; j <= len(text2); j++ {
			tmp := dp[j] // tmp 记录的是二维dp矩阵正上方的值
			if text1[i-1] == text2[j-1] {
				dp[j] = last + 1 // last 记录的是二维dp矩阵左上方的值
			} else {
				dp[j] = maxs(tmp, dp[j-1])
			}
			last = tmp
		}
	}

	return dp[len(text2)]

}

func maxs(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {
	fmt.Println(longestCommonSubsequence("abc","def"))
}