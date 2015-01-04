package ansi

import (
	"fmt"

	"github.com/icecrime/ansi/internals"
)

func (parser) HandleRawBytes(v []byte) []byte {
	// Pass-through: raw bytes are written unmodified on the underlying stream.
	return v
}

func (parser) HandleControlSequence(v []byte) []byte {
	seq, _ := internals.ParseControlSequence(v)
	fmt.Printf("%v\n", seq)
	return nil
}

func (parser) HandleTwoCharSequence(v []byte) []byte {
	return nil
}
