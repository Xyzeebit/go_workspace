package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("array definitions");
	var numbers[5] int;
	var cities[5] string;
	var matrix[3][3] int;

	// insert data
	fmt.Println(">>>>> insert data into array");
	for i := 0; i<5; i++ {
		numbers[i] = i;
		cities[i] = fmt.Sprintf("City %d", i);
	}

	// insert matrix data
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			matrix[i][j] = rand.Intn(100);
		}
	}

	// display data

	fmt.Println("display array data");
	for i := 0; i < 5; i++ {
		fmt.Println(numbers[i]);
		fmt.Println(cities[i]);
	}

	fmt.Println("Display matrix data");
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			fmt.Print(matrix[i][j], " ");
		}
	}
}