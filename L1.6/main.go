package main

import (
	"context"
	"fmt"
	"runtime"
	"sync"
	"time"
)

func flag_(flag *bool, wg *sync.WaitGroup) {
	wg.Add(1)
	defer wg.Done()
	for !*flag {
		fmt.Println("Горутина с условием работает")
		time.Sleep(time.Second)
	}
	fmt.Println("Завершение работы горутины по условию (flag)")

}
func context_(ctx context.Context, wg *sync.WaitGroup, str string) {
	wg.Add(1)
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("Завершение работы горутины через контекст (%v)\n", str)
			return
		default:
			fmt.Printf("Горутина с контекстом (%v) работает\n", str)
			time.Sleep(time.Second)

		}

	}

}

func chan_(c chan int, wg *sync.WaitGroup) {
	wg.Add(1)
	defer wg.Done()
	for {
		select {
		case <-c:
			fmt.Println("Завершение работы по каналу")
			return
		default:
			fmt.Println("Горутина с каналом работает")
			time.Sleep(time.Second)
		}
	}
}

func runtime_(wg *sync.WaitGroup) {
	wg.Add(1)
	wg.Done()
	i := 0
	for {
		fmt.Println("Горутина с GoExit работает, value =", i)
		if i == 5 {
			fmt.Println("Завершение работы по GoExit")
			runtime.Goexit()
		}
		time.Sleep(time.Second)
		i++
	}
}

func main() {
	var wg sync.WaitGroup
	ch := make(chan int)
	ctx1, cancel1 := context.WithCancel(context.Background())
	ctx2, cancel2 := context.WithTimeout(context.Background(), 2*time.Second)
	flag := false
	go flag_(&flag, &wg)
	go context_(ctx1, &wg, "WithCancel")
	go context_(ctx2, &wg, "WithTimeout")
	go chan_(ch, &wg)
	go runtime_(&wg)
	time.Sleep(3 * time.Second)
	fmt.Println("flag = true")
	flag = true
	fmt.Println("cancel1()")
	time.Sleep(time.Second)
	cancel1()
	fmt.Println("cancel2()")
	time.Sleep(2 * time.Second)
	cancel2()
	fmt.Println("chan <-")
	ch <- 1
	time.Sleep(2 * time.Second)
	wg.Wait()
	fmt.Println("main() stoped")
}
