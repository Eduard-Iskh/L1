package main

import (
	"fmt"
)

func main() {
	a := []string{"cat", "cat", "dog", "cat", "tree"}
	ans := make(map[string](bool))
	c := make([]string, 0, 2)
	for _, i := range a {
		ans[i] = true
	}
	for key := range ans {
		c = append(c, key)
	}
	fmt.Println(c)
}
