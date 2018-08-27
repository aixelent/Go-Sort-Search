package main

import (
	"fmt"
	"sort"
)

func main() {
	ints := []int{1, 5, 3, 9, 7, 13, 11, 15, 17, 23, 19, 21}
	fmt.Println(ints)

	// sorting slice by ascending
	sort.Sort(sort.IntSlice(ints))
	fmt.Printf("Slice sorted by ascending: %d", ints)

	// reversing slice
	sort.Sort(sort.Reverse(sort.IntSlice(ints)))
	fmt.Print("\nReversed slice: ", ints)
}
