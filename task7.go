package main

import (
	"fmt"
	"math/rand"
	"sync"
)

func filler(myMap map[int]int, mu *sync.Mutex) {
	for i := 0; i < 30; i++ {
		key := rand.Intn(50)
		mu.Lock()
		if _, ok := myMap[key]; !ok {
			myMap[key] = i
		}
		mu.Unlock()
	}
}

func modOpposite(i int) int {
	if i > 0 {
		return -i
	}
	return i
}

func fillerToNegative(myMap map[int]int, mu *sync.Mutex) {
	for i := 0; i < 150; i++ {
		key := rand.Intn(50)
		mu.Lock()
		if v, ok := myMap[key]; ok {
			myMap[key] = modOpposite(v)
		}
		mu.Unlock()
	}
}

// Реализовать конкурентную запись данных в map
func main() {
	var mu sync.Mutex
	myMap := make(map[int]int)
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		filler(myMap, &mu)
	}()

	go func() {
		defer wg.Done()
		fillerToNegative(myMap, &mu)
	}()
	wg.Wait()
	fmt.Println(myMap)
}
