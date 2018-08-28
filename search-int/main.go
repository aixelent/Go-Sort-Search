package main

import (
	"fmt"
	"sort"
)

func main() {
	ints := []int{1, 5, 3, 9, 7, 13, 11, 15, 17, 23, 19, 21}
	elem := 7
	fmt.Println("Is sorted?: ", sort.IntsAreSorted(ints))
	sort.Ints(ints)
	s := sort.SearchInts(ints, elem)
	fmt.Printf("Slice: %d\nElement '%d' found at position %d.\n\n", ints, elem, s)

	elem = 21
	s = sort.Search(len(ints), func(ks int) bool { return ints[ks] >= elem })
	fmt.Println("Is sorted?: ", sort.IntsAreSorted(ints))
	fmt.Printf("Slice: %d\nElement '%d' found at position %d.\n", ints, elem, s)
}
