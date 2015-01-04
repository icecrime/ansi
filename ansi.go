package ansi

import "io"

func NewWriter(dst io.Writer) io.Writer {
	return writer{
		Writer: dst,
	}
}

type writer struct {
	io.Writer
}

func (w writer) Write(b []byte) (int, error) {
	buf := p.parse(b)
	return w.Writer.Write(buf)
}
