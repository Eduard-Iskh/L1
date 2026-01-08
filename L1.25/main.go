package main

import (
	"fmt"
	"time"
)

func sleep(duration time.Duration) {
	<-time.After(duration)

}

func sleep2(duretion time.Duration) {
	timer := time.NewTimer(duretion)
	<-timer.C
}

func main() {
	var t int
	fmt.Println("Введите время")
	fmt.Scan(&t)
	sleep2(time.Duration(t) * time.Second)
	fmt.Println("Работа завершна")
}
