package main

import (
	"fmt"
	"reflect"
)

func isAnagram(s string, t string) bool {
	m1 := make(map[string]int)
	m2 := make(map[string]int)

	for _, v := range s {
		m1[string(v)] += 1
	}
	for _, v := range t {
		m2[string(v)] += 1
	}

	return reflect.DeepEqual(m1, m2)

}

func main() {
	res := isAnagram("rat", "tra")

	fmt.Println(res)
}
