package main

import "fmt"

/**
93. 复原IP地址

给定一个只包含数字的字符串，复原它并返回所有可能的 IP 地址格式。

示例:

输入: "25525511135"
输出: ["255.255.11.135", "255.255.111.35"]

 */
func restoreIpAddresses(s string) []string {
	if len(s) > 12 {
		return []string{}
	}
	ret := &[]string{}
	ip("", s, 3, ret)
	return *ret
}

func ip(path, s string, remain int , ret *[]string)  {

	if remain == -1 && len(s) == 0 {
		*ret = append(*ret, path[:len(path)-1])
	}
	//fmt.Println(path, left , remain , ret)
	val := rune(0)
	sval := ""

	for i:=0;i< 3 && i<len(s);i++{
		val = 10 * val + rune(s[i]) - '0'
		sval += string(s[i])

		if val<256 && val >=0 {
			ip(path +  sval + ".", s[i+1:], remain-1, ret)
		}
		if val ==0 {// 0.00. not allowed
			break
		}
	}
}

func main() {
	fmt.Println(restoreIpAddresses("25525511135"))
}