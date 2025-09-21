package main

import (
	"fmt"
	"math"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	array := [5]float64{2, 4, 6, 8, 10}
	for _, i := range array {
		wg.Add(1)
		go func(i float64) {
			defer wg.Done()
			fmt.Println(int(math.Pow(i, 2)))
		}(i)
	}
	wg.Wait()

}
