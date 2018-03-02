package main

import(
	"os"
)

func main () {

	// CREATE DIRECTORY
	createdirname := "yakuter"
	if err := os.Mkdir(createdirname,0755); err != nil {
		panic(err)
	}

}