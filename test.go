package main

import "fmt"

func main() {
		arr :=[]int{1,2,3,4}
		arr2 := arr[:2:2]
		arr2 = append(arr2, 5)
		fmt.Println(arr2)
}
