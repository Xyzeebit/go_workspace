package main

import (
	"fmt"
	"math"
)

func main() {
	foo();
	circle_area(3.45)
	calculate(5, 2, 4.56)
	advanced_calculate(2, 3, 7.23)
	x, y, z := compute(5.3, 2.89, 1.01, "energy")
	fmt.Println(x, y, z);
	sum := add(1,2,3,4,5)
	fmt.Println(sum)
	closure_func("closure")
	a := 10
	fib := fibonacci(a)
	fmt.Printf("the fibonacci of %d is %d \n", a, fib)
}

func foo() {
	fmt.Println("foo() was called")
}

func circle_area(r float64) {
	area := math.Pi + math.Pow(r, 2);
	fmt.Printf("Circle area (r=%.2f) = %.2f \n", r, area)
}

func calculate(a, b, c float64) {
	result := a * b * c;
	fmt.Printf("a=%.2f, b=%.2f, c=%.2f Result= %.2f \n", a, b, c, result)
}

// a function with parameters and a return value
func advanced_calculate(a,b, c float64) float64 {
	result := a*b*c
	return result
}

func compute(a, b, c float64, name string) (float64, float64, string) {
	result1 := a*b*c
	result2 := a + b + c
	result3 := result1 + result2
	newName := fmt.Sprintf("%s value = %.2f", name,result3)
	return result1, result2, newName
}

func add(numbers ...int) int {
	result := 0
	for _, number := range numbers {
		result += number
	}

	return result
}

func closure_func(name string) {
	hoo := func(a, b int) {
		result := a * b
		fmt.Printf("hoo() = %d \n", result)
	}
	joo := func(a, b int) int {
		return a * b + a
	}
	fmt.Printf("closure_func(%s) was called\n",name)
	hoo(2,3)
	val := joo(4, 5);
	fmt.Printf("val from joo() = %d \n",val)
}

func fibonacci(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	}
	return (fibonacci(n - 1) + fibonacci(n - 2))
}