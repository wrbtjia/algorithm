package main

import "fmt"

/**
3
给定一个字符串，请你找出其中不含有重复字符的 最长子串 的长度。
输入: "abcabcbb"
输出: 3
解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
 */
func lengthOfLongestSubstring(s string) int {
	freq:=[256]int{0}

	l:=0
	r:=-1
	res :=0
	for l < len(s) {
		if r+1 < len(s) && freq[s[r+1]] == 0 {
			r++
			freq[s[r]] ++

		}else {
				freq[s[l]] --
				l++
		}
		res = max(res,r-l+1)
	}
	return res
}
func max(a,b int)int{
	if a>b{
		return a
	}else{
		return b
	}
}



func main() {
fmt.Println(lengthOfLongestSubstring("abcabcbb"))



}