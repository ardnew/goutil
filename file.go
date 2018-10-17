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

// function PathExists() determines if a file or directory -can be verified- to
// exist at the given path under the effective permissions of this running
// process. if it exists, a true value and the associated stat struct is
// returned. otherwise, a false value and a nil stat struct is returned.
//
//  N.B., this function provides no way of distinguishing an existing file with
//        insufficient permissions and a non-existing file! in both cases, you
//        will receive a false return because in both cases there does not exist
//        a file that can be referenced at the given path.
//
func PathExists(p string) (bool, os.FileInfo) {

	info, err := os.Stat(p)
	if err != nil && (os.IsNotExist(err) || os.IsPermission(err)) {
		return false, nil
	}
	return true, info
}
