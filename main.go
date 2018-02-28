package main2

import(
	"fmt"
	"os"
	"time"
	"syscall"
)

func main () {

	// GOLANG FILE ACTIONS

	// CONSTANTS TO USE WHILE OPENING A FILE
	/*
	const (
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
	)
	*/

	// FILEINFO INTERFACE USED BY OS PACKAGE.
	// Example: When you open a file with os.OpenFile, you can get file name with file.Name().
	/*
	type FileInfo interface {
		Name() string       // base name of the file
		Size() int64        // length in bytes for regular files; system-dependent for others
		Mode() FileMode     // file mode bits
		ModTime() time.Time // modification time
		IsDir() bool        // abbreviation for Mode().IsDir()
		Sys() interface{}   // underlying data source (can return nil)
	}
	*/

	// FILE MODES/TYPES USED BY OS PACKAGE
	/*
	const (
        // The single letters are the abbreviations
        // used by the String method's formatting.
        ModeDir        FileMode = 1 << (32 - 1 - iota) // d: is a directory
        ModeAppend                                     // a: append-only
        ModeExclusive                                  // l: exclusive use
        ModeTemporary                                  // T: temporary file; Plan 9 only
        ModeSymlink                                    // L: symbolic link
        ModeDevice                                     // D: device file
        ModeNamedPipe                                  // p: named pipe (FIFO)
        ModeSocket                                     // S: Unix domain socket
        ModeSetuid                                     // u: setuid
        ModeSetgid                                     // g: setgid
        ModeCharDevice                                 // c: Unix character device, when ModeDevice is set
        ModeSticky                                     // t: sticky

        // Mask for the type bits. For regular files, none will be set.
        ModeType = ModeDir | ModeSymlink | ModeNamedPipe | ModeSocket | ModeDevice

        ModePerm FileMode = 0777 // Unix permission bits
	)
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

	// RENAME FILE OR DIRECTORY
	oldname := "yakuter"
	renname := "yakuter2"
	if err := os.Rename(oldname,renname); err != nil {
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
	changedir := "../"
	if err := os.Chdir(changedir); err != nil {
		panic(err)
	}

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

	// LEARN TEMP DIRECTORY
	fmt.Println(os.TempDir())

	// LEARN FILE PERMISSION
	fileperm, err := os.Stat("read.txt")
	if err != nil {
		panic(err)
	}

	mode := fileperm.Mode();

	fmt.Println(mode.Perm())

}