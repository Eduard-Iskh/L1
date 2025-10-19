package main

import "fmt"

func main() {
	a, b := 5, 10
	fmt.Printf("После замены: a = %d, b = %d", a, b)
	a = a + b
	b = a - b
	a = a - b
	fmt.Printf("\nПосле замены: a = %d, b = %d", a, b)
}
