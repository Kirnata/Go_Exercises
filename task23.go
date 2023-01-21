package main

import "fmt"

func removeByPos(nums []int, i int) []int {
	switch {
	case i < 0 || i >= len(nums):
		fmt.Println("out of range position")
		return nums
	case i == len(nums)-1:
		return nums[:i]
	default:
		return append(nums[:i], nums[i+1:]...)
	}
}

// Удалить i-ый элемент из слайса
func main() {
	var n int
	fmt.Scan(&n)
	nums := make([]int, n)
	for i := range nums {
		fmt.Scan(&nums[i])
	}
	var i int
	fmt.Scan(&i)
	resNums := removeByPos(nums, i)
	fmt.Printf("original array: %v\n", nums)
	fmt.Printf("array after remove: %v\n", resNums)
}
