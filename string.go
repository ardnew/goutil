// Package goutil provides commonly used convenience functions for manipulating
// certain data types and also provides reusable type-generic algorithms
//
//   file: string.go
//   desc: routines for manipulating strings
//   auth: ardnew
//
package goutil

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"fmt"
	"math/rand"
	"time"
	"unicode/utf8"
)

var (
	alphaRune        []rune
	alphaNumericRune []rune
)

func init() {
	alphaRune = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	alphaNumericRune = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	rand.Seed(time.Now().UTC().UnixNano())
}

func MD5(s string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(s)))
}

func SHA1(s string) string {
	return fmt.Sprintf("%x", sha1.Sum([]byte(s)))
}

func SHA256(s string) string {
	return fmt.Sprintf("%x", sha256.Sum256([]byte(s)))
}

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

func RuneConcat(s, t string) (string, int) {
	sn := utf8.RuneCountInString(s)
	tn := utf8.RuneCountInString(t)
	// some initial checks to see if we can skip the alloc/copying entirely
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

func Concat(s, t string) (string, int) {
	sn := len(s)
	tn := len(t)
	// some initial checks to see if we can skip the alloc/copying entirely
	switch {
	case 0 == sn && 0 == tn:
		return "", 0
	case 0 == sn:
		return t, tn
	case 0 == tn:
		return s, sn
	}
	r := make([]byte, sn+tn)
	m := copy(r, s)
	n := copy(r[m:], t)
	return string(r), m + n
}

func RandAlpha(n uint) string {
	s := make([]rune, n)
	m := len(alphaRune)
	for i := uint(0); i < n; i++ {
		s[i] = alphaRune[rand.Intn(m)]
	}
	return string(s)
}

func RandAlphaNumeric(n uint) string {
	s := make([]rune, n)
	m := len(alphaNumericRune)
	for i := uint(0); i < n; i++ {
		s[i] = alphaNumericRune[rand.Intn(m)]
	}
	return string(s)
}

func SizeStr(bytes int64, showBytes bool) string {

	kb := float32(bytes) / 1024.0
	mb := float32(kb) / 1024.0
	gb := float32(mb) / 1024.0
	tb := float32(gb) / 1024.0
	ss := ""

	switch {
	case uint64(tb) > 0:
		ss = fmt.Sprintf("%.3g TiB", tb)
	case uint64(gb) > 0:
		ss = fmt.Sprintf("%.3g GiB", gb)
	case uint64(mb) > 0:
		ss = fmt.Sprintf("%.3g MiB", mb)
	case uint64(kb) > 0:
		ss = fmt.Sprintf("%.3g KiB", kb)
	default:
		ss = fmt.Sprintf("%d B", bytes)
		showBytes = false
	}

	if showBytes {
		ss = fmt.Sprintf("%s (%d B)", ss, bytes)
	}
	return ss
}
