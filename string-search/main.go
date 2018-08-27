package main

import (
	"sort"
	"fmt"
)

func main() {
	stringSlice := []string{"Java", "Scala", "Golang", "Python", "Rust", "Cobol", "Elixir", "Javascript"}
	fmt.Printf("Slice: %s", stringSlice)
	fmt.Println("\nIs sorted?:", sort.StringsAreSorted(stringSlice))
	elem := "Python"

	// sorting
	sort.Strings(stringSlice)
	fmt.Println("\nIs sorted?:", sort.StringsAreSorted(stringSlice))
	pos := sort.SearchStrings(stringSlice, elem)
	fmt.Printf("Slice: %s\nElement '%s' found at position %d:\n\n", stringSlice, elem, pos)
}
