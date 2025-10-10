package main

import (
	"fmt"
	"sync"
	"time"
)

type Data struct {
	mut  sync.RWMutex
	data map[string]int
}

func (d *Data) write(key string, value int) {
	d.mut.Lock()
	defer d.mut.Unlock()
	fmt.Printf("Запись данных data[%v] = %d в data\n", key, value)
	time.Sleep(time.Second)
	d.data[key] = value
}

func main() {
	fmt.Printf("\nНачало записи\n\n")
	time.Sleep(500 * time.Millisecond)
	var wg sync.WaitGroup
	data := Data{
		data: make(map[string]int),
	}
	wg.Add(3)
	for i := range 10 {
		wg.Add(1)
		go func(j int) {
			defer wg.Done()
			data.write("key", j)
		}(i)
	}
	go func() {
		defer wg.Done()
		data.write("ex_2", 2)
	}()
	go func() {
		defer wg.Done()
		data.write("ex_3", 3)
	}()
	go func() {
		defer wg.Done()
		data.write("ex_4", 10)
	}()

	wg.Wait()
	fmt.Println(data.data)
	fmt.Println("\nВсе горутины закрыты")
	fmt.Println("Завершение работы main")
}
