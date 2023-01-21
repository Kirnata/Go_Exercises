package main

import (
	"fmt"
	"math/big"
)

// Разработать программу, которая перемножает, делит, складывает, вычитает
// две числовых переменных a,b, значение которых > 2^20.
func main() {
	var str1 string
	var str2 string
	var operation string
	fmt.Println("enter the first number:")
	fmt.Scan(&str1)
	fmt.Println("enter the second number:")
	fmt.Scan(&str2)
	fmt.Println("enter the operation:")
	fmt.Scan(&operation)

	x := new(big.Int)
	x.SetString(str1, 10)
	y := new(big.Int)
	y.SetString(str2, 10)
	res := new(big.Int)

	switch operation {
	case "*":
		res.Mul(x, y)
	case "/":
		nul := big.NewInt(0)
		if y.Cmp(nul) == 0 {
			fmt.Println("division by zero")
			return
		}
		res.Div(x, y)
	case "+":
		res.Add(x, y)
	case "-":
		res.Sub(x, y)
	default:
		fmt.Println("Invalid operation: try *, /, + or -")
		return
	}

	fmt.Println(res)
}
