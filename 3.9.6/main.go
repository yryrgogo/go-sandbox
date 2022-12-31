package main

import (
	"io"
	"os"
	"strings"
)

var (
	computer    = strings.NewReader("COMPURTER")
	system      = strings.NewReader("SYSTEM")
	programming = strings.NewReader("PROGRAMMING")
)

func main() {
	var stream io.Reader

	A := io.NewSectionReader(programming, 5, 1)
	S := io.LimitReader(system, 1)
	C := io.LimitReader(computer, 1)
	I1 := io.NewSectionReader(programming, 8, 1)
	I2 := io.NewSectionReader(programming, 8, 1)

	stream = io.MultiReader(A, S, C, I1, I2)

	io.Copy(os.Stdout, stream)
}
