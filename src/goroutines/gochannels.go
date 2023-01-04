package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4, 5}
	ch := make(chan int, len(a))

	for _, v := range a {
		go func(n int) {
			ch <- n * 2;
		}(v)
	}

	for i := 0; i < len(a); i++ {
		fmt.Println(<- ch);
	}

	fmt.Println("..........CountTo.............")
	val, cancel := countTo(10);
	for i := range val {
		if i > 5 {
			break;
		}
		fmt.Println(i)
	}
	cancel();
}

func countTo(max int) (<- chan int, func()) {
	ch := make(chan int);
	done := make(chan struct{});
	cancel := func() { 
		close(done);
	}
	go func()  {
		for i := 0; i < max; i++ {
			select {
			case <- done:
				return;
			case ch <- i:
			}
		}
		close(ch);
	}()
	return ch, cancel;
}