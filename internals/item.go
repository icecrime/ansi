package internals

type ItemType int

type Item struct {
	T     ItemType
	Value []byte
}

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
