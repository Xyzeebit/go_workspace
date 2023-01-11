package main

import (
	"eventsemitter/events"
	"fmt"
	"math/rand"
)

func main() {
	emitter := events.EventsEmitter[int]();
	
	emitter.On("hello", func(n int) {
		fmt.Println(n)
	});
	emitter.Emit("hello", 20);

	emitter.On("count", counter);
	emitter.On("count", countdown);
	startCounting(emitter);
}

func counter(n int) {
	fmt.Print("counter: ", n, " ");
}

func countdown(n int) {
	fmt.Println("count down:", n - 1);
}

func startCounting(emitter *events.Emitter[int]) {
	for i := 0; i < 10; i++ {
		n := rand.Intn(10);
		emitter.Emit("count", n);
	}
}