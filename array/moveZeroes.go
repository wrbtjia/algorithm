package main

import "fmt"

/**
283
给定一个数组 nums，编写一个函数将所有 0 移动到数组的末尾，同时保持非零元素的相对顺序。
输入: [0,1,0,3,12]
输出: [1,3,12,0,0]
 */
func moveZeroes(nums []int)  {

	k:=0
	for i:=0;i< len(nums);i++  {

		if nums[i] > 0 {
			if i != k {
				nums[k],nums[i] = nums[i],nums[k]
				k++
			}else {
				k++
			}
		}

	}

}

func moveZeroes2(nums []int)  {
	z:=0

	for i:=0;i< len(nums);i++  {
		if nums[i] > 0{
			if i != z {
				nums[i],nums[z] = nums[z],nums[i]
				z++
			}else {
					z++
			}
		}
	}

}

func main() {
	var nums =[]int{2,1,0,3,0,9}
	moveZeroes2(nums)
	fmt.Println(nums)


}
