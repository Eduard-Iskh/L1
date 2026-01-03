package main

import "fmt"

func reverse(str []int32, left int, rigth int) []int32 {
	for i, j := rigth, left; i > j; i-- {
		str[i], str[j] = str[j], str[i]
		j++
	}
	return str
}

func main() {
	var txt string
	var ans []int32
	txt = "cлова разделяются одиночным пробелом"
	fmt.Println("Исходная строка:", txt)
	ans = []rune(txt)
	ans = reverse(ans, 0, len(ans)-1)

	left := 0
	rigth := 0
	for index, value := range ans {
		if value == ' ' {
			rigth = index - 1
			ans = reverse(ans, left, rigth)
			left = index + 1
		}
		if index == (len(ans) - 1) {
			rigth = index
			ans = reverse(ans, left, rigth)
		}
	}
	fmt.Println("Результат reverse:", string(ans))
}
