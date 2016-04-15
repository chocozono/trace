package trace

//Tracer is the interface that can record some events
type Tracer interface {
	Trace(...interface{})
}
