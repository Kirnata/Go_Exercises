package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func reverseStr(str string) string {
	strs := strings.Split(str, " ")
	var builder strings.Builder
	for i := len(strs) - 1; i > 0; i-- {
		builder.WriteString(strs[i])
		builder.WriteRune(' ')
	}
	builder.WriteString(strs[0])
	return builder.String()
}

// Разработать программу, которая переворачивает слова в строке.
// Пример: «snow dog sun — sun dog snow»
func main() {
	fmt.Println("Start string:")
	in := bufio.NewScanner(os.Stdin)
	in.Scan()
	str := in.Text()
	fmt.Println(str)
	fmt.Println("After reverse string:")
	fmt.Println(reverseStr(str))

}
