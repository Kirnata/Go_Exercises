package main

import (
	"fmt"
	"time"
)

func ftAdd(counter *int, id int) {
	for {
		*counter++
		fmt.Printf("counter in %d ftAdd: %d\n", id, *counter)
	}

}

func timer(N time.Duration, to chan bool) {
	time.AfterFunc(N*time.Second, func() {
		fmt.Println("Time is over!")
		to <- true
	})
}

func main() {
	counter := 0
	timeCh := make(chan bool)
	go timer(5, timeCh)
	go ftAdd(&counter, 1)
	go ftAdd(&counter, 2)
	<-timeCh
	fmt.Printf("Couner == %d\n", counter)
}
