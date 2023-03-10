package events

type Fn[T any] func(T)

type Emitter[T any] struct {
	events map[string][]Fn[T]
}

func (e *Emitter[T]) Emit(event string, data T) (evt Emitter[T]) {
	listeners, ok := e.events[event]
	if !ok {
		return
	}
	for _, listener := range listeners {
		listener(data)
	}
	return *e
}

func (e *Emitter[T]) On(event string, fn Fn[T]) (evt Emitter[T]) {
	_, ok := e.events[event]
	if !ok {
		v := []Fn[T]{fn}
		e.events[event] = v
		return *e
	}
	e.events[event] = append(e.events[event], fn)
	return *e
}

func EventsEmitter[T any]() *Emitter[T] {
	return &Emitter[T]{
		events: make(map[string][]Fn[T]),
	}
}