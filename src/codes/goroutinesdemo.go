package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	f := func(n int) {
		for i := 0; i < 5; i++ {
			fmt.Println(i, ":", n);
			amt := time.Duration(rand.Intn(250));
			time.Sleep(time.Millisecond * amt);
		}
	}

	for i := 0; i < 10; i++ {
		go f(i);
	}
	var input string;
	fmt.Scanln(&input);
}