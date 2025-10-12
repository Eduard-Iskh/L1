package main

import "fmt"

func SetBit(bit int, number int64, i int) int64 {
	if bit == 1 {
		return number | (1 << (i - 1))
	} else {
		return number &^ (1 << (i - 1))
	}

}

func main() {
	var value int64 = 8
	fmt.Printf("value: %d = 0x%b\n", value, value)
	var result int64
	result = SetBit(1, value, 3)
	fmt.Printf("Установка 1, 3-й бит: result = %d = 0x%b\n", result, result)
	result = SetBit(0, value, 4)
	fmt.Printf("Установка 0, 4-й бит: result = %d = 0x%b", result, result)
}
