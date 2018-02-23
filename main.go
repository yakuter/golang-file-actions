package main

import(
	"fmt"
	"os"
)

func main () {

	// GOLANG FILE ACTIONS

	// OPEN FOR READ, WRITE OR CREATE
	file, err := os.OpenFile("read.txt", os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		panic(err)
	}

	data := make([]byte, 100)
	count, err := file.Read(data)
	if err != nil {
		panic(err)
	}

	fmt.Printf("read %d bytes: %q\n", count, data[:count])

	if err := file.Close(); err != nil {
		panic(err)
	}

	// WRITE
	f, err := os.OpenFile("write.txt", os.O_APPEND|os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		panic(err)
	}

	datapoem := "gözlerim kapalı"
	if _, err := f.Write([]byte(datapoem+"\n")); err != nil {
		panic(err)
	}
	if err := f.Close(); err != nil {
		panic(err)
	}

	// CHMOD
	if err := os.Chmod("read.txt", 0755); err != nil {
		panic(err)
	}

	// CHDIR
	newdir := "/etc"
	if err := os.Chdir(newdir); err != nil {
		panic(err)
	}

	// LIST FILES IN A DIRECTORY
	dirname := "."
	folder, err := os.Open(dirname)
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