package internals

import (
	"bytes"
	"errors"
)

var (
	ErrBadControlSequence = errors.New("malformed control sequence")
)

type SequenceData struct {
	Prefix  byte
	Params  [][]byte
	Inters  []byte
	Command byte
}

func ParseControlSequence(v []byte) (*SequenceData, error) {
	// Immediatly reject any malformed control sequence: it must start with the
	// escape character, and contain at least one prefix and command byte.
	if len(v) < 3 || v[0] != EscapeCode {
		return nil, ErrBadControlSequence
	}

	// Everything between the prefix and the command bytes are arguments: we
	// need to determine where parameters end and intermediate char begin.
	var i int
	end := len(v) - 1
	for i = end - 1; IsInterChar(v[i]); i-- {
	}

	// Value of i marks the separation between (semicolon-separated) parameters
	// and intermediate bytes. One catch: when no parameters are specified, we
	// want to have [][]byte{} rather than [][]byte{[]byte{}}.
	params := [][]byte{}
	if i >= 2 {
		params = bytes.Split(v[2:i+1], []byte{';'})
	}

	return &SequenceData{
		Prefix:  v[1],
		Params:  params,
		Inters:  v[i+1 : end],
		Command: v[end],
	}, nil
}
