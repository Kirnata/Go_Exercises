package main

import (
	"fmt"
	"sort"
)

func beSearch(nums []int, n int) int {
	left := 0
	right := len(nums)
	for left <= right {
		mid := (left + right) / 2
		switch {
		case nums[mid] == n:
			return mid
		case nums[mid] > n:
			right = mid - 1
		case nums[mid] < n:
			left = mid + 1
		}
	}
	return -1
}

// Реализовать быструю сортировку массива (quicksort) встроенными методами языка
func main() {
	var n int
	fmt.Scan(&n)
	nums := make([]int, n)
	for i := range nums {
		fmt.Scan(&nums[i])
	}
	sort.Ints(nums)

	var find int
	fmt.Scan(&find)
	//return
	pos := beSearch(nums, find)
	if pos < 0 {
		fmt.Printf("value %d not found\n", find)
	} else {
		fmt.Printf("value %d found at position %d\n", find, pos)
	}
}
