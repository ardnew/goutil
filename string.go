// Package goutil provides commonly used convenience functions for manipulating
// certain data types and also provides reusable type-generic algorithms
//
//   file: file.go
//   desc: routines for manipulating strings
//   auth: ardnew
//
package goutil

import (
	"unicode/utf8"
)

func Reverse(s string) string {
	rn := utf8.RuneCountInString(s)
	if rn < 2 {
		return s
	}
	r := []rune(s)
	for i, j := 0, rn-1; i < rn/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

func Concat(s, t string) (string, int) {
	sn := utf8.RuneCountInString(s)
	tn := utf8.RuneCountInString(t)
	// some initial checks to see if we can skip the
	// alloc/copying entirely
	switch {
	case 0 == sn && 0 == tn:
		return "", 0
	case 0 == sn:
		return t, tn
	case 0 == tn:
		return s, sn
	}
	r := make([]rune, sn+tn)
	m := copy(r, []rune(s))
	n := copy(r[m:], []rune(t))
	return string(r), m + n
}
