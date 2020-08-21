package main

import "container/heap"

/**
347. 前 K 个高频元素
给定一个非空的整数数组，返回其中出现频率前 k 高的元素。
示例 1:

输入: nums = [1,1,1,2,2,3], k = 2
输出: [1,2]
示例 2:

输入: nums = [1], k = 1
输出: [1]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/top-k-frequent-elements
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

type Feq struct {
	val int
	cnt int
}
type FeqMinHeap []Feq

func (pq *FeqMinHeap) Len() int {
	return len(*pq)
}
func (pq *FeqMinHeap) Less(i, j int) bool {
	return (*pq)[i].cnt < (*pq)[j].cnt
}
func (pq *FeqMinHeap) Swap(i, j int) {
	(*pq)[i], (*pq)[j] = (*pq)[j], (*pq)[i]
}

func (pq *FeqMinHeap) Push(x interface{}) {
	*pq = append(*pq, x.(Feq))
}

func (pq *FeqMinHeap) Pop() interface{} {
	n := len(*pq) - 1
	x := (*pq)[n]
	*pq = (*pq)[:n]
	return x
}
func (pq *FeqMinHeap) Peek() Feq {
	return (*pq)[0]
}

func topKFrequent(nums []int, k int) []int {
	cntMap := make(map[int]int)
	for _, num := range nums {
		cntMap[num]++
	}
	minHeap := &FeqMinHeap{}
	for key, val := range cntMap {
		heap.Push(minHeap, Feq{
			val: key,
			cnt: val,
		})
		if minHeap.Len() > k {
			heap.Pop(minHeap)
		}
	}
	res := make([]int, k)
	i := 0
	for minHeap.Len() > 0 {
		res[i] = minHeap.Pop().(Feq).val
		i++
	}
	return res
}
