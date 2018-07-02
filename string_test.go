package goutil

import (
	"testing"
)

func TestReverse(t *testing.T) {
	cases := []struct {
		name, input, expect string
	}{
		{name: "Reverse 1", input: "", expect: ""},
		{name: "Reverse 2", input: "x", expect: "x"},
		{name: "Reverse 3", input: " ", expect: " "},
		{name: "Reverse 4", input: "ab", expect: "ba"},
		{name: "Reverse 5", input: "界", expect: "界"},
		{name: "Reverse 6", input: "世界", expect: "界世"},
		{name: "Reverse 7", input: "Hello, world", expect: "dlrow ,olleH"},
		{name: "Reverse 8", input: "Hello, 世界", expect: "界世 ,olleH"},
		{name: "Reverse 9", input: "a\nb", expect: "b\na"},
		{name: "Reverse 10", input: "ab\n", expect: "\nba"},
		{name: "Reverse 11", input: "a\t", expect: "\ta"},
	}
	t.Logf("Reverse #: \"input\" -> (\"expect\", \"actual\")")
	for _, c := range cases {
		actual := Reverse(c.input)
		if actual != c.expect {
			t.Errorf("%v: %q -> (%q, %q)", c.name, c.input, c.expect, actual)
		}
	}
}

func TestConcat(t *testing.T) {
	cases := []struct {
		name, inputA, inputB, expectStr string
		expectInt                       int
	}{
		{name: "Concat 1", inputA: "", inputB: "", expectStr: "", expectInt: 0},
		{name: "Concat 2", inputA: "", inputB: " ", expectStr: " ", expectInt: 1},
		{name: "Concat 3", inputA: " ", inputB: "", expectStr: " ", expectInt: 1},
		{name: "Concat 4", inputA: "a", inputB: "", expectStr: "a", expectInt: 1},
		{name: "Concat 5", inputA: "", inputB: "a", expectStr: "a", expectInt: 1},
		{name: "Concat 6", inputA: "\n", inputB: "", expectStr: "\n", expectInt: 1},
		{name: "Concat 7", inputA: "\tx", inputB: "x\t", expectStr: "\txx\t", expectInt: 4},
		{name: "Concat 8", inputA: "\n", inputB: "\b", expectStr: "\n\b", expectInt: 2},
		{name: "Concat 9", inputA: "Hello, ", inputB: "world!", expectStr: "Hello, world!", expectInt: 13},
		{name: "Concat 10", inputA: "Hello, ", inputB: "世界", expectStr: "Hello, 世界", expectInt: 9},
		{name: "Concat 11", inputA: "世", inputB: "", expectStr: "世", expectInt: 1},
		{name: "Concat 12", inputA: "", inputB: "界", expectStr: "界", expectInt: 1},
		{name: "Concat 13", inputA: "世\na", inputB: "b\t界", expectStr: "世\nab\t界", expectInt: 6},
		{name: "Concat 14", inputA: "\t世", inputB: "世\t", expectStr: "\t世世\t", expectInt: 4},
	}
	t.Logf("Concat #: \"inputA, inputB\" -> ([\"expectStr\", expectInt], [\"actualStr\", actualInt])")
	for _, c := range cases {
		actualStr, actualInt := Concat(c.inputA, c.inputB)
		if actualStr != c.expectStr || actualInt != c.expectInt {
			t.Errorf("%v: %q, %q -> ([%q, %v], [%q, %v])",
				c.name, c.inputA, c.inputB, c.expectStr, c.expectInt, actualStr, actualInt)
		}
	}
}
