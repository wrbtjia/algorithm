package main

import "fmt"

func main() {
	hanoi(3,"A","B","C")
}

func hanoi(n int,a,b,c string)  {
	if n <1 {
		return
	}
	if n == 1 {
		fmt.Printf("%s->%s\n",a,c)
	}else {
		hanoi(n-1,a,c,b)
		fmt.Printf("%s->%s\n",a,c)
		hanoi(n-1,b,a,c)
	}

}