package internals

type ItemType int

// Item models a lexer token. The value it hols is the raw slice of bytes
// extracted from the stream.
type Item struct {
	T     ItemType
	Value []byte
}

// For references of possible escape sequences, refer to Standard ECMA-48:
// http://www.ecma-international.org/publications/files/ECMA-ST/Ecma-048.pdf
const (
	EOF ItemType = iota
	RawBytes
	TwoCharSequence
	ControlSequence
)

var (
	ItemTypeName = []string{
		"EOF",
		"RawBytes",
		"TwoCharSequence",
		"ControlSequence",
	}
)

func (t ItemType) String() string {
	return ItemTypeName[t]
}
