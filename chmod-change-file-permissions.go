package main

import(
	"os"
)

func main () {

	// CHMOD
	if err := os.Chmod("read.txt", 0755); err != nil {
		panic(err)
	}

}