package events

import "eventsemitter/events"

type Fn[T any] func(T)

type Emitter[T any] struct {
	events map[string][]Fn[T]
}

func (e *Emitter[T]) emit(event string, data T) (evt Emitter[T]) {
	listeners, ok := e.events[event]
	if !ok {
		return
	}
	for _, listener := range listeners {
		listener(data)
	}
	return *e;
}

func (e *Emitter[T]) on(event string, fn Fn[T]) (evt Emitter[T]) {

}

func EventsEmitter[T any]() Emitter[T] {
	var events Emitter[T]
	return events
}