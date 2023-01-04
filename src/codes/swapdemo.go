package main

import "fmt"

func main() {
	x := 1
	y := 2

	fmt.Printf("before swap x=%d, y=%d\n", x, y);
	swap(&x, &y)
	fmt.Printf("after swap x=%d, y=%d\n", x, y);
}

func swap(x, y *int) {
	tmp := *x
	*x = *y
	*y = tmp
}