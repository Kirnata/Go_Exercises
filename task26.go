package main

import "fmt"

// Разработать программу, которая проверяет, что все символы в строке уникальные (true — если уникальные, false etc).
// Функция проверки должна быть регистронезависимой.
// Например:
// abcd — true
// abCdefAaf — false

func lowerCase(i byte) byte {
	if i >= 'A' && i <= 'Z' {
		return i + 32
	}
	return i
}

func checkUnic(str string) bool {
	arr := make([]int, 25)

	for i := 0; i < len(str); i++ {
		if arr[lowerCase(str[i])-97] > 0 {
			return false
		}
		arr[str[i]-97]++
	}
	return true
}
func main() {
	var str string

	_, err := fmt.Scanln(&str)
	if err != nil {
		return
	}

	fmt.Println(checkUnic(str))
}
