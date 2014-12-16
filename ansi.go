package ansi

import (
	"fmt"
	"io"
	"reflect"

	"github.com/docker/ansi/internals"
)

func NewWriter(dst io.Writer) io.Writer {
	return writer{
		Writer: dst,
	}
}

type writer struct {
	io.Writer
}

func (w writer) Write(p []byte) (int, error) {
	buf := parser{}.parse(p)
	return w.Writer.Write(buf)
}

type parser struct {
}

func (p parser) parse(input []byte) []byte {
	l := internals.NewLexer(input)
	for item := l.NextItem(); item.T != internals.EOF; item = l.NextItem() {
		fmt.Printf("[%s] %q\n", item.T.String(), string(item.Value))

		// Print everything we don't know how to handle as raw bytes.
		var handleFn reflect.Value
		if handleFn = reflect.ValueOf(p).MethodByName("Handle" + item.T.String()); !handleFn.IsValid() {
			handleFn = reflect.ValueOf(p).MethodByName("HandleRawBytes")
		}
		handleFn.Call([]reflect.Value{reflect.ValueOf(item.Value)})
	}
	return nil
}

func (parser) HandleRawBytes(v []byte) {
}

func (parser) HandleControlSequence(v []byte) {
}

func (parser) HandleTwoCharSequence(v []byte) {
}
