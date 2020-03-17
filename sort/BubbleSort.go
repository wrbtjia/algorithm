package main

import "fmt"

//冒泡排序
func main() {
	array :=[]int{11,9,2,8,3,7,4,6,5,10}

	for i:=0;i<len(array)-1;i++  {
		flag := false
		for j:=0;j<len(array)-1-i;j++ {
			if array[j]>array[j+1]{
				array[j],array[j+1]= array[j+1],array[j]
				flag = true
			}
		}
		if !flag {
			break
		}
	}
	fmt.Println(array)
}
