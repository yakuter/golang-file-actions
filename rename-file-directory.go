package main

import(
	"os"
)

func main () {

	// RENAME FILE OR DIRECTORY
	oldname := "yakuter"
	renname := "yakuter2"
	if err := os.Rename(oldname,renname); err != nil {
		panic(err)
	}

}