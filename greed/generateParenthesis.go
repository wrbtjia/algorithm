package main

import "fmt"

/**
给出 n 代表生成括号的对数，请你写出一个函数，使其能够生成所有可能的并且有效的括号组合。
例如，给出 n = 3，生成结果为：
[
  "((()))",
  "(()())",
  "(())()",
  "()(())",
  "()()()"
]
 */
func generateParenthesis(n int) []string {
	var str =[]string{}
	gen("",n,n,&str)
	return str
}

func gen(item string,l,r int,str *[]string)  {
	if l ==0 && r ==0 {
		*str = append(*str, item)
		return
	}
	if l > 0 {
		gen(item+"(",l-1,r,str)
	}
	if l < r {
		gen(item+")",l,r-1,str)
	}
}

func main() {
	fmt.Println(generateParenthesis(3))
}
