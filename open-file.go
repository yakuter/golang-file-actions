package main

import(
	"os"
)

func main () {


	// CONSTANTS TO USE WHILE OPENING A FILE
	/*
	const (
		// Exactly one of O_RDONLY, O_WRONLY, or O_RDWR must be specified.
		O_RDONLY int = syscall.O_RDONLY // open the file read-only.
		O_WRONLY int = syscall.O_WRONLY // open the file write-only.
		O_RDWR   int = syscall.O_RDWR   // open the file read-write.

		// The remaining values may be or'ed in to control behavior.
		O_APPEND int = syscall.O_APPEND // append data to the file when writing.
		O_CREATE int = syscall.O_CREAT  // create a new file if none exists.
		O_EXCL   int = syscall.O_EXCL   // used with O_CREATE, file must not exist.
		O_SYNC   int = syscall.O_SYNC   // open for synchronous I/O.
		O_TRUNC  int = syscall.O_TRUNC  // if possible, truncate file when opened.
	)
	*/

	// OPEN FOR READ, WRITE OR CREATE
	file, err := os.OpenFile("read.txt", os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		panic(err)
	}

	/* .... your code here ........ */

	if err := file.Close(); err != nil {
		panic(err)
	}

}