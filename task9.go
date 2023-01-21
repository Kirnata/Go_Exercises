package main

import (
	"fmt"
)

func stageOne(ch1 chan int, intSl []int) { // пишет в канал 1
	for _, v := range intSl {
		ch1 <- v
	}
	close(ch1)
	fmt.Println("ch1 was closed")
}

func stageTwo(ch1, сh2 chan int) { // читает из 1, пишет во 2
	for v := range ch1 {
		сh2 <- v * 2
	}
	close(сh2)
	fmt.Println("ch2 was closed")
}

// Разработать конвейер чисел.
// Даны два канала: в первый пишутся числа (x) из массива,
// во второй — результат операции x*2,
// после чего данные из второго канала должны выводиться в stdout
func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	intSl := []int{0, 1, 2, 3, 4, 5, 6}

	go stageOne(ch1, intSl)
	go stageTwo(ch1, ch2)

	for v := range ch2 {
		fmt.Println(v)
	}
	fmt.Println("main() stopped")
}
