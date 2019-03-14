package main

import (
	"fmt"
	"sort"
)

func SortInts64(a []int64) {
	sort.Slice(a, func(i, j int) bool {
		return a[i] < a[j]
	})
}

func main() {
	a := []int64{2, 3, 8, 9, 12, 0, 1, -10}
	SortInts64(a)
	fmt.Println(a)
}
