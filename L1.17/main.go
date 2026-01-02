package main

import "fmt"

func search(array []int, number int) int {

	left := 0
	higth := len(array) - 1

	for left <= higth {
		index := (higth-left)/2 + left
		if array[index] == number {
			return index
		} else if array[index] > number {
			higth = index - 1
		} else {
			left = index + 1
		}
	}
	return -1

}

func main() {
	array := []int{0, 1, 2, 3, 7, 8, 9, 10, 12}
	ans := search(array, 12)
	fmt.Println(ans)
}
