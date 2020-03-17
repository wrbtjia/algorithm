package main

import "fmt"

func twoFind(array []int, val int) {
	low := 0
	hig := len(array) - 1

	for low <= hig {
		mid := low + (hig-low)/2

		midVal := array[mid]

		fmt.Println(midVal)

		if midVal == val {
			fmt.Println(midVal, "下标为：", mid)
			return
		} else if midVal > val {
			hig = mid - 1
		} else if midVal < val {
			low = mid + 1
		}

	}
}

/**
二分查找 必须是有序的数组
*/
func main() {
	var er = []int{2,4,6,8, 10, 12, 14, 16, 18,20, 22}

	twoFind(er, 16)

}
