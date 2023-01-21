package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
)

// Реализовать постоянную запись данных в канал (главный поток).
// Реализовать набор из N воркеров, которые читают произвольные данные из канала и выводят в stdout.
// Необходима возможность выбора количества воркеров при старте.
// Программа должна завершаться по нажатию Ctrl+C. Выбрать и обосновать способ завершения работы всех воркеров

func worker(chlInt chan int) {
	for x := range chlInt {
		fmt.Println(x)
	}
}

func startWorkers(countWorkers int, chlInt chan int) {
	for countWorkers > 0 {
		go worker(chlInt)
		countWorkers--
	}
}

func checkSignal(c <-chan os.Signal, done chan bool) {
	if <-c == os.Interrupt {
		done <- true
	}
	return
}

func main() {
	var countWorkers int
	_, err := fmt.Scan(&countWorkers)
	if err != nil {
		return
	}
	chInt := make(chan int)
	c := make(chan os.Signal)
	done := make(chan bool)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	go checkSignal(c, done)
	startWorkers(countWorkers, chInt)
	loop := true
	for loop {
		select {
		case <-done:
			fmt.Println("channel done")
			close(chInt)
			loop = false
		default:
			chInt <- rand.Intn(100)
		}
	}
	fmt.Println("ending...")
}
