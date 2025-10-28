package main

import "fmt"

func check(v interface{}) {
	switch v.(type) {
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	case bool:
		fmt.Println("bool")
	case chan int:
		fmt.Println("chan int")
	case chan string:
		fmt.Println("chan string")
	}
}

func main() {
	i := 42
	s := "hello"
	b := true
	chInt := make(chan int)
	chStr := make(chan string)
	check(i)
	check(s)
	check(b)
	check(chInt)
	check(chStr)
}
