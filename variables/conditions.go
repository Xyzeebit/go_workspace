package main

import ("fmt")

func main() {
	a := 5;
	b := 3;

	fmt.Println(a > b)
	fmt.Println(a < b)
	fmt.Println(a == b)

	if a > b || a - b < a {
		fmt.Println("conditional -->a>b || a-b<a");
	} else {
		fmt.Println("...another");
	}

	option := "a";

	switch option {
	case "c":
		fmt.Println("option is ", option);
	case "b":
		fmt.Println("option is ", option);
	case "a":
		fmt.Println("option is ", option);
	default:
		fmt.Println("another");
	}

	for i:=5; i < 10; i++ {
		fmt.Println(i);
	}

	// Go while loop
	i := 0;
	for i < 10 {
		fmt.Print(i);
		i++;
		if i > 7 {
			break;
		}
	}
	
}