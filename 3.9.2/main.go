package main

import (
	"crypto/rand"
	"os"
)

func main() {
	randReader := rand.Reader
	b := make([]byte, 1024)
	// io.ReadFull(randReader, b)
	randReader.Read(b)

	file, err := os.Create("new_rand.txt")
	if err != nil {
		panic(err)
	}
	file.Write(b)
}
