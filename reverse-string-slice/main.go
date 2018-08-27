package main

import (
	"sort"
	"fmt"
)

func main()  {
	strings := []string{"Java", "Scala", "Golang", "Python", "Rust", "Cobol", "Elixir", "Javascript"}
	fmt.Println(strings)

	// sorting slice by ascending without 'Reverse'
	sort.Sort(sort.StringSlice(strings))
	fmt.Printf("\nSlice sorted by ascending: %s", strings)

	// Reversing originally slice without sorting
	// Before reversing slice will be sorted
	sort.Sort(sort.Reverse(sort.StringSlice(strings)))
	fmt.Printf("\n\nSlice sorted by ascending: %s", strings)

}
