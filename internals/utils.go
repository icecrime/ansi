package internals

const (
	EscapeCode = '\x1B'
)

func IsParamChar(b byte) bool {
	return b >= 0x30 && b <= 0x3F
}

func IsInterChar(b byte) bool {
	return b >= 0x20 && b <= 0x2F
}

func IsCommandChar(b byte) bool {
	return b >= 0x40 && b <= 0x7E
}
