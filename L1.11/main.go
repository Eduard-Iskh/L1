package main

import (
	"fmt"
)

func main() {
	a := []int{1, 2, 3, 4, 4}
	b := []int{3, 4, 5}
	ans := make(map[int](bool))
	c := make([]int, 0, 2)
	for _, i := range a {
		ans[i] = true
	}
	for _, j := range b {
		if ans[j] {
			ans[j] = false
			c = append(c, j)
		}
	}
	fmt.Println(c)
}
