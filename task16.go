package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

// Реализовать быструю сортировку массива (quicksort) встроенными методами языка
func main() {
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	numsRune := strings.Split(input.Text(), " ")
	nums := make([]int, 0, len(numsRune))
	for _, v := range numsRune {
		buf, err := strconv.Atoi(v)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			return
		}
		nums = append(nums, buf)
	}
	sort.Ints(nums)
	fmt.Println(nums)
}
