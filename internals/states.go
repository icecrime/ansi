package internals

import "io"

type stateFn func(*Lexer) stateFn

// Status function lexBytes eats up raw bytes, and switches to lexing an espace
// sequence when it encouters the ESC character code.
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

// Status function lexEscapeSequence lexes an sequence starting with the ESC
// character code. It may be a CSI introduced sequence, or a two char sequence.
func lexEscapeSequence(l *Lexer) stateFn {
	l.pos += 1 // Drop the ESC byte
	next, _ := l.Peek()
	if next == '[' {
		return lexControlSequence
	} else if next >= '@' && next <= '_' {
		return lexTwoCharSequence
	}
	return nil
}

// Status function lexTwoCharSequence.
func lexTwoCharSequence(l *Lexer) stateFn {
	l.pos += 1 // Eat up the command character
	l.Emit(TwoCharSequence)
	return lexBytes
}

// Status function lexControlSequence lexes a CSI introduced sequence. If any
// character in the sequence is out of the allowed range, it falls back to
// lexing raw bytes.
func lexControlSequence(l *Lexer) stateFn {
	// General form of a control sequence is:
	// ESC [ n1 ; n2... [trailiing intermediate characters] letter
	// A control sequence is invalid if it doesn't end with a command char: in
	// this case we revert to treat it as a bunch of raw bytes.
	l.Accept([]byte{'['})
	l.AcceptRunFn(IsParamChar)
	l.AcceptRunFn(IsInterChar)
	if !l.AcceptFn(IsCommandChar) {
		return l.Cancel(lexBytes)
	}

	l.Emit(ControlSequence)
	return lexBytes
}
