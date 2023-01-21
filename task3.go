package main

import "fmt"

// Дана последовательность чисел: 2,4,6,8,10. Найти сумму их квадратов(2^2+4^2+….)
// с использованием конкурентных вычислений
func main() {
	var digits = []int{2, 4, 6, 8, 10}
	ch := make(chan int)

	go func() {
		for _, v := range digits {
			ch <- v * v
		}
		close(ch)
	}()
	sum := 0
	for x := range ch {
		sum += x
	}
	fmt.Println(sum)
}
