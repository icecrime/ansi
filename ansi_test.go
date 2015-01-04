package ansi

import (
	"io/ioutil"
	"testing"
)

func TestANSI(t *testing.T) {
	b := "testing\x1B[1A one two\x1B[1;2;3+mthree\x1B@test\n"
	w := NewWriter(ioutil.Discard)
	w.Write([]byte(b))
}
