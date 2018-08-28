package main

import (
	"sort"
	"fmt"
)

func main() {
	f64 := []float64{8.89024, 3.14, 8.8576, 7.5648, 8.125, 10.1029484, 10.10, 11.387492, 11.387491}
	elem := 10.10
	sort.Float64s(f64)
	s := sort.SearchFloat64s(f64, elem)
	fmt.Println("Is sorted?: ", sort.Float64sAreSorted(f64))
	fmt.Printf("Slice: %f\nElement '%f' found at position %d:\n\n", f64, elem, s)

	elem = 10.1029484
	s = sort.Search(len(f64), func(s int) bool { return f64[s] >= elem})
	fmt.Printf("Slice: %f\nElement '%f' found at position %d:\n\n", f64, elem, s)
}
