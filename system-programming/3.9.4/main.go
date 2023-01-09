package main

import (
	"archive/zip"
	"io"
	"net/http"
	"os"
	"strings"
)

func handler(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/zip")
	w.Header().Set("Content-Disposition", "attachment; filename=ascii_sample.zip")

	zipFileName := createZip()
	file, err := os.Open(zipFileName)
	if err != nil {
		panic(err)
	}
	io.Copy(w, file)
	os.Remove(zipFileName)
}

func createZip() string {
	zipFileName := "new.zip"
	file, err := os.Create(zipFileName)
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

	return zipFileName
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
