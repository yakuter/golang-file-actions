package main

import(
	"fmt"
	"os"
)

func main () {

	// GOLANG FILE ACTIONS

	// CONSTANTS TO USE WHILE OPENING A FILE
	/*
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
	*/

	// OPEN FOR READ, WRITE OR CREATE
	file, err := os.OpenFile("read.txt", os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		panic(err)
	}

	// LEARN UID OR GID OF USER
	gid := os.Getgid()
	uid := os.Geteuid()

	fmt.Println(gid) 	// GID
	fmt.Println(uid)	// UID

	// CHANGE OWNER OF A FILE
	// needs permission. You can use "sudo go run main.go" to run

	if err := os.Chown(file.Name(), gid, uid); err != nil {
		panic(err)
	}

	// PRINT FILE CONTENT
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

	// CREATE DIRECTORY
	createdirname := "yakuter"
	if err := os.Mkdir(createdirname,0755); err != nil {
		panic(err)
	}

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

	// CHDIR
	newdir := "../"
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