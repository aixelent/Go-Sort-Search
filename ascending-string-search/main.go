package main

import (
	"fmt"
	"sort"
)

func main() {
	stringSlice := []string{"Java", "Scala", "Golang", "Python", "Rust", "Cobol", "Elixir", "Javascript"}
	elem := "Python"
	// search without sorting
	s := sort.SearchStrings(stringSlice, elem)
	fmt.Println("Is sorted?: ", sort.StringsAreSorted(stringSlice))
	fmt.Printf("Slice: %s\nElement '%s' found at position %d.\n\n", stringSlice, elem, s)

	// search with sorting
	sort.Strings(stringSlice)
	fmt.Println("Is sorted?: ", sort.StringsAreSorted(stringSlice))
	elem = "Python"
	s = sort.SearchStrings(stringSlice, elem)
	fmt.Printf("Slice: %s\nElement '%s' found at position %d.\n", stringSlice, elem, s)
}
