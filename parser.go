package ansi

import (
	"bytes"
	"fmt"
	"reflect"

	"github.com/icecrime/ansi/internals"
)

var (
	p              = parser{}
	defaultHandler = reflect.ValueOf(p).MethodByName("HandleRawBytes")
)

type parser struct {
}

func (p parser) parse(input []byte) []byte {
	out := bytes.Buffer{}
	lex := internals.NewLexer(input)
	for item := lex.NextItem(); item.T != internals.EOF; item = lex.NextItem() {
		fmt.Printf("[%s] %q\n", item.T.String(), string(item.Value))

		handleFn := p.electHandler(item.T)
		if b := p.callHandler(handleFn, item.Value); b != nil {
			out.Write(b)
		}
	}
	return out.Bytes()
}

func (parser) callHandler(handler reflect.Value, input []byte) []byte {
	res := handler.Call([]reflect.Value{reflect.ValueOf(input)})
	return res[0].Interface().([]byte)
}

func (p parser) electHandler(itemType internals.ItemType) reflect.Value {
	// Print everything we don't know how to handle as raw bytes.
	var handleFn reflect.Value
	if handleFn = reflect.ValueOf(p).MethodByName("Handle" + itemType.String()); !handleFn.IsValid() {
		handleFn = defaultHandler
	}
	return handleFn
}
