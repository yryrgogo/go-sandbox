package main

import (
	"archive/zip"
	"io"
	"os"
	"strings"
)

func main() {
	file, err := os.Create("new.zip")
	if err != nil {
		panic(err)
	}

	zipWriter := zip.NewWriter(file)
	defer zipWriter.Close()

	w1, err := zipWriter.Create("f1.txt")
	if err != nil {
		panic(err)
	}
	io.Copy(w1, strings.NewReader("test zip file1"))

	w2, err := zipWriter.Create("f2.txt")
	if err != nil {
		panic(err)
	}

	io.Copy(w2, strings.NewReader("test zip file2"))
}
