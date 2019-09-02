package main

import (
	"fmt"
	"sort"
)

//56. 合并区间
/*给出一个区间的集合，请合并所有重叠的区间。

示例 1:

输入: [[1,3],[2,6],[8,10],[15,18]]
输出: [[1,6],[8,10],[15,18]]
解释: 区间 [1,3] 和 [2,6] 重叠, 将它们合并为 [1,6].*/


func merge(intervals [][]int) [][]int {
	ret := make([][]int,0)

	if len(intervals) == 0 || len(intervals) == 1{
		return intervals
	}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})


	fmt.Println(intervals)

	for i := 0; i < len(intervals);{
		start := intervals[i][0]
		end := intervals[i][1]

		for {
			if i < len(intervals) && intervals[i][0] <= end{
				if intervals[i][1] > end {
					end = intervals[i][1]
				}
				i++
				continue
			}
			tmp := make([]int,2)
			tmp[0] = start
			tmp[1] = end
			ret = append(ret, tmp)
			break
		}
	}
	return ret

}

func main() {
	//[[1,3],[2,6],[8,10],[15,18]]
  in:=[][]int{{1,3},{2,6},{8,10},{15,18}}

fmt.Println(merge(in))
}