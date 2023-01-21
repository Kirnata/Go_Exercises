package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func quitWithChannel(ch chan bool) {
	select {
	case <-ch:
		fmt.Println("stop quitWithChannel")
	}
}

func quitWithClosedChan(ch chan int) {
	fmt.Println("start quitWithClosedChan")
	for {
		i, open := <-ch
		if !open {
			break
		}
		fmt.Println(i)
	}
	fmt.Println("stop quitWithClosedChan")
}

func cleanAll() {
	fmt.Println("Clean all")
}

func sleepRandom(ch chan bool) {
	randomNumber := rand.Intn(3000)
	fmt.Printf("randomNumber == %d\n", randomNumber)
	time.Sleep(time.Duration(randomNumber) * time.Millisecond)
	if ch != nil {
		ch <- true
	}
}

func quitWithContext(ctx context.Context) {
	fmt.Println("start quitWithContext")
	quit := make(chan bool)
	go sleepRandom(quit)
	select {
	case <-ctx.Done():
		cleanAll()
		fmt.Println("stop quitWithContext by cancel")
	case <-quit:
		fmt.Println("quitWithContext returned")
	}
}

func quitWithTimer(ctx context.Context, wg *sync.WaitGroup) {
	select {
	case <-ctx.Done():
		fmt.Println("stop by canceling")
		wg.Done()
	}
}

// Реализовать все возможные способы остановки выполнения горутины
func main() {
	// 1
	ch1 := make(chan bool)
	go quitWithChannel(ch1)
	time.Sleep(5 * time.Second)
	ch1 <- false
	close(ch1)
	// 2
	ch2 := make(chan int)
	go quitWithClosedChan(ch2)
	for i := 0; i < 5; i++ {
		ch2 <- i
	}
	close(ch2)
	// 3
	ctx, cancel := context.WithCancel(context.Background())
	defer func() {
		fmt.Println("main defer: canceling context")
		cancel()
	}()
	go func() {
		sleepRandom(nil)
		cancel()
	}()
	quitWithContext(ctx)
	// 4
	ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	var wg sync.WaitGroup
	wg.Add(1)
	go quitWithTimer(ctx, &wg)
	wg.Wait()
}
