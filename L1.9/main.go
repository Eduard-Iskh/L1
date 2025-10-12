package main

import (
	"fmt"
	"sync"
	"time"
)

func read(ch_r <-chan int, ch_w chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for val := range ch_r {
		fmt.Printf("\nДобавляем x*2 в канал (%d)\n", val)
		ch_w <- (val)
	}

	close(ch_w)
}

func write(ch_r chan int, array []int, wg *sync.WaitGroup) {
	defer wg.Done()
	for j, val := range array {
		fmt.Printf("\nЗапись %d-ого val = %d в канал\n", j, val)
		ch_r <- val * 2
		time.Sleep(500 * time.Millisecond)
	}
	fmt.Println("Закрываем канал записи")
	close(ch_r)
}

func main() {
	var wg sync.WaitGroup
	values := [...]int{1, 2, 3, 4}
	ch_w := make(chan int)
	ch_r := make(chan int)
	wg.Add(2)
	go write(ch_r, values[:], &wg)
	go read(ch_r, ch_w, &wg)
	go func() {
		for result := range ch_w {
			fmt.Printf("\nРезультат: %d\n", result)
		}
	}()
	fmt.Println("Ждем завершения чтения")
	wg.Wait()
	fmt.Println("Завершение main")
}
