package main

import "fmt"

func revStr(str string) string {
	runes := []rune(str)
	i := 0
	j := len(runes) - 1
	for i < j {
		runes[i], runes[j] = runes[j], runes[i]
		i++
		j--
	}
	return string(runes)
}

// Разработать программу, которая переворачивает подаваемую на ход строку (например: «главрыба — абырвалг»).
// Символы могут быть unicode
func main() {
	var str string
	fmt.Println("Start string:")
	_, err := fmt.Scanln(&str)
	if err != nil {
		return
	}
	fmt.Printf("After reverse string:\n%s\n", revStr(str))

}
