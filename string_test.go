package goutil

import (
	"testing"
)

type stringPair struct {
	a, b string
}

var (
	benchReverseResult   string
	benchReverseInput    map[int]string
	benchConcatResultStr string
	benchConcatResultInt int
	benchConcatInput     map[int]stringPair
)

func init() {
	benchReverseInput = map[int]string{
		0:    "",
		1:    RandAlphaNumeric(1),
		2:    RandAlphaNumeric(2),
		5:    RandAlphaNumeric(5),
		10:   RandAlphaNumeric(10),
		20:   RandAlphaNumeric(20),
		50:   RandAlphaNumeric(50),
		100:  RandAlphaNumeric(100),
		200:  RandAlphaNumeric(200),
		500:  RandAlphaNumeric(500),
		1000: RandAlphaNumeric(1000),
		2000: RandAlphaNumeric(2000),
		5000: RandAlphaNumeric(5000),
	}
	benchConcatInput = map[int]stringPair{
		0:    stringPair{"", ""},
		1:    stringPair{RandAlphaNumeric(1), RandAlphaNumeric(1)},
		2:    stringPair{RandAlphaNumeric(2), RandAlphaNumeric(2)},
		5:    stringPair{RandAlphaNumeric(5), RandAlphaNumeric(5)},
		10:   stringPair{RandAlphaNumeric(10), RandAlphaNumeric(10)},
		20:   stringPair{RandAlphaNumeric(20), RandAlphaNumeric(20)},
		50:   stringPair{RandAlphaNumeric(50), RandAlphaNumeric(50)},
		100:  stringPair{RandAlphaNumeric(100), RandAlphaNumeric(100)},
		200:  stringPair{RandAlphaNumeric(200), RandAlphaNumeric(200)},
		500:  stringPair{RandAlphaNumeric(500), RandAlphaNumeric(500)},
		1000: stringPair{RandAlphaNumeric(1000), RandAlphaNumeric(1000)},
		2000: stringPair{RandAlphaNumeric(2000), RandAlphaNumeric(2000)},
		5000: stringPair{RandAlphaNumeric(5000), RandAlphaNumeric(5000)},
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
	benchReverseResult = r
}

func BenchmarkReverse0(b *testing.B)    { benchReverse(benchReverseInput[0], b) }
func BenchmarkReverse1(b *testing.B)    { benchReverse(benchReverseInput[1], b) }
func BenchmarkReverse2(b *testing.B)    { benchReverse(benchReverseInput[2], b) }
func BenchmarkReverse5(b *testing.B)    { benchReverse(benchReverseInput[5], b) }
func BenchmarkReverse10(b *testing.B)   { benchReverse(benchReverseInput[10], b) }
func BenchmarkReverse20(b *testing.B)   { benchReverse(benchReverseInput[20], b) }
func BenchmarkReverse50(b *testing.B)   { benchReverse(benchReverseInput[50], b) }
func BenchmarkReverse100(b *testing.B)  { benchReverse(benchReverseInput[100], b) }
func BenchmarkReverse200(b *testing.B)  { benchReverse(benchReverseInput[200], b) }
func BenchmarkReverse500(b *testing.B)  { benchReverse(benchReverseInput[500], b) }
func BenchmarkReverse1000(b *testing.B) { benchReverse(benchReverseInput[1000], b) }
func BenchmarkReverse2000(b *testing.B) { benchReverse(benchReverseInput[2000], b) }
func BenchmarkReverse5000(b *testing.B) { benchReverse(benchReverseInput[5000], b) }

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

func benchRuneConcat(p stringPair, b *testing.B) {
	var s string
	var c int
	for n := 0; n < b.N; n++ {
		s, c = RuneConcat(p.a, p.b)
	}
	benchConcatResultStr = s
	benchConcatResultInt = c
}

func BenchmarkRuneConcat0(b *testing.B)    { benchRuneConcat(benchConcatInput[0], b) }
func BenchmarkRuneConcat1(b *testing.B)    { benchRuneConcat(benchConcatInput[1], b) }
func BenchmarkRuneConcat2(b *testing.B)    { benchRuneConcat(benchConcatInput[2], b) }
func BenchmarkRuneConcat5(b *testing.B)    { benchRuneConcat(benchConcatInput[5], b) }
func BenchmarkRuneConcat10(b *testing.B)   { benchRuneConcat(benchConcatInput[10], b) }
func BenchmarkRuneConcat20(b *testing.B)   { benchRuneConcat(benchConcatInput[20], b) }
func BenchmarkRuneConcat50(b *testing.B)   { benchRuneConcat(benchConcatInput[50], b) }
func BenchmarkRuneConcat100(b *testing.B)  { benchRuneConcat(benchConcatInput[100], b) }
func BenchmarkRuneConcat200(b *testing.B)  { benchRuneConcat(benchConcatInput[200], b) }
func BenchmarkRuneConcat500(b *testing.B)  { benchRuneConcat(benchConcatInput[500], b) }
func BenchmarkRuneConcat1000(b *testing.B) { benchRuneConcat(benchConcatInput[1000], b) }
func BenchmarkRuneConcat2000(b *testing.B) { benchRuneConcat(benchConcatInput[2000], b) }
func BenchmarkRuneConcat5000(b *testing.B) { benchRuneConcat(benchConcatInput[5000], b) }

func benchConcat(p stringPair, b *testing.B) {
	var s string
	var c int
	for n := 0; n < b.N; n++ {
		s, c = Concat(p.a, p.b)
	}
	benchConcatResultStr = s
	benchConcatResultInt = c
}

func BenchmarkConcat0(b *testing.B)    { benchConcat(benchConcatInput[0], b) }
func BenchmarkConcat1(b *testing.B)    { benchConcat(benchConcatInput[1], b) }
func BenchmarkConcat2(b *testing.B)    { benchConcat(benchConcatInput[2], b) }
func BenchmarkConcat5(b *testing.B)    { benchConcat(benchConcatInput[5], b) }
func BenchmarkConcat10(b *testing.B)   { benchConcat(benchConcatInput[10], b) }
func BenchmarkConcat20(b *testing.B)   { benchConcat(benchConcatInput[20], b) }
func BenchmarkConcat50(b *testing.B)   { benchConcat(benchConcatInput[50], b) }
func BenchmarkConcat100(b *testing.B)  { benchConcat(benchConcatInput[100], b) }
func BenchmarkConcat200(b *testing.B)  { benchConcat(benchConcatInput[200], b) }
func BenchmarkConcat500(b *testing.B)  { benchConcat(benchConcatInput[500], b) }
func BenchmarkConcat1000(b *testing.B) { benchConcat(benchConcatInput[1000], b) }
func BenchmarkConcat2000(b *testing.B) { benchConcat(benchConcatInput[2000], b) }
func BenchmarkConcat5000(b *testing.B) { benchConcat(benchConcatInput[5000], b) }
