package main

import "fmt"

func main() {
	var array []int

	array = []int{1, 2, 3, 4, 5, 6, 7}

	var i int
	fmt.Println("Введите индекс элемента, который необходимо удалить от 0 до", len(array))
	fmt.Scan(&i)
	if i > len(array)-1 {
		fmt.Println("Идекс выходит за пределы среза")
	} else {
		copy(array[i:], array[i+1:])
		array = array[:cap(array)-1]
	}
	fmt.Println("Результат выполнения программы:\n", array)
}
