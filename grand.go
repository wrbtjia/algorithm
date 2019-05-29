package main

import "fmt"

func grand(n int) int {
	if n == 100 {
		return n
	}

	return grand(n+1) + n
}
func main() {
	s := grand(1)
	fmt.Println(s)
}
