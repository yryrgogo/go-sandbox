package main

import (
	"io"
	"os"
)

func main() {
	file, err := os.Open("main.go")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	newFile, err := os.Create("new_file.txt")
	if err != nil {
		panic(err)
	}
	defer newFile.Close()

	io.Copy(newFile, file)
}
