package internals

import (
	"bytes"
	"testing"
)

func cmpControlSequence(t *testing.T, result, expected *SequenceData) {
	if result.Prefix != expected.Prefix {
		t.Fatalf("Bad prefix for parsed control sequence: got %q, expected %q", result.Prefix, expected.Prefix)
	}

	if len(result.Params) != len(expected.Params) {
		t.Fatalf("Bad length for params array: got %d, expected %d", len(result.Params), len(expected.Params))
	}
	for i, param := range result.Params {
		if exp := expected.Params[i]; bytes.Compare(param, exp) != 0 {
			t.Fatalf("Bad value for parameter %d: got %q, expected %q", i, param, exp)
		}
	}

	if bytes.Compare(result.Inters, expected.Inters) != 0 {
		t.Fatalf("Bad intermediate bytes for parsed control sequence: got %q, expected %q", result.Inters, expected.Inters)
	}

	if result.Command != expected.Command {
		t.Fatalf("Bad command for parsed control sequence: got %q, expected %q", result.Command, expected.Command)
	}
}

func TestCompleteSequence(t *testing.T) {
	b := []byte("\x1B[1;2<;3+m")
	seq, err := ParseControlSequence(b)
	if err != nil {
		t.Fatal(err)
	}

	expected := &SequenceData{
		Prefix:  '[',
		Params:  [][]byte{[]byte("1"), []byte("2<"), []byte("3")},
		Inters:  []byte("+"),
		Command: 'm',
	}
	cmpControlSequence(t, seq, expected)
}

func TestMinimalSequence(t *testing.T) {
	b := []byte("\x1B_t")
	seq, err := ParseControlSequence(b)
	if err != nil {
		t.Fatal(err)
	}

	expected := &SequenceData{
		Prefix:  '_',
		Command: 't',
	}
	cmpControlSequence(t, seq, expected)
}
