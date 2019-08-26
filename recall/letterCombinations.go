package main

import "fmt"

/**
17
给定一个仅包含数字 2-9 的字符串，返回所有它能表示的字母组合。

给出数字到字母的映射如下（与电话按键相同）。注意 1 不对应任何字母。

输入："23"
输出：["ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"].
 */
var kv = map[int]string {
	2: "abc",
	3: "def",
	4: "ghi",
	5: "jkl",
	6: "mno",
	7: "pqrs",
	8: "tuv",
	9: "wxyz",
}

func letterCombinations(digits string) []string {
	var res []string
	if len(digits) == 0 {
		return res
	}
	letterCombinationsCore(digits, 0, &res, "")
	return res
}

func letterCombinationsCore(digits string, index int, pRes *[]string, str string) {
	if index == len(digits) -1 {
		for _,v := range kv[int(digits[index] -'0')] {
			*pRes = append(*pRes, str+string(v))
		}
	}else {
		for _,v := range kv[int(digits[index]-'0')] {
			letterCombinationsCore(digits,index+1,pRes,str+string(v))
		}
	}


}

func main() {

	fmt.Println(int("23"[0]-'0'))
	fmt.Println(letterCombinations("23"))

}