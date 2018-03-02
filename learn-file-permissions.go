package main

import(
	"fmt"
	"os"
)

func main () {

	// LEARN FILE PERMISSION
	fileperm, err := os.Stat("read.txt")
	if err != nil {
		panic(err)
	}

	mode := fileperm.Mode();

	fmt.Println(mode.Perm())

}