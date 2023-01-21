package main

import (
	"fmt"
	"sync"
)

// Написать программу, которая конкурентно рассчитает значение квадратов чисел
// взятых из массива (2,4,6,8,10) и выведет их квадраты в stdout

func main() {
	var arr = [5]int{2, 4, 6, 8, 10}
	var wg sync.WaitGroup
	wg.Add(len(arr))

	for _, v := range arr {
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Println(v * v)
		}()
	}
	wg.Wait()
}
