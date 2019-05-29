package main

import "fmt"

func twoFind(array []int, val int) {
	low := 0
	hig := len(array) - 1

	for low <= hig {
		mid := low + (hig-low)/2

		midVal := array[mid]

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
	var er = []int{1, 2, 5, 7, 15, 25, 30, 36, 39, 51, 67, 78, 80, 82, 85, 91, 92, 97}

	twoFind(er, 97)

}
