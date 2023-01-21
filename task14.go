package main

import "fmt"

func typeCheck(i interface{}) {
	switch i.(type) {
	case int:
		fmt.Println("It's integer!")
	case string:
		fmt.Println("It's string!")
	case bool:
		fmt.Println("It's bool!")
	case chan int:
		fmt.Println("It's channel int!")
	default:
		fmt.Println("Unknown type")
	}
}

// Разработать программу, которая в рантайме способна определить тип переменной:
// int, string, bool, channel из переменной типа interface{}
func main() {
	ch := make(chan int)
	arr := make([]int, 5)
	typeCheck(2)
	typeCheck("aboba")
	typeCheck(true)
	typeCheck(ch)
	typeCheck(arr)
}
