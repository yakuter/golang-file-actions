package main

import(
	"fmt"
	"os"
)

func main () {

	// LIST FILES IN A DIRECTORY
	listdirname := "."
	folder, err := os.Open(listdirname)
	if err != nil {
		panic(err)
	}

	files, 	err := folder.Readdir(0)
	if err != nil {
		panic(err)
	}

	if err := folder.Close(); err != nil {
		panic(err)
	}

	for _, file := range files {
		fmt.Println(file.Name())
	}

}