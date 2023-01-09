package events

type Fn[F any] func(F);

type Emitter struct {
	events map[string][]interface{}
}



func EventsEmitter() Emitter {
	var events Emitter;
	return events
}