package main

import "fmt"

func main() {
	var str string
	fmt.Println("Введите строку")
	fmt.Scanf("%v", &str)
	txt := []rune(str)
	for i := len(txt) - 1; i >= 0; i-- {
		fmt.Printf(string(txt[i]))
	}
	fmt.Println()
}
