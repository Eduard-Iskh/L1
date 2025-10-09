package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"

	//"sync"
	"time"
)

func Write(ctx context.Context, worker *int, wg *sync.WaitGroup, ch chan int) {
	for j := 0; j < *worker; j++ {
		wg.Add(1)
		go func(c <-chan int, wg *sync.WaitGroup) {
			defer wg.Done()
			for {
				select {
				case <-ctx.Done():
					time.Sleep(500 * time.Millisecond)
					fmt.Printf("Завершение работы worker %d\n", j)
					return
				case val, ok := <-c:
					if !ok {
						fmt.Printf("Воркер %d: завершение работы\n", j)
						return
					}
					fmt.Printf("Воркер %d, val = %v\n", j, val)
					time.Sleep(time.Second)
				}
			}
		}(ch, wg)
	}

}

func main() {
	var wg sync.WaitGroup
	n := 5

	fmt.Printf("\nВремя работы программы %d(с)\n", n)

	time.Sleep(500 * time.Millisecond)

	roc := make(chan int)

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(n)*time.Second)

	defer cancel()

	worker := flag.Int("worker", 3, "количество воркеров")

	flag.Parse()

	Write(ctx, worker, &wg, roc)

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT)

	go func() {
		<-sigChan
		fmt.Println("\nПрошло заданное время")
		cancel()
		time.Sleep(3 * time.Second)

	}()

inputLoop:
	for i := 0; ; i++ {
		select {
		case <-ctx.Done():
			time.Sleep(2 * time.Second)
			fmt.Printf("\nРабота программы завершена через %d(с)\n", n)
			break inputLoop

		default:
			fmt.Printf("\nЗапись данных в канал: val = %d\n", i)
			roc <- i
			time.Sleep(10 * time.Millisecond)
		}
	}
	close(roc)
	wg.Wait()
	fmt.Println("Все горутины - воркеры закрыты")
	fmt.Println("Завершение работы main")
}
