package main

import (
	"fmt"
	"os"
)

func main() {

	// can handle symbolic link, but will no follow the link
	filePath := os.Args[1]
	fileInfo, err := os.Lstat(filePath)

	// cannot handle symbolic link
	//fileInfo, err := os.Lstat("file.txt")

	if err != nil {
		panic(err)
	}

	fmt.Println("Name : ", fileInfo.Name())

	fmt.Println("Size : ", fileInfo.Size())

	fmt.Println("Mode/permission : ", fileInfo.Mode())

	// --- check if file is a symlink

	if fileInfo.Mode()&os.ModeSymlink == os.ModeSymlink {
		fmt.Println("File is a symbolic link")
	}

	fmt.Println("Modification Time : ", fileInfo.ModTime())

	fmt.Println("Is a directory? : ", fileInfo.IsDir())

	fmt.Println("Is a regular file? : ", fileInfo.Mode().IsRegular())

	fmt.Println("Unix permission bits? : ", fileInfo.Mode().Perm())

	fmt.Println("Permission in string : ", fileInfo.Mode().String())

	fmt.Println("What else underneath? : ", fileInfo.Sys())

}
