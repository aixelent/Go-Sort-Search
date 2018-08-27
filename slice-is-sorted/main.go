package main

import (
	"fmt"
	"sort"
)

type language struct {
	Name string
	Year int
}

func (p language) String() string {
	return fmt.Sprintf("Person %s: %d\n", p.Name, p.Year)
}

func main() {
	languages := []language{
		{"Java", 1995},
		{"Scala", 2003},
		{"Rust", 2010},
		{"Python", 1991},
		{"Golang", 2009},
		{"Cobol", 1960},
		{"Elixir", 2012},
		{"Javascript", 1995},
	}

	res := sort.SliceIsSorted(languages, func(i, j int) bool { return languages[i].Year < languages[j].Year })
	fmt.Println("Is slice sorted by Name?: ", res)
	// result looks like a originally slice
	for _, v := range languages {
		fmt.Println(v.Name, v.Year)
	}

	// sorting slice
	sort.Slice(languages, func(i, j int) bool { return languages[i].Year < languages[j].Year })
	// checking if slice has been sorted
	res = sort.SliceIsSorted(languages, func(i, j int) bool { return languages[i].Year < languages[j].Year })
	fmt.Println("\nIs slice sorted by Name?: ", res)
	for _, v := range languages {
		fmt.Println(v.Name, v.Year)
	}
}
