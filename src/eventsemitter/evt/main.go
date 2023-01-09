package main

import (
	"eventsemitter/events"
	"fmt"
)

func main() {
	emitter := events.EventsEmitter[interface{}]();
	
	emitter.on("hello", func(n interface{}) {
		fmt.Println(n)
	});
	emitter.emit("hello", 20);
}