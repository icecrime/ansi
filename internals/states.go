package internals

import "io"

const (
	EscapeCode = '\x1B'
)

type stateFn func(*Lexer) stateFn

func lexBytes(l *Lexer) stateFn {
	for l.pos < len(l.input) {
		if l.input[l.pos] == EscapeCode {
			if l.pos > l.start {
				l.Emit(RawBytes)
			}
			return lexEscapeSequence
		}
		if _, err := l.Next(); err == io.EOF {
			break
		}
	}

	if l.pos > l.start {
		l.Emit(RawBytes)
	}
	l.Emit(EOF)
	return nil
}

func lexEscapeSequence(l *Lexer) stateFn {
	l.pos += 1 // Drop the ESC byte
	next, _ := l.Next()
	if next == '[' {
		return lexControlSequence
	} else if next >= '@' && next <= '_' {
		return lexTwoCharSequence
	}
	return nil
}

func lexTwoCharSequence(l *Lexer) stateFn {
	l.Emit(TwoCharSequence)
	return lexBytes
}

func lexControlSequence(l *Lexer) stateFn {
	// General form of a control sequence is:
	// ESC [ [private mode characters] n1 ; n2... letter
	paramRange := []byte("0123456789;")
	isCommandChar := func(b byte) bool { return b >= 0x40 && b <= 0x7E }
	isPrivateModeChar := func(b byte) bool { return b >= 0x20 && b <= 0x2F }

	// A control sequence is invalid if it doesn't end with a command char: in
	// this case we revert to treat it as a bunch of raw bytes.
	l.AcceptRunFn(isPrivateModeChar)
	l.AcceptRun(paramRange)
	if !l.AcceptFn(isCommandChar) {
		return l.Cancel(lexBytes)
	}

	l.Emit(ControlSequence)
	return lexBytes
}
