package main

import (
	"fmt"
	"strconv"
)


/**
402. 移掉K位数字
给定一个以字符串表示的非负整数 num，移除这个数中的 k 位数字，使得剩下的数字最小。
注意:

num 的长度小于 10002 且 ≥ k。
num 不会包含任何前导零。
示例 1 :

输入: num = "1432219", k = 3
输出: "1219"
解释: 移除掉三个数字 4, 3, 和 2 形成一个新的最小的数字 1219。
示例 2 :

输入: num = "10200", k = 1
输出: "200"
解释: 移掉首位的 1 剩下的数字为 200. 注意输出不能有任何前导零。
示例 3 :

输入: num = "10", k = 2
输出: "0"
解释: 从原数字移除所有的数字，剩余为空就是0。

 */

func removeKdigits(num string, k int) string {

	var arr=[]int{}
	for i:=0;i< len(num);i++  {
		x := num[i] -'0'

		//1 4 3 2 2 1 9
		for len(arr) != 0 && arr[len(arr)-1] > int(x) && k>0{
			arr = arr[:len(arr)-1]
			k--
		}
		if int(x) != 0 || len(arr) != 0 {
			arr = append(arr, int(x))
		}

	}

	for len(arr) != 0 && k > 0 {
		arr = arr[:len(arr)-1]
		k--
	}
	var res string =""
	for _,c:=range arr{
		res=res+strconv.Itoa(c)
	}
	if res =="" {
		res = "0"
	}
	return res
}

func main() {

	var nums =[]int{}
	nums = append(nums, 1)
	nums = append(nums, 2)
	nums = append(nums, 3)
	nums = append(nums, 4)
	nums = append(nums, 5)


	res:=removeKdigits("1432219",3)
	fmt.Println(res)
}

