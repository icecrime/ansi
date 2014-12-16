package ansi

import (
	"os"
	"testing"
)

func TestANSI(t *testing.T) {
	b := "testing\x1B[1A one two\x1B[1;2;3;mtree\x1B@test"
	w := NewWriter(os.Stdout)
	w.Write([]byte(b))
}
