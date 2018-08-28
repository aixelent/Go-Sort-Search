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
	return fmt.Sprintf("Language %s: %d\n", p.Name, p.Year)
}

// Sort implements sort.Interface for []language based on the Year field
type Sort []language

func (a Sort) Len() int           { return len(a) }
func (a Sort) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a Sort) Less(i, j int) bool { return a[i].Year < a[j].Year }

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

	fmt.Println("Languages unsorted:\n", languages)
	sort.Sort(Sort(languages))
	fmt.Println("Languages sorted by year:\n", languages)
}
