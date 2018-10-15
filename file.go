// Package goutil provides commonly used convenience functions for manipulating
// certain data types and also provides reusable type-generic algorithms
//
//   file: file.go
//   desc: routines for interacting with the file system
//   auth: ardnew
//
package goutil

import (
	"os"
)

// function PathExists() determines if a file or directory exists at the given
// path. if it exists, a true value and the associated stat struct is returned.
// otherwise, a false value and a nil stat struct is returned.
func PathExists(p string) (bool, os.FileInfo) {
	info, err := os.Stat(p)
	exists := err == nil || !os.IsNotExist(err)
	if !exists {
		return false, nil
	}
	return true, info
}
