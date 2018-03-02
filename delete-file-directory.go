package main

import(
	"os"
)

func main () {

	// REMOVE EMPTY FILE OR DIRECTORY
	removedirname := "yakuter"
	if err := os.Remove(removedirname); err != nil {
		panic(err)
	}

	// REMOVE FILE OR DIRECTORY AND ANY CHILDREN IT CONTAINS
	removealldirname := "yakuter"
	if err := os.RemoveAll(removealldirname); err != nil {
		panic(err)
	}

}