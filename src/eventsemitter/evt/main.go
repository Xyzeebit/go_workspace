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
	for i := 0; i < 10; i++ {
		n := rand.Intn(10);
		emitter.Emit("hello", n);
	}
}