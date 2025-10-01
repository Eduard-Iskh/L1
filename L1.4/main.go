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

func Write(ctx context.Context, worker *int, wg *sync.WaitGroup, ch chan string) {
	for j := 0; j < *worker; j++ {
		wg.Add(1)
		go func(c <-chan string, wg *sync.WaitGroup) {
			defer wg.Done()
			for val := range c {
				select {
				case <-ctx.Done():
					fmt.Println("Завершение работы worker", j,
						" из-за нажатия клавишь ctrl + c")
					return
				default:
					fmt.Printf("Воркер %d, val = %v\n", j, val)
					time.Sleep(time.Second)
				}
			}
		}(ch, wg)
	}
}

func main() {
	var wg sync.WaitGroup

	roc := make(chan string)

	ctx, cancel := context.WithCancel(context.Background())

	defer cancel()

	worker := flag.Int("worker", 3, "количество воркеров")

	flag.Parse()

	Write(ctx, worker, &wg, roc)

	var data string

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT)

	go func() {
		<-sigChan
		fmt.Println("\nПолучен сигнал SIGINT (Ctrl+C), завершаем работу...")
		cancel() // Отменяем контекст - уведомляем все воркеры о завершении

		// Даем воркерам время на корректное завершение
		time.Sleep(100 * time.Millisecond)
		close(roc) // Закрываем канал данных
	}()

inputLoop:
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Прерван ввод данных")
			break inputLoop
		default:
			fmt.Println("Введите данные (или Ctrl+C для выхода):")
			fmt.Scan(&data)
			select {
			case roc <- data:
				// Данные успешно отправлены
			case <-ctx.Done():
				break inputLoop
			}
		}

	}

	close(roc)
	wg.Wait()
	fmt.Println("main() stoped")
}
