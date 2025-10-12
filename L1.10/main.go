package main

import (
	"fmt"
	"math"
)

func round(key float64) int {
	var k int = 0
	flag := 1
	if key < 0 {
		flag = -1
	}
	d := int(math.Abs(key))
	for d > 10 {
		k++
		d = d / 10
	}
	return flag * d * int(math.Pow10(k))
}

func main() {
	temp := [...]float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	map_data := make(map[int][]float64)

	for _, val := range temp {
		key := round(val)
		fmt.Println(key)
		if value, ok := map_data[key]; ok {
			fmt.Println(ok, "value", value)
			map_data[key] = append(value, val)
		} else {
			map_data[key] = []float64{val}
		}
	}
	fmt.Println(map_data)
}
