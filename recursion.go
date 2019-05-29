package main

import "fmt"

func recursion(n int) int {
	if n == 1 {
		return 1
	}
	return n * recursion(n-1)
}

func main() {

	//res := recursion(10)

	maps := make(map[int]int, 3)

	maps[1] = 1
	maps[2] = 2
	maps[3] = 3

	fmt.Println(maps)

	maps[1] = 2

	fmt.Println(maps)
	maps[1] = 5
	fmt.Println(maps)
}
