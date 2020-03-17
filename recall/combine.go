package main

import "fmt"

/**
77. 组合
给定两个整数 n 和 k，返回 1 ... n 中所有可能的 k 个数的组合。

示例:

输入: n = 4, k = 2
输出:
[
  [2,4],
  [3,4],
  [2,3],
  [1,2],
  [1,3],
  [1,4],
]

 */


func combine(n int, k int) [][]int {
	var result [][]int
	 temp:= make([]int,0)
	combinations(n,k,1,temp,&result)
	return result
}

func combinations(n int, k int,start int,temp []int,result *[][]int)  {
	if k == len(temp){
		t := make([]int, len(temp))
		copy(t, temp)
		*result = append(*result, t)
	}

	for i := start;i <= n ;i++{
		temp = append(temp, i)
		combinations(n,k,i+1,temp,result)
		temp = temp[:len(temp)-1]
	}

}

func main() {

	fmt.Println(combine(4,2))
}

