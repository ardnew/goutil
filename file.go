// Package goutil provides commonly used convenience functions for manipulating
// certain data types and also provides reusable type-generic algorithms
//
//   file: file.go
//   desc: routines for interacting with the filesystem
//   auth: ardnew
//
package goutil

import (
	"os"
)

func PathExists(p string) (bool, os.FileInfo) {
	info, err := os.Stat(p)
	exists := err == nil || !os.IsNotExist(err)
	if !exists {
		return false, nil
	}
	return true, info
}
