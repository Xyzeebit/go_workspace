package events

type Fn[T any] func(T);

type Emitter struct {
	events map[string][]Fn;
}

func (e *Emitter) emit(event string, data interface{}) (evt Emitter) {
	listeners, ok := e.events[event];
	if !ok {
		return;
	}
	for _, listener := range listeners {
		listener(data);
	}
}

func EventsEmitter() Emitter {
	var events Emitter;
	return events
}