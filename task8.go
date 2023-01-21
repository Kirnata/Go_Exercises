package main

import (
	"fmt"
	"log"
)

func bitDoer(num, i, zeroOne int64) (int64, error) {
	if i < 0 || i > 63 { // 64-й бит знака
		err := fmt.Errorf("Invalid value of i: %d\n", i)
		return num, err
	}
	if zeroOne == 1 {
		num |= 1 << i
	} else {
		num &^= 1 << i
	}
	return num, nil
}

// Дана переменная int64.
// Разработать программу которая устанавливает i-й бит в 1 или 0
func main() {
	var digit, i, bitFl int64
	_, err := fmt.Scan(&digit, &i, &bitFl)
	if err != nil {
		log.Printf("%s", err)
		return
	}

	newDigit, err := bitDoer(digit, i, bitFl)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(newDigit)
	}
}
