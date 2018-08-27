package main

import (
	"fmt"
	"sort"
)

func main() {
	f64 := []float64{8.89024, 3.14, 8.8576, 7.5648, 8.125, 10.1029484, 10.10, 11.387492, 11.387491}
	fmt.Println("Are sorted?: ", sort.Float64sAreSorted(f64))
	fmt.Println("After sorting:", f64)

	// sorting
	sort.Float64s(f64)
	fmt.Println("Are sorted?:", sort.Float64sAreSorted(f64))
	fmt.Println("Before sorting:", f64)
}
