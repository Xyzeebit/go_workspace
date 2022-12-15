package main

import "fmt"

func main() {
	var x int;

	x = 10;

	fmt.Println(x);
	fmt.Println(&x);

	var num *int;
	val := new(int);

	num = new(int);
	*num = x; // set value

	val = &x; // set address

	fmt.Println("===pointer num===")
	fmt.Println(*num) // print a value of x
	fmt.Println(num) // print address of x
	fmt.Println("===pointer val===")
	fmt.Println(*val) // print a value of x
	fmt.Println(val) // 
}