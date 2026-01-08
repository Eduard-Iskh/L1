package main

import (
	"fmt"
	"strings"
)

func main() {
	simv := make(map[rune]int)

	var str string
	fmt.Println("Введите строку")
	fmt.Scan(&str)
	//str = "abcdplokijuhHHHH"
	lower := strings.ToLower(str)
	for _, i := range lower {
		_, ok := simv[i]
		if !ok {
			simv[i] += 1
		} else {
			fmt.Println(false)
			return
		}
	}
	fmt.Println(true)
}
