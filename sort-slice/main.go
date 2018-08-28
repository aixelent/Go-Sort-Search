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
	return fmt.Sprintf("Languageeeee %s: %d\n", p.Name, p.Year)
}

func main() {
	languages := []language {
		{"Java", 1995},
		{"Scala", 2003},
		{"Rust", 2010},
		{"Python", 1991},
		{"Golang", 2009},
		{"Cobol", 1960},
		{"Elixir", 2012},
		{"Javascript", 1995},
	}

	sort.Slice(languages, func(i ,j int) bool { return languages[i].Name < languages[j].Name })
	fmt.Printf("Sorting by Name (ascending): \n%s", languages)

	sort.Slice(languages, func(i, j int) bool { return languages[i].Name > languages[j].Name })
	fmt.Printf("\nSorting by Name (ascending): \n%s", languages)

	sort.Slice(languages, func(i, j int) bool { return languages[i].Year > languages[j].Year })
	fmt.Printf("\nSorting by Year (ascending): \n%s", languages)

	sort.Slice(languages, func(i, j int) bool { return languages[i].Year < languages[j].Year })
	fmt.Printf("\nSorting by Year (ascending): \n%s", languages)
}
