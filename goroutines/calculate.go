package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Printf("goroutine demo\n");

	// run function in background
	go calculate()

	index := 0

	go func() {
		for index < 6 {
			fmt.Printf("go func() index = %d \n", index)
			var result float64 = 2.4 * float64(index)
			fmt.Printf("go func() result = %.2f\n", result)

			time.Sleep(500 * time.Millisecond)
			index++
		}
	}()

	// run in the background
	go fmt.Printf("go printed in the background \n")

	// press ENTER to exit
	var input string
	fmt.Scanln(&input)
	fmt.Println("Done")
}

func calculate() {
	i := 12
	for i < 15 {
		fmt.Printf("calculate() = %d\n", i)
		var result float64 = 8.4 * float64(i)
		fmt.Printf("calculate() result= %.2f \n", result)

		time.Sleep(500 * time.Microsecond)
		i++
	}
}