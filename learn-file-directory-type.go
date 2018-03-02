package main

import(
	"fmt"
	"os"
)

func main () {

	// LEARN FILE/DIRECTORY TYPE
	filetype, err := os.Stat("read.txt")
	if err != nil {
		panic(err)
	}

	switch mode := filetype.Mode(); {

		case mode.IsRegular():
			fmt.Println("regular file")

		case mode.IsDir():
			fmt.Println("directory")

		case mode&os.ModeSymlink != 0:
			fmt.Println("symbolic link")

		case mode&os.ModeNamedPipe != 0:
			fmt.Println("named pipe")
	}

}