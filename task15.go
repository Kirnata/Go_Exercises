package main

import "fmt"

// К каким негативным последствиям может привести данный фрагмент кода, и как это исправить?
// Приведите корректный пример реализации.

//var justString string
//func someFunc() {
//	v := createHugeString(1 << 10)
//	justString = v[:100]
//}
//
//func main() {
//	someFunc()
//}

// 1.Сделаем переменную justString локальной:
// в многопоточных приложениях использование глобальных переменных
// разными потоками приводит к росту ошибок
// 2) Использование string для взятия фрагмента строки может привести к некорректному выводу,
// если символ занимает более одного байта

func createHugeString(length int) []rune {
	str := make([]rune, 0, length)
	for i := 0; i < length; i++ {
		str = append(str, '乑')
	}
	return str
}

func main() {
	var justString string
	v := createHugeString(1 << 10)
	justString = string(v[:10])
	fmt.Println(justString)
}
