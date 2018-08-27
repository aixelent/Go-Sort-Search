package main

import (
	"fmt"
	"sort"
)

func main() {
	ints := []int{1, 5, 3, 9, 7, 13, 11, 15, 17, 23, 19, 21}
	fmt.Println("Is sorted?:", sort.IntsAreSorted(ints))
	fmt.Println("Before sorting:", ints)

	// sorting
	sort.Ints(ints)
	fmt.Println("Is sorted?:", sort.IntsAreSorted(ints))
	fmt.Println("After sorting:", ints)
}
