package main

import (
	"fmt"
	"math/rand"
	"time"
)

func readData(ch chan int) {
	count := 0

	for x := range ch {
		fmt.Println(x)
		count++
	}
}

func writeData(ch chan int) {
	for {
		ch <- rand.Int()
	}
}

func timer(N time.Duration, to chan bool) {
	time.AfterFunc(N*time.Second, func() {
		fmt.Println("Time is over!")
		to <- true
	})
}

// Разработать программу, которая будет последовательно отправлять значения в канал,
// а с другой стороны
// канала — читать. По истечению N секунд программа должна завершаться

func main() {
	var N time.Duration
	_, err := fmt.Scan(&N)
	if err != nil {
		return
	}
	to := make(chan bool)
	ch := make(chan int)

	go timer(N, to)
	go writeData(ch)
	go readData(ch)
	c := <-to
	if c {
		close(ch)
	}
}
