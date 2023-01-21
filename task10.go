package main

import "fmt"

// Дана последовательность температурных колебаний:
// -25.4, -27.0 13.0, 19.0, 15.5, 24.5, -21.0, 32.5.
// Объединить данные значения в группы с шагом в 10 градусов.
// Последовательность в подмножноствах не важна
func main() {
	temperature := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	tMap := make(map[int][]float64)

	for _, val := range temperature {
		key := (int(val) / 10) * 10
		tMap[key] = append(tMap[key], val)
	}

	for key, val := range tMap {
		fmt.Println(key, ": ", val)
	}
}
