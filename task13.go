package main

import "fmt"

// Поменять местами два числа без создания временной переменной
func main() {
	k := 2
	i := 4
	fmt.Printf("k == %d, i := %d\n", k, i)
	k, i = i, k
	fmt.Println("Swapping...")
	fmt.Printf("k == %d, i := %d\n", k, i)
}
