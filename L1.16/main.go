package main

import "fmt"

func quickSort(numbers []int, left int, right int) {
	lhold := left
	rhold := right
	pivot := numbers[left]
	for left < right {
		for numbers[right] > pivot && left < right {
			right--
		}
		if left != right {
			numbers[left] = numbers[right]
			left++
		}
		for (numbers[left] < pivot) && (left < right) {
			left++
		}
		if left != right {
			numbers[right] = numbers[left]
			right--
		}
	}
	numbers[left] = pivot // ставим разрешающий элемент на место
	index := left
	left = lhold
	right = rhold
	if left < index {
		quickSort(numbers, left, index-1)
	}
	if right > index {
		quickSort(numbers, index+1, right)
	}
}

func main() {
	data := []int{10, 4, 2, 14, 67, 2, 11, 33, 1, 15}
	quickSort(data, 0, len(data)-1)
	fmt.Println(data)
}
