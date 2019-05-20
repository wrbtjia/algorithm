package main

import "fmt"

func main() {
	var datas =[10]int{3,7,1,4,0,8,2,6,5,9}

	for i:=1;i<len(datas);i++ {
		for j:=i; j>0 && datas[j] < datas[j-1];j-- {
				datas[j],datas[j-1] = datas[j-1],datas[j]

		}

	}

	fmt.Println(datas)

}
