package main

import (
	"fmt"
	"sync"
	"time"
)

func add(score *int, mu *sync.Mutex, wg *sync.WaitGroup, i int) {
	defer wg.Done()
	fmt.Println("Начало работы горутины", i)
	time.Sleep(2 * time.Second)
	mu.Lock()
	*score += 1
	mu.Unlock()
	time.Sleep(time.Second)
	fmt.Println("Завершнеие горутины", i)

}

func main() {
	var mu sync.Mutex
	var wg sync.WaitGroup
	score := 0
	for i := range 4 {
		wg.Add(1)
		go add(&score, &mu, &wg, i)

	}
	wg.Wait()
	fmt.Println(score)
}
