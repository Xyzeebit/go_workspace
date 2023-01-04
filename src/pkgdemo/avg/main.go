package main

import (
	"fmt"
	"pkgdemo/math"
)

func main() {
	xs := []float64{1, 2, 3, 4, 0.3}
	avg := math.Average(xs)
	fmt.Println(avg);
	min := math.Min(xs);
	fmt.Println(min);
	max := math.Max(xs);
	fmt.Println(max);
}