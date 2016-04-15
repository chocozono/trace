package trace

import "io"

//Tracer is the interface that can record some events
type Tracer interface {
	Trace(...interface{})
}

type tracer struct {
	out io.Writer
}

func (t *tracer) Trace(a ...interface{}) {}

func New(w io.Writer) Tracer {
	return nil
}
