package main

import "fmt"

/**
131. 分割回文串
给定一个字符串 s，将 s 分割成一些子串，使每个子串都是回文串。

返回 s 所有可能的分割方案。

示例:

输入: "aab"
输出:
[
  ["aa","b"],
  ["a","a","b"]
]
 */


func partition(s string) [][]string {
	str := []byte(s)
	res := make([][]string,0)
	backTracking(str,[]string{},&res)
	return res
}

func backTracking(s []byte,temp []string,res *[][]string){
	if len(s) == 0{
		tmp := make([]string,len(temp))
		copy(tmp,temp)

		*res = append(*res,tmp)
		return
	}
	for i:=1;i<=len(s);i++{
		if isPali(string(s[:i])){
			fmt.Println(string(s[:i]))
			backTracking(s[i:],append(temp, string(s[:i])),res)
		}
	}
}
func isPali(s string) bool{

	i :=0
	j:=len(s)-1
	for i<j{
		if s[i] != s[j]{
			return false
		}
		i++
		j--
	}
	return true
}

func main() {
	fmt.Println(partition("aab"))
}