package main;

import "fmt";

func main() {
	var str string;
	var n, m int;
	var nm float32;

	str = "Hello, go";
	n = 10;
	m = 50;
	nm = 2.45;

	fmt.Println("value of str=", str);
	fmt.Println("value of n=", n);
	fmt.Println("value of m=", m);
	fmt.Println("value of nm=", nm);

	var city string = "London";
	var x int = 100;

	fmt.Println("value of city=", city);
	fmt.Println("value of x=", x);

	country := "DE";
	val := 15;

	fmt.Println("value of country=", country);
	fmt.Println("value of val=", val);

	// defining multiple variables
	var (
		name string
		email string
		age int
	)
	name = "John"
	email = "john@email.com"
	age = 27

	fmt.Println("value of name=", name);
	fmt.Println("value of email=", email);
	fmt.Println("value of age=", age);
}