package main

import (
	"fmt"
	"strconv"
)

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
	res:=removeKdigits("10",2)
	fmt.Println(res)
}
