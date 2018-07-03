package goutil

import (
	"testing"
)

var (
	benchResult       string
	benchReverseInput map[int]string
)

func init() {
	benchReverseInput = map[int]string{
		0:    "",
		1:    RandAlphaNumeric(1),
		2:    RandAlphaNumeric(2),
		3:    RandAlphaNumeric(3),
		10:   RandAlphaNumeric(10),
		20:   RandAlphaNumeric(20),
		30:   RandAlphaNumeric(30),
		100:  RandAlphaNumeric(100),
		200:  RandAlphaNumeric(200),
		300:  RandAlphaNumeric(300),
		1000: RandAlphaNumeric(1000),
		2000: RandAlphaNumeric(2000),
		3000: RandAlphaNumeric(3000),
	}
}

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

func benchReverse(s string, b *testing.B) {
	var r string
	for n := 0; n < b.N; n++ {
		r = Reverse(s)
	}
	benchResult = r
}

func BenchmarkReverse0(b *testing.B)    { benchReverse(benchReverseInput[0], b) }
func BenchmarkReverse1(b *testing.B)    { benchReverse(benchReverseInput[1], b) }
func BenchmarkReverse2(b *testing.B)    { benchReverse(benchReverseInput[2], b) }
func BenchmarkReverse3(b *testing.B)    { benchReverse(benchReverseInput[3], b) }
func BenchmarkReverse10(b *testing.B)   { benchReverse(benchReverseInput[10], b) }
func BenchmarkReverse20(b *testing.B)   { benchReverse(benchReverseInput[20], b) }
func BenchmarkReverse30(b *testing.B)   { benchReverse(benchReverseInput[30], b) }
func BenchmarkReverse100(b *testing.B)  { benchReverse(benchReverseInput[100], b) }
func BenchmarkReverse200(b *testing.B)  { benchReverse(benchReverseInput[200], b) }
func BenchmarkReverse300(b *testing.B)  { benchReverse(benchReverseInput[300], b) }
func BenchmarkReverse1000(b *testing.B) { benchReverse(benchReverseInput[1000], b) }
func BenchmarkReverse2000(b *testing.B) { benchReverse(benchReverseInput[2000], b) }
func BenchmarkReverse3000(b *testing.B) { benchReverse(benchReverseInput[3000], b) }

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
