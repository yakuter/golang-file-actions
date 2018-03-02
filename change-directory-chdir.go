package main

import(
	"os"
)

func main () {

	// CHDIR
	changedir := "../"
	if err := os.Chdir(changedir); err != nil {
		panic(err)
	}

}