package main

import "sort"

/**
贪心算法 分糖分果
假设你是一位很棒的家长，想要给你的孩子们一些小饼干。但是，每个孩子最多只能给一块饼干。
对每个孩子 i ，都有一个胃口值 gi ，这是能让孩子们满足胃口的饼干的最小尺寸；并且每块饼干 j ，都有一个尺寸 sj 。
如果 sj >= gi ，我们可以将这个饼干 j 分配给孩子 i ，这个孩子会得到满足。你的目标是尽可能满足越多数量的孩子，并输出这个最大数值。

输入: [1,2,3], [1,1]

输出: 1
 */
func findContentChildren(g []int, s []int) int {

	sort.Ints(g)
	sort.Ints(s)
	child := 0
	cook := 0
	for child <len(g) && cook < len(s)  {
		if g[child] <= s[cook] {
			child ++
		}
		cook ++
	}
	return child
}

/*func sort(datas []int)  {
	for i:=1;i<len(datas);i++ {
		for j:=i; j>0 && datas[j] < datas[j-1];j-- {
			datas[j],datas[j-1] = datas[j-1],datas[j]

		}

	}

}*/
func main()  {
	
}



