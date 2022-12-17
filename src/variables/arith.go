package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b int;

	a = 5;
	b = 10;
	c := a + b;

	fmt.Printf("%d + %d = %d \n", a, b, c);

	d := b - a;

	fmt.Printf("%d - %d = %d \n", b, a, d);

	e := float32(a) / float32(b);

	fmt.Printf("%d / %d = %.2f \n", b, a, e);

	f := a * b;

	fmt.Printf("%d * %d = %d \n", a, b, f);

	// ..
	/* Mathematical functions */
	// ..

	n := 2.4;
	y := math.Pow(n, 2);

	fmt.Printf("%.2f^%d = %.2f \n", n, 2, y)

	m := math.Sin(n);

	fmt.Printf("Sin(%.2f) = %.2f \n", n, m);

	square := math.Sqrt(n * 3.5);

	fmt.Printf("Sqrt(%.2f * %.2f) = %.2f \n", n, e, square);

	var k int = 10;
	fmt.Printf("k is %d \n", k);
	k++;
	fmt.Printf("k++ is now %d \n", k);

}