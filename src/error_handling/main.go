package main

import "fmt"

func main() {
	// calculate()
	// demoPanic()
	calculate2()
}

func calculate() {
	fmt.Println("---------------Error handling------------")
	a := 10
	b := 0
	c := 0

	c = a / b

	fmt.Printf("result = %.2f\n", c)
	fmt.Println("Done")
}

func demoPanic() {
	fmt.Println("---------------Error handling------------")
	defer func() {
		fmt.Println("do something")
	}()
	panic("this is a panic from demoPanic()")
	fmt.Println("this message will never show")
}

func calculate2() {
	fmt.Println("---------------Error handling------------")
	c := 0
	defer func ()  {
		a := 10
		b := 0

		c = a / b

		if error := recover(); error != nil {
			fmt.Println("Recovering...", error)
			fmt.Println(error)
		}
	}()
	fmt.Printf("Result = %d \n", c);
	fmt.Println("Done")

}