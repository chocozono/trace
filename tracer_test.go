package trace

import (
	"bytes"
	"testing"
)

func TestNew(t *testing.T) {
	var buf bytes.Buffer
	tracer := New(&buf)
	if tracer == nil {
		t.Error("Return value from New() is null")
	} else {
		tracer.Trace("Hello, trace package")
		if buf.String() != "Hello, trace package\n" {
			t.Errorf("Wrong output value: '%s'", buf.String())
		}
	}
}
