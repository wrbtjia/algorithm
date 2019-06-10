package main


func findDuplicates(nums []int) []int {
	maps := make(map[int]int)
	res :=[]int{}
	for i:= 0;i<len(nums);i++{
		if _,ok:=maps[nums[i]];ok{
			res = append(res,nums[i])
		}
		maps[nums[i]]=nums[i]
	}
	return res
}

func main() {
	in :=[]int{1,1}
	findDuplicates(in)
}