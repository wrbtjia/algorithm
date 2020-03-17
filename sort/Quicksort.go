package main

import "fmt"

/**

算法描述
步骤为：

从数列中挑出一个元素，称为”基准”（pivot），
重新排序数列，所有比基准值小的元素摆放在基准前面，所有比基准值大的元素摆在基准后面（相同的数可以到任何一边）。
在这个分区结束之后，该基准就处于数列的中间位置。这个称为分区（partition）操作。
递归地（recursively）把小于基准值元素的子数列和大于基准值元素的子数列排序。
递归到最底部时，数列的大小是零或一，也就是已经排序好了。这个算法一定会结束，因为在每次的迭代（iteration）中，它至少会把一个元素摆到它最后的位置去。

一般采用的都是原地排序版本，这样不需要分配额外的空间，对速度有提升。

 */

 /**
 实现
单路快排
递归版本实现

首先通过随机数选取比较的点
遍历 head + 1-> tail
大于mid 则和tail交换 并且tail–
小于等于 则交换head 并且head++ i++
疑问：

为什么在大于的情况下 index i 没有移位？
因为交换后的tail，并没有保证一定会大于mid元素，所以需要再次进行比较

为什么在小于的情况下 index i 需要移位？
因为第一次交换的时候是arr[0] -> mid，所以交换后一定满足arr[i] <= mid，所以可以移位操作。否则遍历无法结束

  */


func quickSort(arr []int, start, end int) {
	if start < end {
		i, j := start, end
		key := arr[(start+end)/2]
		for i <= j {
			for arr[i] < key {
				i++
			}
			for arr[j] > key {
				j--
			}
			if i <= j {
				arr[i], arr[j] = arr[j], arr[i]
				i++
				j--
			}
		}
		if start < j {
			quickSort(arr, start, j)
		}
		if end > i {
			quickSort(arr, i, end)
		}
	}
}

func main() {
	arr := []int{3, 7, 9, 8, 38, 93, 12, 222, 45, 93, 23, 84, 65, 2}
	qqqq(arr, 0, len(arr)-1)
	fmt.Println(arr)
}

func qqqq(arr []int,s,e int)  {
	if s < e{
		x,y := s,e
		mid := arr[(s+e)/2]
		fmt.Println(mid)
		for x<=y {
			for arr[x] < mid {
				x++
			}
			for arr[y] > mid {
				y--
			}
			if x <= y {
				arr[x],arr[y] = arr[y],arr[x]
				x++
				y--
			}
		}
		qqqq(arr,s,y)
		qqqq(arr,x,e)

	}
}

func kkk(arr []int,k,j int)  {
	if k<j {
		i,p:=k,j
		mid:=arr[(k+j)/2]
		for i<=p {
			for arr[i]<mid {
				i++
			}
			for arr[p]>mid {
				p--
			}
			if i<=p {
				arr[i],arr[p]=arr[p],arr[i]
				i++
				p--
			}
		}
		kkk(arr,k,p)
		kkk(arr,i,j)


	}
}