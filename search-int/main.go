package main

import (
	"sort"
	"fmt"
)

func main() {
	ints := []int{1, 5, 3, 9, 7, 13, 11, 15, 17, 23, 19, 21}
	// searching element
	elem := 11
	// searching in (where, what)
	s := sort.SearchInts(ints, elem)
	fmt.Println("Is sorted?: ", sort.IntsAreSorted(ints))
	// elem found at position '5' because slice are sorted before searching
	// adn result is not sorted in the output
	fmt.Printf("Slice: %d\nElement '%d' found at position %d.\n\n", ints, elem, s)

	elem = 3
	sort.Ints(ints)
	fmt.Println("Is sorted?: ", sort.IntsAreSorted(ints))
	s = sort.SearchInts(ints, elem)
	fmt.Printf("Slice: %d\nElement '%d' found at position %d.\n\n", ints, elem, s)

	elem = 21
	sort.Ints(ints)
	fmt.Println("Is sorted?: ", sort.IntsAreSorted(ints))
	s = sort.Search(len(ints), func(i int) bool { return ints[i] >= elem })
	fmt.Printf("Slice: %d\nElement '%d' found at position %d.\n", ints, elem, s)
}
