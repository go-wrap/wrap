package wrap

import (
	"reflect"
	"testing"
)

func same(a, b interface{}) bool { return reflect.DeepEqual(a, b) }

func equals(t *testing.T, a, b interface{}) {
	t.Helper()
	if !same(a, b) {
		t.Errorf("\nexpected: %v\n  actual: %v\nto be equal", a, b)
	}
}

var wrapTests = []struct {
	in  string
	out string
	n   int
}{
	{"", "", 10},
	{"Hello world!", "Hello worl\nd!", 10},
	{"Hello world!", "Hello world!", 20},
	{"Hello\nworld!", "Hello\nworld!", 10},
}

func TestWrap(t *testing.T) {
	for _, tt := range wrapTests {
		out := Wrap(tt.in, tt.n)
		equals(t, out, tt.out)
	}
}

var spaceTests = []struct {
	in  string
	out string
	n   int
}{
	{"", "", 10},
	{"Hello world!", "Hello\nworld!", 5},
	{"Hello world!", "Hello\nworld!", 10},
	{"Hello world!", "Hello world!", 20},
	{"Hello\nworld!", "Hello\nworld!", 10},
}

func TestSpace(t *testing.T) {
	for _, tt := range spaceTests {
		out := Space(tt.in, tt.n)
		equals(t, out, tt.out)
	}
}
