package main

import "fmt"

// Имеется последовательность строк - (cat, cat, dog, cat, tree)
// создать для нее собственное множество.
func main() {
	strs := []string{"cat", "cat", "dog", "cat", "tree"}
	set := make([]string, 0)
	m := make(map[string]int)
	for _, v := range strs {
		m[v]++
		if m[v] == 1 {
			set = append(set, v)
		}
	}
	fmt.Printf("start strings: %v\n", strs)
	fmt.Printf("result set: %v\n", set)
}
