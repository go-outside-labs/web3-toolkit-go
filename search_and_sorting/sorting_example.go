package main

import (
	"fmt"
	"sort"
)

func main() {

	strs := []string{"c", "a", "b"}

	// sort is in place
	sort.Strings(strs)
	fmt.Println("Strings:", strs)

	ints := []int{7, 2, 4}
	sort.Ints(ints)
	fmt.Println("Ints:   ", ints)

	// check if a slice is already in sorted order
	s := sort.IntsAreSorted(ints)
	fmt.Println("Sorted: ", s)
}
