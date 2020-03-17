package main

import "fmt"

/**
75
给定一个包含红色、白色和蓝色，一共 n 个元素的数组，原地对它们进行排序，使得相同颜色的元素相邻，并按照红色、白色、蓝色顺序排列。

此题中，我们使用整数 0、 1 和 2 分别表示红色、白色和蓝色。

输入: [2,0,2,1,1,0]
输出: [0,0,1,1,2,2]
 */

func sortColors(nums []int)  {

	zero :=-1
	two := len(nums)
	for i:=0;i< two;  {
		if nums[i] == 1 {
			i++
		}else if nums[i] == 2 {
			two--
			nums[i],nums[two] = nums[two],nums[i]
		}else {
			zero ++
			nums[i],nums[zero]=nums[zero],nums[i]
			i++
		}
	}
}



func main() {
	var nums = []int{2,0,2,1,1,0}
	sortColors2(nums)
	fmt.Println(nums)

}



/*输入: [2,0,2,1,1,0]
输出: [0,0,1,1,2,2]*/

func sortColors2(nums []int)  {
	r :=-1
	b := len(nums)
	for i:=0;i< b;  {
		if nums[i] == 1 {
			i++
		}else if nums[i] == 2{
				b--
				nums[i],nums[b] = nums[b],nums[i]
		}else {
			r++
			nums[i],nums[r] = nums[r],nums[i]
			i++
		}
	}
}