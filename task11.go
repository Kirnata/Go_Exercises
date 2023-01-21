package main

import "fmt"

func ftInter(arr1, arr2 []int) []int {
	res := make([]int, 0)
	myMap := make(map[int]int)
	for _, v := range arr1 {
		myMap[v]++
	}
	for _, v := range arr2 {
		if _, ok := myMap[v]; ok {
			res = append(res, v)
			myMap[v]--
			if myMap[v] < 1 {
				delete(myMap, v)
			}
		}
	}
	return res
}

// Реализовать пересечение двух неупорядоченных множеств
func main() {
	arr1 := []int{2, 2, 2}
	arr2 := []int{1, 5, 2, 8, -2, 5, 2}
	fmt.Println(ftInter(arr1, arr2))
}
