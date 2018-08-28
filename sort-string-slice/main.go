package main

import (
	"fmt"
	"sort"
)

func main() {
	stringSlice := []string{"Java", "Scala", "Golang", "Rust", "Cobol", "Elixir", "Python", "Javascript"}
	fmt.Println("Is sorted?:", sort.StringsAreSorted(stringSlice))
	fmt.Println("Before sorting:", stringSlice)

	// sorting
	sort.Strings(stringSlice)
	fmt.Println("\nIs sorted?:", sort.StringsAreSorted(stringSlice))
	fmt.Println("After sorting:", stringSlice)
}
