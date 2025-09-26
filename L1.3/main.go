package main

import (
	"flag"
	"fmt"
	"sync"

	//"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	worker := flag.Int("worker", 3, "количество воркеров")
	flag.Parse()

	roc := make(chan string)
	for j := 0; j < *worker; j++ {
		wg.Add(1)
		go func(c <-chan string, wg *sync.WaitGroup) {
			for val := range c {
				fmt.Println(val, "воркер", j)
				time.Sleep(time.Second)
			}
			wg.Done()
		}(roc, &wg)
	}
	var data string
	for {
		fmt.Println("Введите данные")
		fmt.Scan(&data)
		if data == "end" {
			fmt.Println("Завершение цикла ввода")
			close(roc)
			break
		} else {
			roc <- data
		}

	}
	wg.Wait()
	fmt.Println("main() stoped")
}
