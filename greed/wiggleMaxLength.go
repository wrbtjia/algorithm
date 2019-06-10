package main


func wiggleMaxLength(nums []int) int {
	if len(nums) < 2 {
		return len(nums)
	}
	begin := 0
	up := 1
	down := 2

	state := begin
	lenght := 1

	for i:=1;i<len(nums);i++ {
		switch state {
		case begin:
			if nums[i-1] < nums[i] {
				state = up
				lenght++
			}else if nums[i-1] > nums[i] {
				state = down
				lenght++
			}
			break
		case up:
			if nums[i-1] > nums[i] {
				state = down
				lenght++
			}
			break
		case down:
			if nums[i-1] < nums[i] {
				state = up
				lenght++
			}
			break
		}
	}
	return lenght

}

func main() {

}