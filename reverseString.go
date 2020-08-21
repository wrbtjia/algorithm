package main

import "fmt"

/**
反转字符串
 */

func main() {
	str := "i am mans"
	s := []byte(str)
	reverseString(s)
	l :=0
	for i:=0;i<len(s);i++{
		if string(s[i]) == " " {
			reverseString(s[l:i])
			l=i+1
		}

		if i+1 == len(s) {
			reverseString(s[l:i+1])
		}
	}

	fmt.Println(string(s))
}

func reverseString(by []byte)  {

	r :=len(by)
	for i:=0;i<len(by)/2;i++ {
		by[i],by[r-1-i] = by[r-1-i],by[i]
	}
	fmt.Println(string(by))
}